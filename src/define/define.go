/********************************************************************************************
/* define types/structs
/* ----
/*  License: BSD
/* ----
/*  0.01 define structs & types - 2019.1.15 - Nick cKing
/********************************************************************************************/

package define



//------------------------------------- res --------------------------------------------------

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


//-------------------------- A stock params --------------------------------------------------
type T_A struct{
        Pe    T_pe
        Volr  T_volr               // (trade vol)/Cmc ratio
        Tnr   T_tnr                // daily
        Mtsr  T_mtsr               // mtss/Cmc ratio.  only valid from 2014-05-05
        Pb    T_pb
        Cmc   T_cmc                // circulation market value
}


type T_cmc struct{
        Cmc_total float64
        Cmc_sh    float64
        Cmc_sz    float64
        Cmc_tim   float64
        Cmc_gem   float64
}


type T_pe struct{
        Pe_total  float64
        Pe_sh     float64
        Pe_sz     float64
        Pe_gem    float64
        Pe_tim    float64            // 科创板
        Pe_hs300  float64          // 沪深300
        Pe_sh50   float64
}


type T_pb struct{
        Pb_total  float64
        Pb_hs300  float64          // 沪深300
        Pb_sh50   float64
        Pb_sh     float64
        Pb_sz     float64
}


type T_mtsr struct{
        Mtsr_total float64
        Mtsr_sh    float64
        Mtsr_sz    float64
}


type T_volr struct{
	Volr_total float64
        Volr_gem   float64
        Volr_sh    float64
        //Volr_sha   float64
        Volr_sz    float64
        //Volr_szm   float64
}


type T_tnr struct{
	Tnr_total  float64
        Tnr_sh     float64
        //Tnr_sha    float64
        Tnr_sz     float64
        //Tnr_szm    float64
        Tnr_gem    float64
}


//------------------------ stock A calca parameters ------------------------------------------



//------------------------ configuration parameters ------------------------------------------
const(
        // JQ == JoinQuant) | UQ == Uqer.io | RQ == RiceQuant | BQ == BigQuant | tushare
        QIF_VENDOR = "JQ"
)



//----------------------- adjustable parameters ----------------------------------------------
const(
	WIN_SMP_SIZE        = 12         // every window max sample size, some data may lost
        PRE_SMP_NUM         = 5		 // effective sample data number
        DATA_SMOOTH_METHOD  = "LSE"      // least square method, MLE, Good-Turing
)

const(
	QIF_ACCESS_INTVL    = 200        // ms
)
//----------------------- don't need modification --------------------------------------------
const(
        TIME_LAYOUT_STR   = "2006-01-02 15:04:05"
        TIME_LAYOUT_SHORT = "2006-01-02"
)


const(
        FN_RES_CALC        = "../data/res_calc.json"
        Fn_RES_RRUN        = "../data/res_rrun.json"
        FN_BOT_PUC_DATE    = "../data/botpunch_date"
        FN_BOT_RLX_DATE    = "../data/botrelax_date"
        FN_BOT_DATA        = "../data/botdataRec"
)
