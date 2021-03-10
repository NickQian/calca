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
	Evt_Tag    string  `json:"evt_tag"`         // date, "BOT_RLX"/"BOT_PUC"/"TOP_HOT"/"TOP_CRZ"

        Cmc        T_cmc   `jason:"CMC"`            // circulation market value
        Pe         T_pe	   `jason:"PE"`	   	    // Market: PE
        Pb         T_pb	   `jason:"PB"` 	    // Market: PB
        Tnr        T_tnr   `jason:"TNR"`            // Emo   : Tnr(daily)

        //Mtsr       T_mtsr  `jason:"MTSR"`         // Emo   : mtss/Cmc ratio.  only valid from 2014-05-05
        //Volr       T_volr  `jason:"VOLR"`         // Emo   : (trade vol)/Cmc ratio
}


type T_cmc struct{
        Cmc_total  float64 `jason:"Cmc_total"`
        Cmc_sh     float64 `jason:"Cmc_sh"`
        Cmc_sz     float64 `jason:"Cmc_sz"`         // whole shenzhen market
	Cmc_hs300  float64 `jason:"Cmc_hs300"`
	Cmc_szm    float64 `jason:"Cmc_szm"`        // shenzhen main board. big change after 2015.5.20
	Cmc_smb    float64 `jason:"Cmc_smb"`        // small and medium board
        Cmc_gem    float64 `jason:"Cmc_gem"`        // Growth Enterprise Market
        //Cmc_tim    float64 `jason:"Cmc_tim"`      // Technology Innovation Board
}


type T_pe struct{
        Pe_total   float64 `jason:"Pe_total"`
        Pe_sh      float64 `jason:"Pe_sh"`
        Pe_sz      float64 `jason:"Pe_sz"`
        Pe_hs300   float64 `jason:"Pe_hs300"`         // 沪深300
        Pe_szm     float64 `jason:"Pe_szm"`
        Pe_smb     float64 `jason:"Pe_smb"`
        Pe_gem     float64 `jason:"Pe_gem"`
        //Pe_tim     float64 `jason:"Pe_tim"`           // 科创板
}


type T_pb struct{
        Pb_total   float64 `jason:"Pb_total"`
        Pb_sh      float64 `jason:"Pb_sh"`
        Pb_sz      float64 `jason:"Pb_sz"`
        Pb_hs300   float64 `jason:"Pb_hs300"`         // 沪深300
        Pb_szm     float64 `jason:"Pb_szm"`
        Pb_smb     float64 `jason:"Pb_smb"`
        Pb_gem     float64 `jason:"Pb_gem"`
}


type T_tnr struct{
	Tnr_total  float64 `jason:"Tnr_total"`
        Tnr_sh     float64 `jason:"Tnr_sh"`
        Tnr_sz     float64 `jason:"Tnr_sz"`
        Tnr_hs300  float64 `jason:"Tnr_hs300"`
        Tnr_szm    float64 `jason:"Tnr_szm"`
        Tnr_smb    float64 `jason:"Tnr_smb"`
        Tnr_gem    float64 `jason:"Tnr_gem"`
}


type T_mtsr struct{
        Mtsr_total float64 `jason:"Mtsr_total"`
        Mtsr_sh    float64 `jason:"Mtsr_sh"`
        Mtsr_sz    float64 `jason:"Mtsr_sz"`
        Mtsr_hs300 float64 `jason:"Mtsr_hs300"`
        Mtsr_gem   float64 `jason:"Mtsr_gem"`
}


type T_volr struct{
	Volr_total float64 `jason:"Volr_total"`
        Volr_sh    float64 `jason:"Volr_sh"`
        Volr_sz    float64 `jason:"Volr_sz"`
        Volr_hs300    float64 `jason:"Volr_sz"`
        Volr_gem   float64 `jason:"Volr_gem"`
}

//----------------------- Market factors --------------------------------------
type T_Market struct{
	mktVal		T_MktVal
	trdEmo		T_TrdEmo
	pubPly		T_PubPly
}

type T_MktVal struct{
	pe_total float64
	pb_total float64
}

type T_TrdEmo struct{
	volr_total float64                 // vol / cap
	tnr_total  float64
	mtsr_total float64
}

type T_PubPly struct{
	cp   float64                      // credit policy:tighten or relax
	nbir float64			 // national debt
	ir     float64                    // interest rate
}


//------------------------ mipos data ------------------------------------------
type T_Mipos struct{
	BotRlx   []float64       // -Pi
	BotPuc   []float64       // punch ~  0
	TopHot   []float64       // 0     ~  hot
	TopCrz   []float64       // hot   ~  Pi
	Eqpo       float64       // 0: equilibrium position
	Bi         float64
	Ti         float64
	Cw_pos     float64
	Gnd        float64
	Vcc        float64
}


// ------- run data ------------
type T_Rundata struct{
	LastBotDate  string
	LastTopDate  string
}



//-------------------- ERR messages ----------------------
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
	FFT_FLT_INTEREST_PTS = 20         // 10/15
	FFT_FLT_PAD_LEN      = 0.0        // 0.2 * len
)



//============================= don't need modification ==============================
const(
        TIME_LAYOUT_STR    = "2006-01-02 15:04:05"
        TIME_LAYOUT_SHORT  = "2006-01-02"
)



const(  DATE_HS300_START   = "2005-04-29"
        DATE_MKT_SMB_START = "2006-02-10"
	DATE_MKT_GEM_START = "2010-06-18"
	DATE_2015_0520EVT  = "2015-05-20"
	DATE_MKT_TIM_START = "2019-08-22"
)


const(  RUN_DIR = "/home/nk/calca/src/" )


const(
	FN_RUN_DATA        = RUN_DIR + "data/run_data.json"
        FN_RES_CALC        = RUN_DIR + "data/res_calc.json"
        FN_RES_RRUN        = RUN_DIR + "data/res_rrun.json"
)

const(
        FN_DATE_BOT_PUC    = RUN_DIR + "data/man/botpunch_date"
        FN_DATE_BOT_RLX    = RUN_DIR + "data/man/botrelax_date"
        FN_DATE_TOP_HOT    = RUN_DIR + "data/man/tophot_date"
        FN_DATE_TOP_CRZ    = RUN_DIR + "data/man/topcrz_date"
        FN_EVT_REC_DAT     = RUN_DIR + "data/his/hisEvtRecData.json"
        FN_EVT_AVG_DAT     = RUN_DIR + "data/his/hisEvtAvgData.json"
        FN_EGG_BOT_RLX     = RUN_DIR + "data/proc/botrelax_cha.json"
        FN_EGG_BOT_PUC     = RUN_DIR + "data/proc/botpunch_cha.json"
        FN_EGG_TOP_CRZ     = RUN_DIR + "data/proc/topcrz_cha.json"
        FN_EGG_TOP_HOT     = RUN_DIR + "data/proc/tophot_cha.json"
)
