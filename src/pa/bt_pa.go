/*********************************************************
/* bottom calculation
/* ----
/*  License: BSD
/* ----
/* v0.1   bottom calculation. --- 2019.1.15
/*********************************************************/

package pa

import (
        . "cmn"
	"fmt"
        . "define"
        //"github.com/sjwhitworth/golearn/base"
        //"github.com/sjwhitworth/golearn/ensemble"
        //"github.com/sjwhitworth/golearn/evaluation"
       )





/************************************************************************
/*2018.12: 按市盈率等指标来看，目前的熊市已经创了历史新低的估值了。
历史新低没有，还是比2014年5、6月份高一点，当时最低是9.76，现在是12.48，不过前后四年的对比是价值股
上天了，如果剔除茅台、平安等价值股对指数的影响，那么现在已经是9.76附近

/-------------------------两融余额----------------------------------

// 这样2015年度A股正式的两融业务总额最高是2.26万亿，其他杠杆资金是配资，而掐死配资业务后，很多类似我这样的赌徒都只能选择券商的两融业务
//...，然后加上这几年市场规模扩大太快，两融标的池子的两融券种也大幅提高，那么这轮行情到顶峰时，两市融资余额过3万亿，甚至是到4万亿都将是常态

//2010年3月31日才开始发展两融业务，2012、2013年股民参与两融还不是普遍性，直到2015年初牛市行情确认后，两融业务才真正快速发展，"

*/



//---- PE ----
const(
PE_BOTTOM_REC6_SH_CINDEX = 13.1
PE_BOTTOM_REC5_SH_CINDEX = 12.0         //最低为11.91，出现在2014年
PE_BOTTOM_REC4_SH_CINDEX = 14.0
PE_BOTTOM_REC3_SH_CINDEX = 18.0
)

//--- 破净 ---
const(
PB_LESS_THAN1_MAX_NUM_REC7 =  428               // REC4 = 2018.12.27
TOTAL_ASTOCK_NUM_REC7      = 3553

PB_LESS_THAN1_MAX_NUM_REC5 =  168               // REC4 = 2018.12.27
TOTAL_ASTOCK_NUM_REC5      = 1575

PB_LESS_THAN1_MAX_NUM_REC4 =  173               // REC4 = 2018.12.27
TOTAL_ASTOCK_NUM_REC4      = 1305
)

const(
    EXTREME_VOL_TOP_DIV_BOT = 11
)


/***************************************** 底部函数 *******************************************************
/* 常用的评价函数方法有"线性加权和法"、"极大---极小法"和"理想点法"
// 一些常见的权重计算方法包括：熵值法、主成分分析、因子分析、AHP层次分析法、模糊综合评价、灰色关联法、TOPSIS、DEA包络分析等计算权重的
// 方法，有兴趣可以登录官网查看和使用。
/* PE:
/* PB: "上面提到的规律在发展中国家股市中也适用。如上图所示，如果该国股市的市净率低于1.22，那么接下来4年
/*      的股市平均回报为每年13%左右。如果市净率介于1.22和2.76之间，那么接下来4年的股市年回报在9%左右。而如果
/*      该国股市的市净率超过2.76，那么接下来股市下跌的概率会比较大，每年的回报为-5%左右  -----作者：伍治坚
***********************************************************************************************************/


//------------------------- level 1: f1确定权重(熵权法) -----------------------------
// wmap: weight output map use entropy method
func f1_weight_cha()(w []float64 ){

	dmEigRlx, dmEigPuc, dmEigHot, dmEigCrz, _ := GetEigDm(FN_EGG_BOT_RLX,FN_EGG_BOT_PUC,FN_EGG_TOP_CRZ,FN_EGG_TOP_HOT )
	fmt.Printf("<f1_weight_cha> dmEigPuc:%v, dmEigHot:%v,dmEigCrz:%v  \n", dmEigPuc, dmEigHot, dmEigCrz )

	w = WeightEnt(dmEigRlx)
	return
}



//---------------------------- level 2: f2 : 估值函数 ------------------------------
// [Normalization] + [Weight]
func EvalPos(eig_min, eig_max[]float64, dicMkt map[string]float64 )(pos float64){
	var dicMktS = make(map[string](map[string]float64) )
	dicMktS["dummyTag1"] = dicMkt
	//dicMktS["dummyTag2"] = dicMkt

	dm := Eggs2Dm(dicMktS)                                  // dm is [][]float64
	DmClean(&dm)             				// clean before use
	if DEBUG_MIPOS{ fmt.Printf("Info: <EvalCurPos>  dm: %v \n", dm)  }

	// step 2: Normalize it bases on eig_min, eig_max
	d_p := Norm_EvtsDm(dm, eig_min, eig_max)           // dm[0] is 1d slice
	if DEBUG_MIPOS{ fmt.Printf("Info: <EvalCurPos>  d_p: %v \n", d_p) }

	// step 3: weight it
	pos_slc := f2_evalEigs(d_p)
	fmt.Printf("Info : <EvalCurPos>  pos_slc: %v \n", pos_slc)

	return pos_slc[0]
}


