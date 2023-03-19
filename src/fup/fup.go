/****************************************************************************************
/* fetch data from qif, update model parameter("calca" to the updating current status )
/* collect history data/characters.
/* ----
/*  License: BSD
/* ----
/* v0.1  init. --- 2019.10.21
/****************************************************************************************/



package fup


import (
	"cmn"
	"fmt"
	dfn "define"
	"qif"
//	"pa"
)



// fetch data from qif; update model parameter
func Fup(allK bool)(suc bool, err error){
	// get last valid trade day
	RecentTradeDays := qif.GetTradeDays(cmn.TodayStr, dfn.PRE_SMP_NUM)
	todayIsValid, err := cmn.Is_in(RecentTradeDays, cmn.TodayStr)

	// fetch data from qif
	if todayIsValid{
		if allK{
			//suc, K_sh,K_sz,K_sh300,K_gem,K_star := cmn.PullHisKs_Ix(cmn.TodayStr)
			cmn.PullHisKs_Ix(cmn.TodayStr)
		}else{
			cmn.PullTodayKs_Ix()
		}
	}else{
		fmt.Printf("## Warining:<Fup> today is not valid trade day. today: %v, recent trade days: %v \n", cmn.TodayStr, RecentTradeDays )
	}

	// calculate mipos
	//pa.Mipos("A");

	return
}


func UpdateCW_Model()(bool){
        //modelCfg := Mipos_Model()
	return true
}




// distill CURRENT event character. manually usage
func DistilEvt(fn_date string) bool{
	eggs, _, suc := cmn.GetBtsData(dfn.FN_DATE_BOT_PUC, dfn.FN_EVT_REC_DAT, dfn.FN_EVT_AVG_DAT )
	fmt.Print("info: <DistilEvt> : eggs: ", eggs)
	return suc
}


// see if there's new b/t event
func HasBtEvent()(botEvt bool, topEvt bool){
	//rundata, _ := cmn.ReadResTrd(dfn.FN_RES_TRD)
	lastOp, _ := cmn.ReadLastOp(dfn.FN_TRD_LAST_OP)


	/*Bot_Rlx := cmn.ReadBtDate(dfn.FN_DATE_BOT_RLX)
	lastBotRlxDate := Bot_Rlx[0]
	Bot_Puc := cmn.ReadBtDate(dfn.FN_DATE_BOT_PUC)
	lastBotPucDate := Bot_Puc[0]

	Top_Hot := cmn.ReadBtDate(dfn.FN_DATE_TOP_HOT)
	lastTopHotDate := Top_Hot[0]
	Top_Crz := cmn.ReadBtDate(dfn.FN_DATE_TOP_CRZ)
	lastTopCrzDate := Top_Crz[0]

	if lastOp.Date == lastBotRlxDate || rundata.LastBotDate == lastBotPucDate{
		return true, false
	}
	if rundata.LastTopDate == lastTopHotDate || rundata.LastTopDate == lastTopCrzDate{
                return false, true
        } */

	if lastOp.Op_evt_type == dfn.BOT_PUC || lastOp.Op_evt_type == dfn.BOT_RLX{
		return true, false
	}
	

        return false, false
}
