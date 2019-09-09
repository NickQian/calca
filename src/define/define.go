/**************************************************************************
/* define types/structs
/* ----
/*  License: BSD
/* ----
/*  0.01 define structs & types - 2019.1.15 - Nick cKing
/**************************************************************************/

package define



//------------------------------- res --------------------------------------

type CalRes struct{
        Bi        int
        Ti        int
	Mix_cw    [2]int
	Scan_res  [9]int
        Auto_res  float64
}


type SimRes struct{
        CurState        string
        CurValue        int32
        CurStock     [3]int32
        CurStockVol  [3]int32
        CurDate         string      //time.Date
}


//-------------------------- A stock params --------------------------------
type T_A struct{
        Cmv   T_cmv                // circulation market value
        Pe    T_pe
        Pb    T_pb
        Mtsr  T_mtsr               // mtss/cmv ratio.  only valid from 2014-05-05
        Volr  T_volr               // (trade vol)/cmv ratio
        Tnr   T_tnr                // daily
}


type T_cmv struct{
        Cmv_total float64
        Cmv_sh float64
        Cmv_sz float64
        Cmv_tim float64
        Cmv_gem float64
}


type T_pe struct{
        Pe_total  float64
        Pe_hs300 float64          // 沪深300
        Pe_sh50 float64
        Pe_sh float64
        Pe_sz float64
        Pe_tim float64            // 科创板
        Pe_gem float64
}


type T_pb struct{
        Pb_total  float64
        Pb_hs300 float64          // 沪深300
        Pb_sh50 float64
        Pb_sh float64
        Pb_sz float64
}


type T_mtsr struct{
        Mtsr_total float64
        Mtsr_sh float64
        Mtsr_sz float64
}


type T_volr struct{
        Volr float64
}


type T_tnr struct{
        Tnr float64
}



//-------------------- configuration parameters -------------------------------
const(
        // JQ == JoinQuant) | UQ == Uqer.io | RQ == RiceQuant | BQ == BigQuant | tushare
        QIF_VENDOR = "JQ"
)



//--------------------- adjustable parameters --------------------------------
const (
        PRE_SMP_NUM         = 10
        DATA_SMOOTH_METHOD  = "LSE"      // least square method, MLE, Good-Turing
)


//-------------------- don't need modification -------------------------------
const (
        TIME_LAYOUT_STR   = "2006-01-02 15:04:05"
        TIME_LAYOUT_SHORT = "2006-01-02"
)

const (
        FN_RES_CALC  = "../data/res_calc.json"
        Fn_RES_RRUN  = "../data/res_rrun.json"
        FN_BOT_DATE  = "../data/bot_date"
        FN_BOT_DATA  = "../data/bot_data"
)
