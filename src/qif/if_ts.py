#!/usr/bin/env python
# -*- coding:utf-8 -*-
""" Tushare的python接口.   https://waditu.com/
#  ----
#  License: BSD
#  ----
#  0.1: init version: TS数据历史：从2004年1月开始提供 - 2020.9 - Nick cKing
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
4   000300.SH @ (hs300) :   valid from 2005
5   000905.SH   (中证500)
6   399001.SZ * (sz 001)
7   399005.SZ * (中小板指)
8   399006.SZ * (创业板指): valid from 2010
9   399016.SZ   (深证创新)
10  399300.SZ   (sz hs300)
11  399905.SZ   (中证500 )
"""
#------------------------------------------------
# eg. date='20181018'
def getMarketMap(day):
        df = getMarket(day)

	pe_sh,  pe_hs300,  pe_szm,  pe_smb,  pe_gem  = None, None, None, None, None
	tnr_sh, tnr_hs300, tnr_szm, tnr_smb, tnr_gem = None, None, None, None, None
        pb_sh,  pb_hs300,  pb_szm,  pb_smb,  pb_gem  = None, None, None, None, None
        cmc_sh, cmc_hs300, cmc_szm, cmc_smb, cmc_gem = None, None, None, None, None

	if df.empty:
                print("pyInfo: Err: <getMarket> result is empty! ")
                return {}
        else:
	    try:
                pe_sh    = (df.loc[df['ts_code']=='000001.SH', 'pe']).values.tolist()[0]
                tnr_sh   = (df.loc[df['ts_code']=='000001.SH', 'turnover_rate_f']).values.tolist()[0]
		pb_sh    = (df.loc[df['ts_code']=='000001.SH', 'pb']).values.tolist()[0]
		cmc_sh   = (df.loc[df['ts_code']=='000001.SH', 'float_mv']).values.tolist()[0]   #cmc: circulating_market_capacity


		pe_szm    = (df.loc[df['ts_code']=='399001.SZ', 'pe']).values.tolist()[0]
		tnr_szm   = (df.loc[df['ts_code']=='399001.SZ', 'turnover_rate_f']).values.tolist()[0]
                pb_szm    = (df.loc[df['ts_code']=='399001.SZ', 'pb']).values.tolist()[0]
                cmc_szm   = (df.loc[df['ts_code']=='399001.SZ', 'float_mv']).values.tolist()[0]

		pe_hs300 = (df.loc[df['ts_code']=='000300.SH', 'pe']).values.tolist()[0]
                tnr_hs300= (df.loc[df['ts_code']=='000300.SH', 'turnover_rate_f']).values.tolist()[0]
                pb_hs300 = (df.loc[df['ts_code']=='000300.SH', 'pb']).values.tolist()[0]
	        cmc_hs300= (df.loc[df['ts_code']=='000300.SH', 'float_mv']).values.tolist()[0]

                pe_smb   = (df.loc[df['ts_code']=='399005.SZ', 'pe']).values.tolist()[0]
                tnr_smb  = (df.loc[df['ts_code']=='399005.SZ', 'turnover_rate_f']).values.tolist()[0]
                pb_smb   = (df.loc[df['ts_code']=='399005.SZ', 'pb']).values.tolist()[0]
                cmc_smb  = (df.loc[df['ts_code']=='399005.SZ', 'float_mv']).values.tolist()[0]

                pe_gem   = (df.loc[df['ts_code']=='399006.SZ', 'pe']).values.tolist()[0]
		tnr_gem  = (df.loc[df['ts_code']=='399006.SZ', 'turnover_rate_f']).values.tolist()[0]
                pb_gem   = (df.loc[df['ts_code']=='399006.SZ', 'pb']).values.tolist()[0]
		cmc_gem  = (df.loc[df['ts_code']=='399006.SZ', 'float_mv']).values.tolist()[0]

                #vol_sh  = pdf[2:3]['money'][2]
                #mtss_sh, mtss_sz = getMtss(day)

	    except (IndexError, UnboundLocalError), err:
	    	print("@Python Warning: day to query may has no hs300<000300, from year 2005.1> or gem<399006, from year 2010>. Err msg:", err)
		if pe_gem == None and pe_hs300 == None and pe_smb == None:
			return  { 'pe_sh' :pe_sh,  'pe_szm' :pe_szm,            # minimal (1)
                        	  'pb_sh' :pb_sh,  'pb_szm' :pb_szm,            # minimal (2)
                          	  'tnr_sh':tnr_sh, 'tnr_szm':tnr_szm,           # minimal (3)
				  'cmc_sh':cmc_sh, 'cmc_szm':cmc_szm,
                        	}
                elif pe_gem == None and pe_smb == None:
                	return  { 'pe_sh' :pe_sh,  'pe_szm' :pe_szm, 'pe_hs300' : pe_hs300,
                                  'pb_sh' :pb_sh,  'pb_szm' :pb_szm, 'pb_hs300' : pb_hs300,
                                  'tnr_sh':tnr_sh, 'tnr_szm':tnr_szm,'tnr_hs300':tnr_hs300,
				  'cmc_sh':cmc_sh, 'cmc_szm':cmc_szm, 'cmc_hs300':cmc_hs300,
				}
		else :     #pe_gem == None:
                        return  {  'pe_sh': pe_sh,  'pe_szm': pe_szm,  'pe_hs300': pe_hs300,  'pe_smb': pe_smb,
                                   'pb_sh': pb_sh,  'pb_szm': pb_szm,  'pb_hs300': pb_hs300,  'pb_smb': pb_smb,
                                  'tnr_sh':tnr_sh, 'tnr_szm':tnr_szm, 'tnr_hs300':tnr_hs300, 'tnr_smb':tnr_smb,
                                  'cmc_sh':cmc_sh, 'cmc_szm':cmc_szm, 'cmc_hs300':cmc_hs300, 'cmc_smb':cmc_smb,
                                }
	    else:
                return  {  'pe_sh' :pe_sh,  'pe_szm' :pe_szm,  'pe_hs300': pe_hs300,  'pe_smb': pe_smb, 'pe_gem': pe_gem,
                           'pb_sh' :pb_sh,  'pb_szm' :pb_szm,  'pb_hs300': pb_hs300,  'pb_smb': pb_smb, 'pb_gem': pb_gem,
                          'tnr_sh':tnr_sh, 'tnr_szm':tnr_szm, 'tnr_hs300':tnr_hs300, 'tnr_smb':tnr_smb,'tnr_gem':tnr_gem,
			  'cmc_sh':cmc_sh, 'cmc_szm':cmc_szm, 'cmc_hs300':cmc_hs300, 'cmc_smb':cmc_smb,'cmc_gem':cmc_gem,
                         # 'vol_sh':vol_sh,'vol_szm':vol_szm,'vol_sz':vol_sz,'vol_gem':vol_gem,
                         # 'mtss_sh':mtss_sh,'mtss_sz':mtss_sz,
                        }



