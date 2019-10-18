/*********************************************************
/* test cmnfcd.go
/* ----
/*  License: BSD
/* ----
/* v0.01  init version --- 2019.7.12
/*********************************************************/

package cmn

import (
         "testing"
         dfn "define"
	"github.com/stretchr/testify/assert"
)



func TestOperateTime(t *testing.T){
        res := OperateTime()
        t.Logf(" %t ", res)
}


func TestReadBotDate(t *testing.T){
        //date := ReadBotDate(dfn.FN_BOT_PUC_DATE)
	date := ReadBotDate(dfn.FN_BOT_RLX_DATE)

        if len(date) == 0{
                t.Logf("Err_Testing: <Get_Bot_Date>: len(d) is  %d ", len(date) )
                t.Errorf("Err_Testing: <Get_Bot_Date>: len(d) is  %d ", len(date))
        }else{
                t.Logf(" date: %s ", date)
        }

        if cap(date) == 0{
                t.Errorf("Err: <TestReadBotDate>: cap(d) is  %d ", cap(date))
        }
}



func TestGetBotWindow(t *testing.T){
        bw := GetBotWindow("2019-01-02")
        if len(bw) == 0{
        	t.Error("Error: <GetBotWindow>: bot window len is 0. Maybe internet access problem? ")
        }else{
	        t.Logf("Window: %v ", bw)
	}
}


func TestGetBotWindow_raw(t *testing.T){
        bw := GetBotWindow_raw("2019-01-02", 10)
        t.Logf("Window: %v ", bw)
}



func TestGetBotsDate(t *testing.T){
        botsDate := GetBotsDate(dfn.FN_BOT_RLX_DATE)
        t.Logf("<TestGetBotsDate>: %v ", botsDate)
}


func TestGetBotsData(t *testing.T){
	var a = make([]dfn.T_A, 50)
	suc := GetBotsData(dfn.FN_BOT_PUC_DATE, a)
	t.Logf("<TestGetBotsData>: %v ", suc)
	assert.True(t, suc)
	t.Log("<TestGetBotsData> ;;;assert passed. A: ", a)
}


func TestProcBotsData(t *testing.T){
	var a_punch, a_relax = make([]dfn.T_A, 100), make([]dfn.T_A, 100)
	GetBotsData(dfn.FN_BOT_PUC_DATE, a_punch)
        GetBotsData(dfn.FN_BOT_RLX_DATE, a_relax)

        dicBot := ProcBotsData(a_punch, a_relax)
	t.Log("<TestProcBotsData> result: dicBot: ", dicBot)
}
