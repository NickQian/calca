#!/usr/bin/env python
# -*- coding:utf-8 -*-
""" Tushare(聚宽)的JQData的python接口.   www.joinquant.com/data
#  ----
#  License: BSD
#  ----
#  0.1: init version - 2020.9 - Nick cKing
"""

import datetime
import tushare as ts


pro = ts.pro_api()

#--- login ---
# usr & password not used
def Login_TS(usr, password):
	ts.set_token('7aaef57854cac90bf596075e9cd60d0c256b62477e312aa509069fcd')
	print("@:Python Info: tushare pro version:", ts.__version__)
	return  True, True, True


# ------------- kline --------------
# daily, eg: '000001.SH', '20181010'
def getKlineNum(index, startDay, endDay):
	df = pro.index_daily(ts_code=index, start_date= startDay, end_date=endDay)
	print ("@:Python Info: Getting Kline... index, startDay, endDay", index, startDay, endDay)
	kclose = df['close'].tolist()
	kclose.reverse()
	return kclose


# turn the number list in to String
def getKline(index, startDay, endDay):
	k_num = getKlineNum(index, startDay, endDay)
	return ['{:.2f}'.format(x) for x in k_num]




#-------- Market property ------------------------
"""ts_code trade_date total_mv(总市值) float_mv(流通市值）total_share float_share free_share turnover_rate turnover_rate_f pe  pe_ttm  pb
0   000001.SH * 20190104  3.219598e+13  2.338122e+13  4.500358e+12  ...           0.51             1.32  12.04   11.08  1.25
1   000005.SH   20190104  1.481860e+12  7.004018e+11  1.738149e+11  ...           1.75             2.52  14.52   13.70  1.34
2   000006.SH   地产指数[000006]实时行情_东方财富
3   000016.SH   (sh50)
4   000300.SH * (hs300)
5   000905.SH   (zz500)
6   399001.SZ * (sz 001)
7   399005.SZ   (中小板指)
8   399006.SZ * (创业板指)
9   399016.SZ   (深证创新)
10  399300.SZ   (sz hs300)
11  399905.SZ   (中证500 )
"""
#------------------------------------------------
# eg. date='20181018'
def getMarketMap(day):
        df = getMarket(day)
        #print ("----pdf is:---", pdf)
        if df.empty:
                print("pyInfo: Err: <getMarket> result is empty! ")
                return {}
        else:
                pe_sh    = (df.loc[df['ts_code']=='000001.SH', 'pe_ttm']).ix[0] #[0] is just the row index of returned data series
		pe_sh300 = (df.loc[df['ts_code']=='000300.SH', 'pe_ttm']).ix[4]
		pe_sz    = (df.loc[df['ts_code']=='399001.SZ', 'pe_ttm']).ix[6]
                pe_gem   = (df.loc[df['ts_code']=='399006.SZ', 'pe_ttm']).ix[8]

                tnr_sh    = (df.loc[df['ts_code']=='000001.SH', 'turnover_rate_f']).ix[0]
                tnr_sh300 = (df.loc[df['ts_code']=='000300.SH', 'turnover_rate_f']).ix[4]
		tnr_sz    = (df.loc[df['ts_code']=='399001.SZ', 'turnover_rate_f']).ix[6]
		tnr_gem   = (df.loc[df['ts_code']=='399006.SZ', 'turnover_rate_f']).ix[8]

		pb_sh    = (df.loc[df['ts_code']=='000001.SH', 'pb']).ix[0] #[0] is just the row index of returned data series
                pb_sh300 = (df.loc[df['ts_code']=='000300.SH', 'pb']).ix[4]
                pb_sz    = (df.loc[df['ts_code']=='399001.SZ', 'pb']).ix[6]
                pb_gem   = (df.loc[df['ts_code']=='399006.SZ', 'pb']).ix[8]

		print ("###: pe:",  pe_sh,  pe_sh300,  pe_sz,  pe_gem)
		print ("###: tnr:", tnr_sh, tnr_sh300, tnr_sz, tnr_gem)
		print ("###: pb:",  pb_sh,  pb_sh300,  pb_sz,  pb_gem)

                #vol_sh  = pdf[2:3]['money'][2]
                #mtss_sh, mtss_sz = getMtss(day)
                #cmc_sha = pdf[0:1]['circulating_market_cap'][0]

                return  { 'pe_sh' :pe_sh,  'pe_sh300' :pe_sh300,   'pe_sz' :pe_sz,  'pe_gem' :pe_gem,
                          'tnr_sh':tnr_sh, 'tnr_sh300':tnr_sh300,  'tnr_sz':tnr_sz, 'tnr_gem':tnr_gem,
                          'pb_sh' :pe_sh,  'pb_sh300' :pe_sh300,   'pb_sz' :pe_sz,  'pb_gem' :pe_gem,
			 # 'cmc_sh':cmc_sh,'cmc_szm':cmc_szm,'cmc_sz':cmc_sz,'cmc_gem':cmc_gem,
                         # 'vol_sh':vol_sh,'vol_szm':vol_szm,'vol_sz':vol_sz,'vol_gem':vol_gem,
                         # 'mtss_sh':mtss_sh,'mtss_sz':mtss_sz,
                        }




# eg. date='20181018'
def getMarket(date):
	df = pro.index_dailybasic(trade_date= date)               # , fields='ts_code,trade_date,turnover_rate,pe')
	print("###df:", df)
	return df


#------------------------ Trade days -----------------------------------
def getTradeDays(end_date_str, num_prev):
        validays   = []
        numpre     = int(num_prev)
        end_Date   = datetime.datetime.strptime(end_date_str, '%Y-%m-%d')
	delta = datetime.timedelta(days = numpre)
	start_Date = end_Date - delta
	#exchange  cal_date  is_open

        tradays_df = pro.query('trade_cal', start_date=start_Date.strftime('%Y%m%d'), end_date=end_Date.strftime('%Y%m%d'))    # exchange  cal_date  is_open
        for day_df in tradays_df:
		if day_df['is_open'] is true:
	        	list_tradays = day_df['cal_date'].tolist()                                         # datetime in list
                	validays.append(day.strftime("%Y-%m-%d") )
        return validays






if __name__ ==  "__main__":
	Login_TS("abc", "ddd")
	kline = getKline("000001.SH","20190103", "20190120")
	print ("Try to get market.. type_of_element, data:",  type(kline[0]), kline )

	print("---testing <getMarketMap>--- \n")
	print( getMarketMap('20210210') )            #20190104
	print("--- testing <getTradeDays>--- \n")
	print( getTradeDays('20190105', '10')  )