# eg. date='20181018'
# ts: only valid from 20040105??
def getMarket(date):
	df = pro.index_dailybasic(trade_date= date)               # , fields='ts_code,trade_date,turnover_rate,pe')
	print("# date, df:", date, df)
	return df


#------------------------ Trade days -----------------------------------
# input eg: (2014-06-13, 10)
def getTradeDays(end_date_str, num_prev):
        validays   = []
        numpre     = int(num_prev) + 10
        print("numpre:" ,numpre)
        end_Date   = datetime.datetime.strptime(end_date_str, '%Y-%m-%d') #  '%Y-%m-%d')
	delta = datetime.timedelta(days = numpre -1)
	start_Date = end_Date - delta

        days_df = pro.query('trade_cal', start_date=start_Date.strftime('%Y%m%d'), end_date=end_Date.strftime('%Y%m%d'))    # exchange  cal_date  is_open
	for i in range( numpre):
		if (days_df['is_open']).ix[i] == 1:
	        	valid_traday = (days_df['cal_date']).ix[i]                                         # datetime in list
			validays.append(valid_traday.encode("utf-8") )
        #print("### tradays_df: ", days_df )
	#print("org: return: validays[-6:-1]" , validays, validays[0:int(num_prev)], validays[-6:-1] )
        return validays[-int(num_prev) : ]        # ??? validays[-int(num_prev)-1 : -1 ]






if __name__ ==  "__main__":
	Login_TS("abc", "ddd")
	#kline = getKline("000001.SH","20190103", "20190120")
	#print ("Try to get market.. type_of_element, data:",  type(kline[0]), kline )

	print("---testing <getMarketMap>--- \n")
	print( getMarketMap('20150522') )   #20050607         # 20040105 --> 20190104  now: 20210210
	print("--- testing <getTradeDays>--- \n")
	print( getTradeDays('2014-06-13', '5')  )             # date golang format
