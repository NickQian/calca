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



func TestReadCalRes(t *testing.T){
	calRes, err := ReadCalRes(dfn.FN_RES_CALC)
	if err != nil{
		t.Errorf("Err during Testing:<ReadCalRes>. Err is: %v", err)
	}else{
		t.Logf("<ReadCalRes> result Bi: %v, Ti: %v, Scan_res:%v ", calRes.Bi, calRes.Ti, calRes.ScanRes)
	}
}


func TestReadBtDate(t *testing.T){
        //date := ReadBotDate(dfn.FN_BOT_PUC_DATE)
	date := ReadBtDate(dfn.FN_DATE_BOT_RLX)
        if len(date) == 0{
                t.Logf("Err during Testing: <Get_Bot_Date>: len(d) is  %d ", len(date) )
                t.Errorf("Err during Testing: <Get_Bot_Date>: len(d) is  %d ", len(date))
        }else{
                t.Logf(" date: %s ", date)
        }

        if cap(date) == 0{
                t.Errorf("Err: <TestReadBotDate>: cap(d) is  %d ", cap(date))
        }
}


func TestGetBtWindow(t *testing.T){
        bw := GetBtWindow("2019-01-02")
        if len(bw) == 0{
        	t.Error("Error: <GetBotWindow>: bot window len is 0. Maybe internet access problem? ")
        }else{
	        t.Logf("Window: %v ", bw)
	}
}


func TestGetBtWindow_raw(t *testing.T){
        bw := GetBtWindow_raw("2019-01-02", 10)
        t.Logf("Window: %v ", bw)
}



func TestGetBtsDate(t *testing.T){
        botsDate := GetBtsDate(dfn.FN_DATE_BOT_RLX )
        t.Logf("<TestGetBotsDate>: %v ", botsDate)
}


func TestGetBtsData(t *testing.T){
	//var a = make([]dfn.T_A, 50)
	eggMaps, _, suc := GetBtsData(dfn.FN_DATE_BOT_PUC, dfn.FN_EVT_REC_DAT, dfn.FN_EVT_AVG_DAT)
	t.Logf("t.Logf <TestGetBtsData> result: %v ", suc)
	assert.True(t, suc)
	t.Logf("t.Log <TestGetBtsData>. eggMap: %v ", eggMaps)
}


func TestGetEigDm(t *testing.T){

	dmEigRlx, dmEigPuc, dmEigHot, dmEigCrz, _ := GetEigDm(dfn.FN_EGG_BOT_RLX, dfn.FN_EGG_BOT_PUC, dfn.FN_EGG_TOP_CRZ, dfn.FN_EGG_TOP_HOT)
	t.Log("<TestGetEigDm> result: ", dmEigRlx, dmEigPuc, dmEigHot, dmEigCrz)
}



func TestReadResTrd(t *testing.T){
	resTrd, err := ReadResTrd(dfn.FN_RES_TRD)
	if err != nil{
		t.Errorf("Err during testing:<ReadRrunRes>. Err is: %v", err)
	}else{
		t.Logf(" @resTrd: %v", resTrd )
	}
}


func TestReadLastOp(t *testing.T){
	lastOpDat, err := ReadLastOp(dfn.FN_TRD_LAST_OP)
	if err != nil{
		t.Errorf("Err during Testing: <ReadRunData>: Err is  %v ", err)
	}else{
		t.Logf(" @ lastOpDat is: %v", lastOpDat )
	}
}