// eval multiple Eigs (by applying "weight")
func f2_evalEigs(dm_eigs [][]float64)(so []float64){
	dm_eigs_T := TranposeDm(dm_eigs)

	for _, evtEig := range dm_eigs_T{
		so = append(so, f2_evalOnEig(evtEig) )
	}
	//fmt.Printf("===> <f2_evalEigs>, so: %v \n", so)
	return
}

// apply "weight" on single eig
func f2_evalOnEig(dm_evt []float64 )(sum float64){
	sum = 0.0
        w_ent := []float64{0.30,0.01,0.02, 0.33,0.01,0.02, 0.33,0.01,0.02}   //pe, pb, tnr       // dummy. not use entropy method now
        r     := []float64{0.5, 0.0, 0.0,  0.4, 0.0, 0.0,  0.1, 0.0, 0.0 }

        w_scr := WeightSwc(w_ent, r)

	for j, v := range dm_evt{
		eva := w_scr[j] * v
		sum += eva
	}
	//fmt.Printf("--> <f2_evalOnEig> w_scr: %v, sum:%v  \n", w_scr, sum)
	return
}

//------------ Gen eval(估值) parameters ---------------
// Gen Crz/Hot/Relax/Puc character values.
// Average + Normalize + Weight(Factors)
func Chrp( )(crz_slc, hot_slc, puc_slc, rlx_slc []float64, eigMin, eigMax []float64){

	dmEigRlx, dmEigPuc, dmEigHot, dmEigCrz, _ := GetEigDm( FN_EGG_BOT_RLX,FN_EGG_BOT_PUC,FN_EGG_TOP_CRZ,FN_EGG_TOP_HOT )
	var dm_eigAll [][]float64 = make([][]float64, len(dmEigRlx) )

	DmAppend(&dm_eigAll, dmEigRlx)
	DmAppend(&dm_eigAll, dmEigPuc)
	DmAppend(&dm_eigAll, dmEigHot)
	DmAppend(&dm_eigAll, dmEigCrz)
	DmClean(&dm_eigAll)
	//fmt.Printf("Info: <Chrp>, dm_eigAll: %v   \n", dm_eigAll)

	// step 1: get min_max
	min_cha, max_cha := GetMinMax( dm_eigAll )               // (min, max []float64)
	if DEBUG_CHRP { fmt.Printf("Info: <Chrp>, min_cha: %v, max_cha: %v   \n", min_cha, max_cha) }

	// step 2: Normalize crz/hot/puc/rlx
	DmClean(&dmEigRlx); DmClean(&dmEigPuc); DmClean(&dmEigHot); DmClean(&dmEigCrz)
	//var dmNormCrz, dmNormHot, dmNormPuc, dmNormRlx [][]float64
	dmNormCrz := Norm_EvtsDm(dmEigCrz, min_cha, max_cha)
	dmNormHot := Norm_EvtsDm(dmEigHot, min_cha, max_cha)
	dmNormPuc := Norm_EvtsDm(dmEigPuc, min_cha, max_cha)
	dmNormRlx := Norm_EvtsDm(dmEigRlx, min_cha, max_cha)
	if DEBUG_CHRP { fmt.Printf("Info: <Chrp>, dmNormCrz:%v\n, dmNormHot:%v\n, dmNormPuc:%v\n, dmNormRlx:%v  \n", dmNormCrz, dmNormHot, dmNormPuc, dmNormRlx) }

	// setp 3: weight it (each type)
	crz_slc = f2_evalEigs(dmNormCrz )
	hot_slc = f2_evalEigs(dmNormHot )
	puc_slc = f2_evalEigs(dmNormPuc )
	rlx_slc = f2_evalEigs(dmNormRlx )
	if DEBUG_CHRP{ fmt.Printf("Info: <Chrp:Weighted>: crz_slc:%v, hot_slc:%v, puc_slc:%v, rlx_slc:%v  \n", crz_slc, hot_slc, puc_slc, rlx_slc) }

	eigMin = min_cha
	eigMax = max_cha

	return
}



//---------------- level 3: f3 : bot_market result ---------------------------------
// Now it's dummy func
// 3 factors -> 1 indication value. 
func f3_tri_MktEval(mkt T_Market )(ind float64){       // 市场底，政策底，情绪底，
	//crz, hot, puc, rlx := Chrp()
	return
}




