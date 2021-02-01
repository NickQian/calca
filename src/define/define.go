/********************************************************************************************
/* define types/structs
/* ----
/*  License: BSD
/* ----
/*  0.01 define structs & types - 2019.1.15 - Nick cKing
/********************************************************************************************/

package define

import(
	"time"
)

//------------------------------------- res --------------------------------------------------

type T_CalRes struct{
        Bi              int
        Ti              int
	Mix_cw          [2]int
	Scan_res        [9]int
        Rrun_res        float64
}


type T_SimRes struct{
        CurValue        int
        CurCode         [3]T_CodeInfo
        CurState        string
        CurDate         string      //time.Date
}

type T_CodeInfo struct{
	Code            string
	OrgValue        int
	CurValue        int
}

//-------------------------- A stock params --------------------------------------------------
type T_A struct{
	EventType  string  `json:"evt_type"`    // "BOT_RLX"/"BOT_PUC"/"TOP_HOT"/"TOP_CRZ"
        Pe         T_pe	   `jason:"PE"`		// Market: PE
        Pb         T_pb	   `jason:"PB"` 	// Market: PB
        Volr       T_volr  `jason:"VOLR"`       // Emo   : (trade vol)/Cmc ratio
        Tnr        T_tnr   `jason:"TNR"`        // Emo   : Tnr(daily)
        Mtsr       T_mtsr  `jason:"MTSR"`       // Emo   : mtss/Cmc ratio.  only valid from 2014-05-05
        Cmc        T_cmc   `jason:"CMC"`        // circulation market value
}


type T_cmc struct{
        Cmc_total  float64 `jason:"Cmc_total"`
        Cmc_sh     float64 `jason:"Cmc_sh"`
        Cmc_sz     float64 `jason:"Cmc_sz"`
        Cmc_tim    float64 `jason:"Cmc_tim"`
        Cmc_gem    float64 `jason:"Cmc_gem"`
}


type T_pe struct{
        Pe_total   float64 `jason:"Pe_total"`
        Pe_sh      float64 `jason:"Pe_sh"`
        Pe_sz      float64 `jason:"Pe_sz"`
        Pe_gem     float64 `jason:"Pe_gem"`
        Pe_tim     float64 `jason:"Pe_tim"`           // 科创板
        Pe_hs300   float64 `jason:"Pe_hs300"`         // 沪深300
        Pe_sh50    float64 `jason:"Pe_sh50"`
}


type T_pb struct{
        Pb_total   float64 `jason:"Pb_total"`
        Pb_hs300   float64 `jason:"Pb_hs300"`         // 沪深300
        Pb_sh50    float64 `jason:"Pb_sh50"`
        Pb_sh      float64 `jason:"Pb_sh"`
        Pb_sz      float64 `jason:"Pb_sz"`
}


type T_mtsr struct{
        Mtsr_total float64 `jason:"Mtsr_total"`
        Mtsr_sh    float64 `jason:"Mtsr_sh"`
        Mtsr_sz    float64 `jason:"Mtsr_sz"`
}


type T_volr struct{
	Volr_total float64 `jason:"Volr_total"`
        Volr_gem   float64 `jason:"Volr_gem"`
        Volr_sh    float64 `jason:"Volr_sh"`
        Volr_sz    float64 `jason:"Volr_sz"`
}


type T_tnr struct{
	Tnr_total  float64 `jason:"Tnr_total"`
        Tnr_sh     float64 `jason:"Tnr_sh"`
        Tnr_sz     float64 `jason:"Tnr_sz"`
        Tnr_gem    float64 `jason:"Tnr_gem"`
}


//------------------------ mipos data ------------------------------------------
type T_Mipos struct{
	Bot_trunk  float64       // -Pi
	Bot_punch  float64       // punch ~  0
	Top_hot    float64       // 0     ~  hot
	Top_crzy   float64       // hot   ~  Pi
	Eqpo       float64       // 0
}


type T_Rundata struct{
	LastBotDate  string
	LastTopDate  string
}


type T_ErrCalca struct{
        Code       int
	QifErr     string
        CptErr     string
        IoErr      string
        where      string
        time       time.Time
}



//------------------------ configuration parameters ------------------------------------------
const(
        // JQ == JoinQuant) | UQ == Uqer.io | RQ == RiceQuant | BQ == BigQuant | TS == tushare
        QIF_VENDOR         = "TS"
)



//----------------------- adjustable parameters ----------------------------------------------
const(
	WIN_SMP_SIZE       = 12         // every window max sample size, some data may lost
        PRE_SMP_NUM        = 5		 // effective sample data number
	WIN_SIZE           = PRE_SMP_NUM
        DATA_SMOOTH_METHOD = "LSE"      // least square method, MLE, Good-Turing
)

const(
	QIF_ACCESS_INTVL     = 100        // ms

	FFT_FA_INTEREST_PTS  = 50
	FFT_FLT_INTEREST_PTS = 10         // 10/15
	FFT_FLT_PAD_LEN      = 0.2        // 0.2 * len
)

//----------------------- don't need modification --------------------------------------------
const(
        TIME_LAYOUT_STR    = "2006-01-02 15:04:05"
        TIME_LAYOUT_SHORT  = "2006-01-02"
)


const(
	FN_RUN_DATA        = "../data/run_data.json"
        FN_RES_CALC        = "../data/res_calc.json"
        Fn_RES_RRUN        = "../data/res_rrun.json"
        FN_BOT_PUC_REC_DAT = "../data/his/hisEvtRecData.json"
        FN_BOT_PUC_AVG_DAT = "../data/his/hisEvtAvgData.json"
        FN_BOT_PUC_DATE    = "../data/man/botpunch_date"
        FN_BOT_RLX_DATE    = "../data/man/botrelax_date"
        FN_TOP_HOT_DATE    = "../data/man/tophot_date"
        FN_BOT_CRZ_DATE    = "../data/man/topcrz_date"
)
