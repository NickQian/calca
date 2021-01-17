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


ts.set_token('your token here')


# daily
def GetK

df = pro.index_daily(ts_code='399300.SZ')

#或者按日期取

df = pro.index_daily(ts_code='399300.SZ', start_date='20180101', end_date='20181010')
