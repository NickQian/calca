/**************************************************************************
/* define types/structs
/* ----
/*  License: BSD
/* ----
/*  0.01 define structs & types - 2019.1.15 - Nick cKing
/**************************************************************************/

package cmn



//------------------------------- res --------------------------------------

const (
    FN_RES_CALC  = "../data/res_calc.json"
    Fn_RES_RRUN = "../data/res_rrun.json"
    FN_BOT_DATE  = "../data/bot_date"
    FN_BOT_DATA  = "../data/bot_data"
)

type CalRes struct{
        bi        int
        ti        int
	mix_cw    [2]int
	scan_res  [9]int
        auto_res  float64
}


type SimRes struct{
        curState        string
        curValue        int32
        curStock     [3]int32
        curStockVol  [3]int32
        curDate         string      //time.Date
}


//-------------------------- A stock params --------------------------------
type T_A struct{
        cmv   T_cmv                // circulation market value
        pe    T_pe
        pb    T_pb
        mtsr  T_mtsr               // mtss/cmv ratio.  only valid from 2014-05-05
        volr  T_volr               // (trade vol)/cmv ratio
        tnr   T_tnr                // daily
}

type T_cmv struct{
        cmv_total float64
        cmv_sh float64
        cmv_sz float64
        cmv_tim float64
        cmv_gem float64 
}

type T_pe struct{
        pe_total  float64
        pe_hs300 float64          // 沪深300
        pe_sh50 float64
        pe_sh float64
        pe_sz float64
        pe_tim float64            // 科创板
        pe_gem float64
}

type T_pb struct{
        pb_total  float64
        pb_hs300 float64          // 沪深300
        pb_sh50 float64
        pb_sh float64
        pb_sz float64
}

type T_mtsr struct{
        mtsr_total float64
        mtsr_sh float64
        mtsr_sz float64
}

type T_volr struct{
        volr float64
}

type T_tnr struct{
        tnr float64
}



//-------------------------------- 数据选择 --------------------------------




//----------------------- configuration params -----------------------------
const(
        // JQ == JoinQuant) | UQ == Uqer.io | RQ == RiceQuant | BQ == BigQuant | tushare 
        QIF_VENDOR = "JQ"
)

//wind quantitative interface data
type Qdata struct{
    pe  T_pe
    pb  T_pb
    mtss,  mtsr  float64   //?余额占流通市值比?
}


//---------------------------------------------------------------



