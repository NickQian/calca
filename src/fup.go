/*********************************************************
/* fetch data from qif, update model parameter
/* ----
/*  License: BSD
/* ----
/* v0.1  . --- 2019.10.21
/*********************************************************/

import (
	. "cmn"
	"fmt"
)



func UpdateCW_Model()(){
        modelCfg := Mipos_Model()

}


// calculate today postion & scan
func CalcaToday()(bi, ti int, mix_cw float64){
	//(1) current position
        bot_relax, bot_punch, top_hot, top_crazy = Mipos(policy="A", todayA= &dfn.T_A)
                // mix_cw: model mix of (Casino) - (weighing machine)

	//(2) scan
        sug := Scan()
        if err = WriteRes(res){
        	Log.Println(err)
        }
}


// see if there's new b/t event
func HasBtEvent()(botEvt bool, topEvt bool){
	rundata, _ := ReadRunData(FN_RUN_DATA)

	Bot_Rlx := ReadBtDate(FN_BOT_RLX_DATE)
	lastBotRlxDate := Bot_Rlx[0]
	Bot_Puc := ReadBtDate(FN_BOT_PUC_DATE)
	lastBotPucDate := Bot_Puc[0]

	Top_Hot := ReadBtDate(FN_TOP_HOT_DATE)
	lastTopHotDate := Top_Hot[0]
	Top_Crz := ReadBtDate(FN_TOP_CZR_DATE)
	lastTopCrzDate := Top_Crz[0]

	if rundata.LastBotDate == lastBotRlxDate || rundata.LastBotDate == lastBotPucDate{
		return true, false
	}
	if rundata.LastTopDate == lastTopHotDate || rundata.LastTopDate == lastTopCrzDate{
                return false, true
        }

        return false, false
}