/***********************************  情绪底 *************************************
--- 成交量 ---
//俗话说，多头不死，空头不止。一般底部来临时，绝大多数人都对市场绝望了，该斩仓出局的人大多都已经退出，剩
下的是一些零星的卖出盘，基本没有太多人愿意买入。而地量形成的原因是没有人愿意卖出也没有人愿意买入，股价
处于极度低迷的状态，也意味着市场缺乏增量资金，市场情绪不高。但另一方面也意味着空头势力衰弱，多空双方暂
时达成平衡，市场底部往往伴随地量盘整。
对A股市场而言，长期的市场特征都是底部永远是在市场万念俱灰的情况下完成，市场顶部永远是在市场狂欢中铸成，
这是永远的规律。
******************************/
//------- 两融余额 -------
// "东吴现在两融余额是8.8亿"
// "国金现在的两融余额是17.03亿，2015年股灾爆发前后，国金盘口的两融余额是89亿，现在才17.03亿"



//----- 两融余额占市值比/流通市值比 -----
// "这几天国金盘口的融资买入额占总市值的占比才5%左右。市场疯狂时，融资买入额占市值占比是20-30%。"
//主升段时，两融余额占流通市值的比例是8、9%，


//---- 融资买入额占成交比 -----
//2014年券商板块行情的主升段是11-12月初，当时东北的盘口是融资买入额动辄就是7、8亿，甚至10亿，现在是多少？2.93亿
//"两融交易额占成交额的占比，多时是40%，现在东北证券盘口的两融成交额占日成交额的占比是多少？是22%"
//"上轮行情东北证券的两融买入额最大时是15.84亿，那么这轮行情如果起来了，两融资金当天买入东北证券可否到单日20亿？而现在的 $
//-------------------------两融余额----------------------------------

// 这样2015年度A股正式的两融业务总额最高是2.26万亿，其他杠杆资金是配资，而掐死配资业务后，很多类似我这样的赌徒都只能选择券商的两融业务
//...，然后加上这几年市场规模扩大太快，两融标的池子的两融券种也大幅提高，那么这轮行情到顶峰时，两市融资余额过3万亿，甚至是到4万亿都将是常态

//2010年3月31日才开始发展两融业务，2012、2013年股民参与两融还不是普遍性，直到2015年初牛市行情确认后，两融业务才真正快速发展，"




/*
func bot_emo_calc(int last_top_vol, cur_vol)(res float64){
	
   res = ( ltopvol / cur_vol ) /  (EXTREME_VOL_TOP_DIV_BOT)
	
}

func getLastTopVol()ltopvol int{
	req_item ：= "TRADE_VOLUME"
	curDate date := LastTopDate
	last_top_vol := qif(LastBotDat, req_item)

	for i:=0; i<len(last_top_vol); i++{
		ltopvol += last_top_vol[i]
	}
	return ltopvol
}

func getCurVol()curVol int{
	req_item := 
	curDate date := 
	cur_trade_Vol := qif(curDate, req_item)	

	for i:=0; i<len(last_top_vol); i++{
		curVol += last_top_vol[i]
	}
}
*/


//----------------- 新开户数量 ----------------
//2019.3 市场高峰期时，每周开户数量可以过百万户，现在每周开户数量才30万户，说明市场还没开始疯狂，那么等这次市场每周

/* ---  bot ind -------
中国股市作为一个周期性很强的市场，每当牛市中后期，“中国梦”、“黄金十年”之类的概念满天飞，而熊市末期，则到处流传各
类“鬼故事”。似乎唯有市场现象规律性的统计结果才可作为投资的决策。至少有4个现象可以作为市场见底的信号：

1、成交量： 降10/11：A 股市场，上涨期的顶部交易量与随后的下跌底部交易量之间有个大约的 11:1 关系（即底部成交约为前期成交
顶部的 1/11）。市场本轮下跌前的成交量单日高峰达到约 2.4 万亿元左右，前期 A 股日均成交量也已经下行至 2200 亿元附近，
在接近这一历史规律。

2、下跌幅度： 中位数达70%左右。历史上的几次深度调整时期，如1993-1994 年、2001-2005 年、2007-2008 年，个股的下跌幅度中
位数均在 70%左右（分别为75%、76%、69%，2009-2014 年由于存在结构性行情，跌幅中位数为 12%）。当前市场个股与 2015 年
中的高点相比，跌幅中位数已经达到 73%。


4、市场估值：  三大指数1月4日的市盈率（TTM），和上证指数前期三次大底部当日的市盈率（TTM）比起来，
中证500指数已是历史最低，而上证指数和沪深300指数则只高于2013年的那个底部（表4）。

3、估值调整幅度： 2008 与2014 年低点时动态市盈率较前期高点下滑中位数分别为 72%和51%，动态市净率下滑中位
数为 75%和 57%。目前市场动态市盈率和动态市净率较 2015 年高点时期的下滑中位数分别为 78%和 79%。

5. 融资额度降低： 

*/


/******************************* 政策底 *********************************/



