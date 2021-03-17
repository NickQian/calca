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
        "cmn"
        "fmt"
	. "define"
	"qif"
	//"plotit"
        )



/***********************************************************************
/* Training model
/*
/***********************************************************************/
// train model param according history data
func Tm_mipos(policy string)(o T_Mipos){
/*
	switch policy{
	case "A": bot_trunk, bot_punch, top_crz, eqpo = MiposGm_Pa(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	case "B": bot_trunk, bot_punch, top_crz, eqpo = MiposGm_Pb(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	case "C": bot_trunk, bot_punch, top_crz, eqpo = MiposGm_Pc(botR, botP[]dfn.T_A, topH, topC[]dfn.T_A)
	default: Log.Panicln("Critical Error: <Mipos_Model> Policy not correct.Input is:", policy)
*/
	return
}


// Estimate today market.
// policy is defined in define file
func Mipos(policy string)( cwRes T_Mipos){
	//var crz, hot, puc, rlx []float64

	switch policy{
        case "A": cwRes = Mipos_Pa( )
	case "B": cwRes = Mipos_Pb( )
	case "C": cwRes = Mipos_Pc( )
	default: cmn.Log.Panicln("Critical Error: <Mipos> Policy not correct.Input is:", policy)
	}
	return
}


// call by "fup" manually. <Mipos_Daily_Pa> will do daily update
func Mipos_Pa()(cwRes T_Mipos){
	// step 1: get c.h.r.p values from eig
	crz, hot, puc, rlx, cha_min, cha_max := pa.Chrp( )

	// step 2: get current market
	days    := qif.GetTradeDays(cmn.TodayStr, LEN_MIPOS_K)

	// step 3: eval current market
	for i, day := range days[:len(days)-1]{     // last day = today; there's no today data
		dicMkt     := cmn.GetMarket(day)
		pos        := pa.EvalPos(cha_min, cha_max, dicMkt )
		cwRes.Poslc = append(cwRes.Poslc, pos)
		cmn.Delay_Qif_Intvl( )
		fmt.Printf("#####<Mipos_Pa> i:%v, day:%v, cwRes.Poslc: %v   \n", i, day, cwRes.Poslc)
	}

	cwRes.BotRlx = rlx
        cwRes.BotPuc = puc
        cwRes.TopHot = hot
        cwRes.TopCrz = crz                  		   // []float64
	cwRes.Pos    = cwRes.Poslc[len(cwRes.Poslc)-1]     // last
	gnd, _      := cmn.GetMinMax_1d(rlx)
	_,  vcc     := cmn.GetMinMax_1d(crz)
        cwRes.Gnd    = gnd         		    	   //   float64
        cwRes.Vcc    = vcc
	halfLen      := (vcc - gnd)/2
	cwRes.Eqpo   = halfLen + gnd
        if cwRes.Pos > cwRes.Eqpo {
		cwRes.Ti = (cwRes.Pos - cwRes.Eqpo)/halfLen
		cwRes.Bi = 0

	}else{
		cwRes.Bi = (cwRes.Eqpo - cwRes.Pos)/halfLen
        	cwRes.Ti =  0
	}

	return
}

func Mipos_Daily_Pa()(float64){

	// get cha_min/max
        _, _, _, _, cha_min, cha_max := pa.Chrp( )

        // step 2: get current market
        dicMkt  := cmn.GetCurMarket()

        // step 3: eval current market
        pos     := pa.EvalPos(cha_min, cha_max, dicMkt )

	return pos
}


//---------------------------- PB ----------------------------------------
func Mipos_Pb( )(cwRes T_Mipos ){
	fmt.Print("Starting <Mipos_Pb> ...")
	return
}



//----------------------------- PC ---------------------------------------
func Mipos_Pc()( cwRes T_Mipos ){
        return
}




//---------------------------------------------------
func miposDataRW()bool{
	return true
}
