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


ts.set_token('7aaef57854cac90bf596075e9cd60d0c256b62477e312aa509069fcd')
pro = ts.pro_api()
print("tushare pro version:", ts.__version__)

# daily, eg: '000001.SH', '20181010'
def GetK(index, startDay, endDay):
	df = pro.index_daily(ts_code=index, start_date= startDay, end_date=endDay)
	#print ("@:python info, df:", df)
	kclose = df['close'].tolist()
	return kclose



if __name__ ==  "__main__":
	print ("Try to get market data:", GetK("000001.SH","20190911", "20190916") )
