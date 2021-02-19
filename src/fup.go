/**************************************************************
/* fetch data from qif, update model parameter("calca" to the updating current status )
/* collect history data/characters.
/* ----
/*  License: BSD
/* ----
/* v0.1  . --- 2019.10.21
/**************************************************************/

package main


import (
	. "cmn"
	"fmt"
	dfn "define"
)



func UpdateCW_Model()(bool){
        //modelCfg := Mipos_Model()
	return true
}


// distill current event character
func DistilEvt(fn_date string) bool{
	eggs, _, suc := GetBtsData(fn_date)
	fmt.Print("info: <DistilEvt> : eggs: ", eggs)
	return suc
}


// see if there's new b/t event
func HasBtEvent()(botEvt bool, topEvt bool){
	rundata, _ := ReadRunData(dfn.FN_RUN_DATA)

	Bot_Rlx := ReadBtDate(dfn.FN_DATE_BOT_RLX)
	lastBotRlxDate := Bot_Rlx[0]
	Bot_Puc := ReadBtDate(dfn.FN_DATE_BOT_PUC)
	lastBotPucDate := Bot_Puc[0]

	Top_Hot := ReadBtDate(dfn.FN_DATE_TOP_HOT)
	lastTopHotDate := Top_Hot[0]
	Top_Crz := ReadBtDate(dfn.FN_DATE_TOP_CRZ)
	lastTopCrzDate := Top_Crz[0]

	if rundata.LastBotDate == lastBotRlxDate || rundata.LastBotDate == lastBotPucDate{
		return true, false
	}
	if rundata.LastTopDate == lastTopHotDate || rundata.LastTopDate == lastTopCrzDate{
                return false, true
        }

        return false, false
}
