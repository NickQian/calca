#!/usr/bin/env python
# -*- coding:utf-8 -*-
"""# JoinQuant(聚宽)的JQData的python接口.   www.joinquant.com/data
#  ----
#  License: BSD
#  ----
#  0.2: with jqdatasdk-1.8.7. Has data from 2005                  - 2021.2
#  0.1: init version(jqdatasdk-1.6.5). Only has data from 2009??  - 2019.2 - Nick cKing. 
"""


import datetime
import jqdatasdk as jq
import numpy as np
import pandas as pd

from sqlalchemy import Integer, Column, create_engine, ForeignKey
from sqlalchemy.orm import relationship, joinedload, subqueryload, Session

"""old version login"""
#import logging
#logging.basicConfig(filename='./pyinfo.log', format='[%(asctime)s-%(filename)s-%(levelname)s:%(message)s]',level=logging.DEBUG, filemode='a', detefmt='%Y-%m-%d%I:%M:%S %p')

def Login_JQ(username,password):
	print ("PyInfo: JQ login ing....... \n")
	suc = False
	try:
		JQLoginRes = jq.auth(username, password)
		print ("PyInfo: JQLoginRes:", JQLoginRes)
		logging.info("Pyinfo:Login_JQ success.")
		suc = True
		print ("PyInfo: JQLogin successfully.")
	except:
		suc = False
		print ("PyInfo: JQLogin Failed! ")
	return suc



def getMarketMap(day):
	pdf = getMarket(day)
	print ("----pdf is:---", pdf)
	pe_sha,  pe_sh,  pe_sz,  pe_gem,  pe_szm  = None, None, None, None, None
	pb_sha,  pb_sh,  pb_sz,  pb_gem,  pb_szm  = None, None, None, None, None
	tnr_sha, tnr_sh, tnr_sz, tnr_gem, tnr_szm = None, None, None, None, None
	cmc_sha, cmc_sh, cmc_sz, cmc_gem,  pb_szm  = None, None, None, None, None

	if pdf.empty:
		print("pyInfo: Err: <getMarket> result is empty! ")
		return {}
	else:
		pe_sha  = pdf[0:1]['pe_average'][0]               # [0]上海A股. [0]->[0:1]
		pe_sh   = pdf[2:3]['pe_average'][2]               # [2]上海市场 [2]->[2:3]
		pe_szm  = pdf[4:5]['pe_average'][4]               # [4]深主板
		pe_sz   = pdf[5:6]['pe_average'][5]               # [5]深市场
		pe_gem  = pdf[6:7]['pe_average'][6]               # [6]创业板
		tnr_sha = pdf[0:1]['turnover_ratio'][0]
		tnr_sh  = pdf[2:3]['turnover_ratio'][2]
		tnr_szm = pdf[4:5]['turnover_ratio'][4]
		tnr_sz  = pdf[5:6]['turnover_ratio'][5]
		tnr_gem = pdf[6:7]['turnover_ratio'][6]
		cmc_sha = pdf[0:1]['circulating_market_cap'][0]
		cmc_sh  = pdf[2:3]['circulating_market_cap'][2]
		cmc_szm = pdf[4:5]['circulating_market_cap'][4]
		cmc_sz  = pdf[5:6]['circulating_market_cap'][5]
		cmc_gem = pdf[6:7]['circulating_market_cap'][6]
		#vol_sha = pdf[0:1]['money'][0]
		#vol_sh  = pdf[2:3]['money'][2]
		#vol_szm = pdf[4:5]['money'][4]
		#vol_sz  = pdf[5:6]['money'][5]
		#vol_gem = pdf[6:7]['money'][6]
		mtss_sh, mtss_sz = getMtss(day)
		return  {'pe_sha':pe_sha,  'pe_sh':pe_sh,    'pe_szm':pe_szm,   'pe_sz':pe_sz,   'pe_gem':pe_gem,
        	         'pb_sha':pb_sha,  'pb_sh':pb_sh,    'pb_szm':pb_szm,   'pb_sz':pb_sz,   'pe_gem':pb_gem,
        		 'tnr_sha':tnr_sha,'tnr_sh':tnr_sh,  'tnr_szm':tnr_szm, 'tnr_sz':tnr_sz, 'tnr_gem':tnr_gem,
        	  	 'cmc_sha':cmc_sha,'cmc_sh':cmc_sh,  'cmc_szm':cmc_szm, 'cmc_sz':cmc_sz, 'cmc_gem':cmc_gem,
        	  	 'mtss_sh':mtss_sh,'mtss_sz':mtss_sz,
        	  	 #'vol_sha':vol_sha,'vol_sh':vol_sh,'vol_szm':vol_szm,'vol_sz':vol_sz,'vol_gem':vol_gem,
        		}



