/*********************************************************
/* top file of mix model rate% & index position
/* ----
/*  License: BSD
/* ----
/* v0.1  . --- 2019.1.15
/*********************************************************/

package mipos



import (
        "pa"
        //"pb"
        //"pc"
        . "cmn"
        "fmt"
        )



/***********************************************************************
/* mipos* is "估值函数"
/*
/***********************************************************************/
// train model param according history data
func Tm_mipos(policy string)(o T_Mipos){

	switch policy{
	case "A": bot_trunk, bot_punch, top_crz, eqpo = MiposGm_Pa(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	case "B": bot_trunk, bot_punch, top_crz, eqpo = MiposGm_Pb(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	case "C": bot_trunk, bot_punch, top_crz, eqpo = MiposGm_Pc(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	default: Log.Panicln("Critical Error: <Mipos_Model> Policy not correct.Input is:", policy)
}


// Estimate today market.
// policy is defined in define file
func Mipos(policy string)(bi, ti int, mix_cw float64){
	switch policy{
        case "A": bi, ti, mix_cw = Mipos_Pa(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	case "B": bi, ti, mix_cw = MiposEf_Pb(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	case "C": bi, ti, mix_cw = MiposEf_Pc(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	default: Log.Panicln("Critical Error: <Mipos> Policy not correct.Input is:", policy)
}



func Mipos_Pa()(crz, hot, puc, rlx []float64, pos float64){
	// step 1: get c.h.r.p values from eig
	crz, hot, puc, rlx, eigMin, eigMax := Chrp()

	// step 2: get current market
	dicMkt  := GetCurMarket()
	// step 3: eval current market
	pos := EvalCurPos(eigMin, eigMax, dicMkt )
	return
}

