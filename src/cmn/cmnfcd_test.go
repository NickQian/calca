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


// single code
func TestPullTodayPos(t *testing.T){
	yestPos := PullTodayPos("000001.SH", dfn.FN_KLINE_SH)
	t.Logf("yestPos: %v ", yestPos)
}


// today's Klines of indexes(sh/sz/sh300/gem...)
func TestPullTodayKs_Ix(t *testing.T){
	suc := PullTodayKs_Ix()
	t.Logf("<PullTodayKs> result: %v \n", suc)
}

// pull indexs 4 codes history data
func TestPullHisKs_Ix(t *testing.T){
	//_, K := PullHisKs_Ix(YesterdayStr)
	_, k_sh, k_sz, k_sh300, k_gem := PullHisKs_Ix("20230118")
	t.Logf("<TestPullHisKs_Ix> sh_K len: %v, k_sz len: %v, k_sh300 len:%v, k_gem len:%v ", len(k_sh), len(k_sz), len(k_sh300), len(k_gem) )
}


// pull single code history data
func TestPullHisK_Is(t *testing.T){
	startDate := DateStrRmvSlash(dfn.DATE_LAST_BOT2)
	endDate  :=  DateStrRmvSlash(YesterdayStr)

	PullHisIs("300059.SZ", startDate,  endDate,  dfn.FN_KLINE_SUG1B)        // dong cai
	//PullHisIs("600109.SH", startDate,  endDate,  dfn.FN_KLINE_SUG1A)        // guo jin
	//PullHisIs("603986.SH", "20170707", endDate,  dfn.FN_KLINE_SUG2A)        // zhao yi
	//PullHisIs("002049.SZ", startDate,  endDate,  dfn.FN_KLINE_SUG2B)        // ziguang_guowei

	t.Logf("@@ t.Logf: startDate: %v, YesterdayStr: %v \n", startDate, endDate)
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



