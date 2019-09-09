/*********************************************************
/* test cmnfcd.go
/* ----
/*  License: BSD
/* ----
/* v0.01  init version --- 2019.7.12
/*********************************************************/

package cmn_test

import (
         "testing"
         "cmn"
         "define"
)



func TestGetBotsDate(t *testing.T){
        botsDate := cmn.GetBotsDate()
        t.Logf("<TestGetBotsDate>: %v ", botsDate)
}


func TestGetBotWindow(t *testing.T){
        bw := cmn.GetBotWindow("2019-01-02", 10)
        t.Logf("<TestGetBotWindown>: %v ", bw)
}


func TestOperateTime(t *testing.T){
        res := cmn.OperateTime()
        t.Logf("<OperateTime> %t ", res)
}


func TestReadBotDate(t *testing.T){
        date := cmn.ReadBotDate(define.FN_BOT_DATE)

        if len(date) == 0{
                t.Logf("Err_Testing: <Get_Bot_Date>: len(d) is  %d ", len(date) )
                t.Errorf("Err_Testing: <Get_Bot_Date>: len(d) is  %d ", len(date))
        }else{
                t.Logf("<TestReadBotDate>: date result: %s ", date)
        }

        if cap(date) == 0{
                t.Errorf("Err: <TestReadBotDate>: cap(d) is  %d ", cap(date))
        }
}



