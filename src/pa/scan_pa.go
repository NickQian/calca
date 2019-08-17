// 本个股扫描完全基于价投“称重机模型”。
// 涛哥教导：即便是赌，也不能压注在垃圾上；因为垃圾身上你不知道会发生什么
// -------------------------------------------------------------------------------------------------------------------
// v0.01 - policy A的个股扫描: 基于 1) & 4).   --- 2018.8.30

package pa

import (
        "fmt"
        )

// --------------------------估值函数 ----------------------------
func pa_single() (int, int, int, int){
    sug_code := [] int{}
    cb_intension := int{50}
    
 
}



// ====================== 三碗面 ================================
// 基本面决定趋势，技术面决定位置，情绪面决定日常波动区间




//----------------------- 优质公司ROE-----------------------------




//---------- 使用毛利润与资产之比（Gross-Profit-to-Assets,GP/A）--------





//----------------------- 盈利能力 和 PE  -----------------------
// 盈利能力(毛利率、净利润率的高低)保证上行，PE保证价格便宜




//-------------------------是否龙头------------------------------
// 未来新增产能情况（决定未来可发展的可持续性），行业排序（决定是否是行业龙头及辅助判断坑风险能力）



//------------------------ 营收占款的占比 -----------------------
//（决定财务结构是否健康）



//------------------------ 新增业务收入情况----------------------
//（决定业务增长情况）




//-------------------------- 董事长打分--------------------------




//------------------------- 员工参与度 ----------------------


//我个人选票原则，只在沪深300+中证500这800只成分股里选


//----------------------- 选股范围 ----------------------------
//我这几年来圈定的选股范围，一直是沪深300+中证500这800只股票，这话我说过10次不止了，从这800只股票里选标的，踩雷的概率会
//降低很多。然后我此前也说过，A股国际化后，只有成分股才会受到外资的追捧，后期大多数非指标股、成分股会逐渐被市场边缘化...



//=========================== Filter =============================
//------------------------- 不碰 ----------------------------------
//然后养殖不碰、养花种草、农业类的不碰，科技软件基本不碰、环保工程的不碰，公用事业不碰，综合类企业不碰，
//频繁重组资产并购的不碰，民企基本不碰，影视娱乐传媒行业不碰，我有这些不碰原则后，然后我又只在沪深300+中证500里
//选票，这样我个人踩雷的概率已经极低。



//-----------------------  高换手率 -------------------------------


///////////////////////// JoinQuant py低估价值选股代码 //////////////////////
'''
1.市净率小于2；
2.负债比例高于市场平均值;
3.企业的流动资产至少是流动负债的1.2倍;
4.每年四次调仓，即在1/4/7/10月调仓；
5.可加入止损(十天HS300跌幅达10%清仓)；
'''

## 初始化函数，设定要操作的股票、基准等等
def initialize(context):
    # 设定指数
    g.stockindex = '000300.XSHG' 
    # 设定沪深300作为基准
    set_benchmark('000300.XSHG')
    # True为开启动态复权模式，使用真实价格交易
    set_option('use_real_price', True) 
    # 设定成交量比例
    set_option('order_volume_ratio', 1)
    # 股票类交易手续费是：买入时佣金万分之三，卖出时佣金万分之三加千分之一印花税, 每笔交易佣金最低扣5块钱
    set_order_cost(OrderCost(open_tax=0, close_tax=0.001, \
                             open_commission=0.0003, close_commission=0.0003,\
                             close_today_commission=0, min_commission=5), type='stock')
    # 最大持仓数量
    g.stocknum = 10

    ## 自动设定调仓月份（如需使用自动，注销下段）
    f = 4  # 调仓频率
    log.info(range(1,13,12/f))
    g.Transfer_date = range(1,13,12/f)
    
    ## 手动设定调仓月份（如需使用手动，注释掉上段）
    # g.Transfer_date = (3,9)
    
    #根据大盘止损，如不想加入大盘止损，注释下句即可
    # run_daily(dapan_stoploss, time='open') 
    
    ## 按月调用程序
    run_monthly(trade, monthday=20, time='open')

## 交易函数
def trade(context):
    # 获取当前月份
    months = context.current_dt.month
    # 如果当前月为交易月
    if months in g.Transfer_date:
        ## 获得Buylist
        Buylist = check_stocks(context)
        
        ## 卖出
        if len(context.portfolio.positions) > 0:
            for stock in context.portfolio.positions.keys():
                if stock not in Buylist:
                    order_target(stock, 0)

        ## 分配资金
        if len(context.portfolio.positions) < g.stocknum :
            Num = g.stocknum  - len(context.portfolio.positions)
            Cash = context.portfolio.cash/Num
        else: 
            Cash = 0

        ## 买入
        if len(Buylist) > 0:
            for stock in Buylist:
               if stock not in context.portfolio.positions.keys():
                   order_value(stock,Cash)
    else:
        return
    
## 选股函数
def check_stocks(context):
    # 获取沪深成分股
    security = get_index_stocks(g.stockindex)

    Stocks = get_fundamentals(query(
            valuation.code,
            valuation.pb_ratio,
            balance.total_assets,
            balance.total_liability,
            balance.total_current_assets,
            balance.total_current_liability
        ).filter(
            valuation.code.in_(security),
            valuation.pb_ratio < 2, #市净率低于2
            balance.total_current_assets/balance.total_current_liability > 1.2 #流动资产至少是流动负债的1.2倍
        ))
    
    # 计算股票的负债比例
    Stocks['Debt_Asset'] = Stocks['total_liability']/Stocks['total_assets']
    # 获取负债比率的市场均值
    me = Stocks['Debt_Asset'].median()
    # 获取满足上述条件的股票列表
    Codes = Stocks[Stocks['Debt_Asset'] > me].code

    return list(Codes)

## 根据局大盘止损，具体用法详见dp_stoploss函数说明
def dapan_stoploss(context):
    stoploss = dp_stoploss(kernel=2, n=3, zs=0.1)
    if stoploss:
        if len(context.portfolio.positions)>0:
            for stock in list(context.portfolio.positions.keys()):
                order_target(stock, 0)

## 大盘止损函数
def dp_stoploss(kernel=2, n=10, zs=0.03):
    '''
    方法1：当大盘N日均线(默认60日)与昨日收盘价构成“死叉”，则发出True信号
    方法2：当大盘N日内跌幅超过zs，则发出True信号
    '''
    # 止损方法1：根据大盘指数N日均线进行止损
    if kernel == 1:
        t = n+2
        hist = attribute_history('000300.XSHG', t, '1d', 'close', df=False)
        temp1 = sum(hist['close'][1:-1])/float(n)
        temp2 = sum(hist['close'][0:-2])/float(n)
        close1 = hist['close'][-1]
        close2 = hist['close'][-2]
        if (close2 > temp2) and (close1 < temp1):
            return True
        else:
            return False
    # 止损方法2：根据大盘指数跌幅进行止损
    elif kernel == 2:
        hist1 = attribute_history('000300.XSHG', n, '1d', 'close',df=False)
        if ((1-float(hist1['close'][-1]/hist1['close'][0])) >= zs):
            return True
        else:
            return False


