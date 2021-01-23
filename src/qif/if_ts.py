#!/usr/bin/env python
# -*- coding:utf-8 -*-
"""
# Tushare(聚宽)的JQData的python接口.   www.joinquant.com/data
#  ----
#  License: BSD
#  ----
#  0.1: init version - 2020.9 - Nick cKing
"""

import datetime
import tushare as ts


pro = ts.pro_api()


def Login_TS(usrname, password):
	ts.set_token('7aaef57854cac90bf596075e9cd60d0c256b62477e312aa509069fcd')
	print("@:python info: tushare pro version:", ts.__version__)
	return  True, True, True

# daily, eg: '000001.SH', '20181010'
def getKlineNum(index, startDay, endDay):
	df = pro.index_daily(ts_code=index, start_date= startDay, end_date=endDay)
	print ("@:python info: index, startDay, endDay", index, startDay, endDay)
	kclose = df['close'].tolist()
	kclose.reverse()
	return kclose

# turn the number list in to String
def getKline(index, startDay, endDay):
	k_num = getKlineNum(index, startDay, endDay)
	return ['{:.2f}'.format(x) for x in k_num]



if __name__ ==  "__main__":
	Login_TS("abc", "ddd")
	kline = getKline("000001.SH","20190103", "20210118")
	print ("Try to get market.. type_of_element, data:",  type(kline[0]), kline )
