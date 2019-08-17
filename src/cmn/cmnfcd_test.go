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
         "fmt"
         )



func TestGet_Bot_Date(t *testing.T){
        cmn.Print("@@1 : FN_BOT_DATE is:", cmn.FN_BOT_DATE)

        date := cmn.Get_Bot_Date(cmn.FN_BOT_DATE)

        if len(date) == 0{
                t.Logf("Err_Testing: <Get_Bot_Date>: len(d) is  %d ", len(date) )
                t.Errorf("Err_Testing: <Get_Bot_Date>: len(d) is  %d ", len(date))
        }else{
                t.Logf("<TestGet_Bot_Date>: date result: ", date)
        }

        if cap(date) == 0{
                t.Errorf("Err_Testing: <Get_Bot_Date>: cap(d) is  %d ", cap(date))
        }
}



