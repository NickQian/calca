#!/usr/bin/env python
"""
# JoinQuant(聚宽)的JQData的python接口.   www.joinquant.com/data    
#  ----
#  License: BSD 
#  ----
#  0.1: init version - 2019.2 - Nick cKing
"""
import datetime
import jqdatasdk as jq
import numpy as np
import pandas as pd

from sqlalchemy import Integer, Column, create_engine, ForeignKey
from sqlalchemy.orm import relationship, joinedload, subqueryload, Session

def auth(username,password):
    return jq.auth(username, password)
    
def getCurPE():    
    pdf = getMarketYesterday()
    pe_sha = pdf[0:1]['pe_average'][0]               # 上海A股
    pe_sh  = pdf[2:3]['pe_average'][2]               # 上海市场
    pe_szm = pdf[4:5]['pe_average'][4]               # 深主板
    pe_sz  = pdf[5:6]['pe_average'][5]               # 深市场
    pe_gem = pdf[6:7]['pe_average'][6]               # 创业板
    return  { 'pe_sh':pe_sh, 'pe_sz':pe_sz, 'pe_gem':pe_gem}
    
    
def getCurVol():   # 单位：亿人民币
    pdf = getMarketYesterday()
    vol_sh  = pdf[2:3]['money'][2]                   # 上海市场
    vol_sz  = pdf[5:6]['money'][5]                   # 深市场  
    vol_gem = pdf[6:7]['money'][6]                   # 创业板
    return  { 'vol_sh':vol_sh, 'vol_sz':vol_sz, 'vol_gem':vol_gem}


def getCurTor():   # turnOver ratio. 单位：％
    pdf = getMarketYesterday()
    tor_sh  = pdf[2:3]['turnover_ratio'][2]           # 上海市场
    tor_sz  = pdf[5:6]['turnover_ratio'][5]           # 深市场  
    tor_gem = pdf[6:7]['turnover_ratio'][6]           # 创业板
    return  {'tor_sh':tor_sh,  'tor_sz':tor_sz, 'tor_gem':tor_gem}
    

# 2019.2.28:等待JQdatasdk更新,他们工程师已经建议添加
def getCurMtss():  
    pass

# market PB: 得使用申万数据，使用jy查询     
def getCurPB():    
    pdf = ()
    
    
    
    
#----------------------------- single -----------------------------------

    
# get today price
def getSinglePrice(code):
    today = time.strftime('%Y-%m-%d', time.localtime(time.time()))
    return jq.get_price(code, start_date= today, end_date= today, fields=['open','close'] )

# 获取单个票/某天的融资余额
def getSingleMtss(code, date):
    return jq.get_mtss(code, date)
    

    
    
#--- 公司财务数据 ---
#获取单季度/年度财务数据
get_fundamentals(query(
            #下一行为估值指标，包括代码、市值、pe、pb、ps、pcf，可自由添加其它指标
            valuation.code, valuation.market_cap, valuation.pe_ratio, valuation.pb_ratio,valuation.ps_ratio,valuation.pcf_ratio,\
            #下一行为资产负债表，包括报表发布日期、报表期最后一天、总资产、可自由添加其它指标
            balance.pubDate,balance.statDate,balance.total_assets,\
             #下一行为利润表，包括报营业总收入、净利润、可自由添加其它指标
            income.total_operating_revenue,income.net_profit
            #可自由添加其它表和指标……
 
        ).filter(
 
            valuation.code.in_(stockList), #指定股票为stockList的股票，如果这个代码删掉，则为全部股票
            #可在下面加入各种筛选标准，如筛选出pe<10、Pb<3的股票
            # valuation.pe_ratio < 10,
            # valuation.pb_ratio<3
 
        ).order_by(
            # 按市值降序排列
            valuation.market_cap.desc()
        ).limit(
 
            # 最多返回个数，最大不超过10000行
            10000
        ), date=calDayFormat)
    df["date"]=calDayFormat



    
####================================ 内部使用函数 ============================================

#---------- 获取市场位置信息----------
def getMarket(date):
    pdf =jq.finance.run_query(\      # type(l): pandas.core.frame.DataFrame
                           jq.query(jq.finance.STK_EXCHANGE_TRADE_INFO)\       # type: sqlalchemy.orm.query.Query
                           .filter(jq.finance.STK_EXCHANGE_TRADE_INFO.date==date)\
                           .limit(10))
    return pdf

def getMarketYesterday():
    #today = time.strftime('%Y-%m-%d', time.localtime(time.time()))
    today = datetime.date.today()
    oneday = datetime.timedelta(days=1)
    yesterday = (today - oneday).strftime('%Y-%m-%d')
    return getMarket(yesterday)

# 申万指数，use jy: 
def getSWdata(code, end_date=None, count=None, start_date=None):
    if isinstance(code, str):
        code=[code]
    days = get_trade_days(start_Date, end_date, count)
    
    df = jq.jy.run_query(jq.query(jq.jy.QT_SYWGIndexQuote.InnerCode.distinct().label('InnerCode'))\
                         .filter(jq.jy.QT_SYWGIndexQuote.TradingDay==day,  ))
                   
    code_df = jq.jy.run_query(jq.query(jq.jy.SecuMain.InnerCode, jq.jy.SecuMain.SecuCode, jq.jy.SecuMain.ChiName)\
                                     .filter(jq.jy.SecuMain.SecuCode.in_(df.Innercode)))
    df = 

#---------- 获取成分股 -------------
# 获取现在上证50列表
def getSh50()
    return jq.get_index_stocks('000016.XSHG')
    
def getSh180()  # 上证180
    return jq.get_index_stocks('000010.XSHG')

def getHs300()  # 沪深300
    return jq.get_index_stocks('000300.XSHG')


def getZz500()  # 中证500
    return jq.get_index_stocks('000905.XSHG')
    

    
    
    
    
if __name__ == "main":
    auth("18602122079", "calcaapi")
    