# 2019.2.28:等待JQdatasdk更新, 已经建议他们工程师添加
def getMtss(date):   #date eg: '2019-05-23'
	pd = jq.finance.run_query(      # type(l): pandas.core.frame.DataFrame
	                          jq.query(jq.finance.STK_MT_TOTAL)       # type: sqlalchemy.orm.query.Query
	                          .filter(jq.finance.STK_MT_TOTAL.date==date)
	                          .limit(10)
	                          )
	#print("<getMtss> pd, pd[0:1]:", pd)
	mtss_sh = pd[0:1]['fin_value'][0]
	mtss_sz = pd[1:2]['fin_value'][1]
	return mtss_sh, mtss_sz



def getPE(day):
    	pdf = getMarket(day)
    	pe_sha = pdf[0:1]['pe_average'][0]               # 上海A股
    	pe_sh  = pdf[2:3]['pe_average'][2]               # 上海市场
    	pe_szm = pdf[4:5]['pe_average'][4]               # 深主板
    	pe_sz  = pdf[5:6]['pe_average'][5]               # 深市场
    	pe_gem = pdf[6:7]['pe_average'][6]               # 创业板
    	return  { 'pe_sh':pe_sh, 'pe_sz':pe_sz, 'pe_gem':pe_gem}


def getVol(day):   # 单位：亿人民币
    pdf = getMarket(day)
    vol_sh  = pdf[2:3]['money'][2]                   # 上海市场
    vol_sz  = pdf[5:6]['money'][5]                   # 深市场
    vol_gem = pdf[6:7]['money'][6]                   # 创业板
    return  { 'vol_sh':vol_sh, 'vol_sz':vol_sz, 'vol_gem':vol_gem}


def getTor(day):   # turnOver ratio. 单位：％
    pdf = getMarket(day)
    tor_sh  = pdf[2:3]['turnover_ratio'][2]           # 上海市场
    tor_sz  = pdf[5:6]['turnover_ratio'][5]           # 深市场
    tor_gem = pdf[6:7]['turnover_ratio'][6]           # 创业板
    return  {'tor_sh':tor_sh,  'tor_sz':tor_sz, 'tor_gem':tor_gem}



# market PB: 得使用申万数据，使用jy查询
def getPB(day):
	pdf = ()




#------------------------ Trade days -----------------------------------
def getTradeDays(dateStr, num_prev):
	validays = []
	numpre = int(num_prev)
	date = datetime.datetime.strptime(dateStr, '%Y-%m-%d')
	tradays = jq.get_trade_days(end_date = date, count=numpre)
	list_tradays = tradays.tolist()         # datetime in list
	for day in list_tradays:
		validays.append(day.strftime("%Y-%m-%d") )
	return validays

#----------------------------- single -----------------------------------
# get today price
"""get_price 与 get_fundamentals_continuously 接口 panel 参数将固定为 False
注意：0.25 以上版本 pandas 不支持 panel，如使用该数据结构和相关函数请注意修改"""
def getSinglePrice(code):
    today = time.strftime('%Y-%m-%d', time.localtime(time.time()))
    return jq.get_price(code, start_date= today, end_date= today, fields=['open','close'] )


# 获取单个票/某天的融资余额
def getSingleMtss(code, date):
    return jq.get_mtss(code,end_date=today,count =1)



"""#--- 公司财务数据 ---
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
"""



####================================ 内部使用函数 ============================================
def getMarket(date):
    pdf =jq.finance.run_query(      # type(l): pandas.core.frame.DataFrame
                              jq.query(jq.finance.STK_EXCHANGE_TRADE_INFO)       # type: sqlalchemy.orm.query.Query
                              .filter(jq.finance.STK_EXCHANGE_TRADE_INFO.date==date).limit(30)
                              )
    print(pdf)
    return pdf


def getMarketYesterday():
	#today = time.strftime('%Y-%m-%d', time.localtime(time.time()))
	today = datetime.date.today()
	oneday = datetime.timedelta(days=1)
	yesterday = (today - oneday).strftime('%Y-%m-%d')
	print ("yesterday is:", yesterday)
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
    #df =




#-------------------------------- 获取成分股 -----------------------------------------------------------
# 获取现在上证50列表
def getSh50():
    return jq.get_index_stocks('000016.XSHG')


def getSh180():  # 上证180
    return jq.get_index_stocks('000010.XSHG')


def getHs300():  # 沪深300
    return jq.get_index_stocks('000300.XSHG')


def getZz500():  # 中证500
    return jq.get_index_stocks('000905.XSHG')







if __name__ == "__main__":
	print ("User Name Login Result:", Login_JQ("18602122079", "calcaapi"))

	print ("Try to get 2019.9.13 market data:", getMarketMap("2004-01-10") )        #2003-01-10   2019-09-11

	days = getTradeDays(dateStr="2019-09-11", num_prev = "10")
	print ("Try to get 2019.9.06 valid days", days )

	#print ("Try to get Mtss_total on 2019.10.11: ", getMtss("2019-10-10"))
	#print ("Try to get Current market:", getMarketYesterday() )
	#print ("Try to get Current PE:", getCurPE() )
