/*********************************************************
/* common funcs & data & constant
/* ----
/*  License: BSD
/* ----
/* v0.01  init version --- 2019.2.12
/*********************************************************/

package cmn

import (
        "fmt"
        "io/ioutil"
        "time"
        "qif"
        //"strconv"
        "encoding/json"
        "os"
	"reflect"
        "log"
        "errors"
        . "define"
        "strings"
        "bytes"
        )


var Log *log.Logger
var T_Now time.Time
var Today string
var A = new(T_A)               // Inst A parameters


type T_A_BT struct{
	PeMap    map[string]float64
	VolrMap  map[string]float64
	MtsrMap  map[string]float64
	ParaMap  map[string]float64
}


// inst the Bottom parameters container with 0
var BotA = &T_A_BT{PeMap   : make(map[string]float64),
  		   VolrMap : make(map[string]float64),
		   MtsrMap : make(map[string]float64),
		   ParaMap : make(map[string]float64),
}


//type Ftp func(a ...interface{})(n int, err error)
//var  Print Ftp = fmt.Println
var Print func(a ...interface{}) (n int, err error)  = fmt.Println



func init(){
        if logfile, err := os.OpenFile("../run.log",os.O_APPEND|os.O_CREATE, 666); err == nil{
	        Log = log.New(logfile, "", log.Ldate | log.Ltime)

        }else{
                panic("<cmnfd-init> Open logfile error")
        }

	initToday()
        initBotpara(BotA)                // point input
	initLogger()
        //Log.Println("<cmn-init> done. T_Now, Today:", T_Now, Today )
}


func initToday(){
        T_Now = time.Now()
        TodaySlice := strings.SplitAfter(T_Now.Format(TIME_LAYOUT_STR), " ")
        Today = strings.TrimSpace(TodaySlice[0])

        Print("<cmn-init>: Today:", Today )
        Print("<cmn-init> done!", TodaySlice)
}


func initBotpara(p_BotA *T_A_BT) (suc bool) {
	suc = false
        (*p_BotA).ParaMap["SWC_pe"]    = 0.22       // Subjective Weight Correction for entropy weight method
        (*p_BotA).ParaMap["SWC_pb"]    = 0.23
        (*p_BotA).ParaMap["SWC_volr"]  = 0.24
        (*p_BotA).ParaMap["SWC_mtsr"]  = 0.23
        (*p_BotA).ParaMap["SWC_tnr"]   = 0.11
        suc = true
        return suc
}

func initLogger()(error){
        logfile, err := os.OpenFile("run.log", os.O_APPEND|os.O_CREATE, 666)
        Log = log.New(logfile, "", log.Ldate | log.Ltime)
        return err
}


//----------------------------------------------------------------------
// read bottom or top date.
// Bottom date will be set manually in the "date/bot_date" file

func ReadFile(fn string)(rbytes []byte, err error){
	rbytes, err = ioutil.ReadFile(fn);

	if err != nil{
	       	fmt.Println("^^$$$$^^: fn, rbytes, err:", fn, rbytes, err)
		Log.Fatal("Error: <Readfile>: read file err, fn is:", fn)
	        return nil, err
	}
	return
}

func ReadBtDate(fn string)(o []string){
        //var botdate []byte
        fbyte, err := ReadFile(fn)
        if err != nil{
        	fmt.Println("Error: <ReadBtDate>, err is:", err)
        	return nil
        }

        for _, line := range bytes.Fields(fbyte){
                if len(line) > 0{                                      // avoid manually blank lines
                        o = append(o, strings.TrimSpace(string(line)) )
                }
        }
        return
}


// ----------------------- Read json -------------------------------------------------------
/* use interface to faciliate future modification
/* JSON 值可以是：数字,字符串(双引号),bool(true/false,) 数组（方括号）, 对象（花括号）, null
/*  Bool<-json Bool; float64<-Json number; string<-json string;
/* []interface{}<-Json array; map[string]interface <- json obj; nil<- Json nul
/*-----------------------------------------------------------------------------------------*/
func JsonReadUnmash(fn_json string)(rdmapIf map[string]interface{}, err error){
        rbytes, err := ReadFile(fn_json)
	var rdIf interface{}           // left of assertion must be interface{}
        if err = json.Unmarshal(rbytes, &rdIf); err != nil{
        	fmt.Println("Error: <JsonReadUnmash> unmarshall error:", err)
        	return
        }

        // type assertion to use interface
        runIf, ok := rdIf.(map[string]interface{})
       	fmt.Println("Info:runIf, ok:", runIf, ok)

        if !ok{
        	fmt.Println("Error: <JsonReadUnmash> type assertion err, ok value:", ok)
        }
        return runIf, nil
}


func ReadRunData(fn_rundata string)(rundata T_Rundata, err error){
        rdIf, err := JsonReadUnmash(fn_rundata)
      	rundata = JsonExtr_Rundata(rdIf)
	return
}


func JsonExtr_Rundata(runIf map[string]interface{})(rundata T_Rundata){
        for k, v := range runIf{
        	//type assetion to use interface
		switch k{     //switch vt := v.(type){
		case "LastTop":
            		vt_if, _ := v.(map[string]interface{})
			for key, value := range vt_if{
		   		if key == "date"{ rundata.LastTopDate = value.(string) }
		        }
		case "LastBot":
            		vb_if, _ := v.(map[string]interface{})
			for key, value := range vb_if{
				if key == "date"{ rundata.LastBotDate = value.(string)}
		        }
		default:
			fmt.Println("#Info: other json key that not processed.")
		}//switch
        }//for
        return
}


func ReadCalRes(fn_calRes string) (calRes T_CalRes, err error){
        rdIf, err := JsonReadUnmash(fn_calRes)
       	calRes = JsonExtr_CalRes(rdIf)
	return
}


func JsonExtr_CalRes(calResIf map[string]interface{})(calRes T_CalRes){
        for k, v := range calResIf{
                switch k{            // switch vt := v.(type){
                case "btm":          // json obj -> map[string]interface{}
	        	v_btm_if, _ := v.(map[string]interface{})
                        fmt.Println("Info: v_btm_if type:", reflect.TypeOf(v_btm_if).String())
                	for key, value := range v_btm_if{
                        	if key == "bi"{ calRes.Bi = int(value.(float64)) }
                        	if key == "ti"{ calRes.Ti = int(value.(float64)) }
                        }
                case "scan_res":     // json array -> []interface{}
	        	v_scan_if, _ := v.([]interface{})
      			fmt.Println("Info: scan_res type:",  reflect.TypeOf(v_scan_if).String())
                        for index, value := range v_scan_if{
                        	calRes.Scan_res[index]    = int(value.(float64))
                        }
		case "rrun_res":
			fmt.Println("Info: rrun_res type:", reflect.TypeOf(v).String())
			calRes.Rrun_res = v.(float64)
                default:
                	fmt.Println("#Info: other json key that not processed.")
                }//switch
        }//for
        return
}


// read real run result
func ReadSimTradeRes()(){

}

//----------------------------------- write json result ---------------------------------
// write calca result(json)
func WriteRes(res *T_CalRes) (err error) {
        //t := time.Now()

        if jsonRes, err := json.Marshal(res); err == nil{
        	fmt.Println("res转json成功")
        }else{
        	if ioutil.WriteFile("res.json", jsonRes, 0644) != nil {
        		return errors.New("写文件res.json出错.")
        	}
        }
        return
}


/**********************************************************************************
/* make bot data(pe_total/volr_total/mtsr_total) to feed Bottom Weighter
/*
/*********************************************************************************/
func ProcBotsData(a_punch, a_relax []T_A)(omap map[string]float64){
	omap = make(map[string]float64)

	// BotA.PeMap['pe_relax'] = peRes
	omap["pe_relax"], _, _   = pe_proc(a_relax)
	omap["pe_punch"], _, _   = pe_proc(a_punch)

	omap["volr_relax"], _, _ = volr_proc(a_relax)
	omap["volr_punch"], _, _ = volr_proc(a_punch)

	omap["mtsr_relax"], _, _ = mtsr_proc(a_relax)
	omap["mrsr_punch"], _, _ = mtsr_proc(a_punch)

	return
}


func pe_proc(a []T_A)(pe_total, pe_sh, pe_sz float64){
	pes_sh, pes_sz, pes_total := []float64{}, []float64{}, []float64{}
	for i:=0; i<len(a);i++{
		if a[i].Pe.Pe_sh != 0 &&  a[i].Pe.Pe_sz != 0{
			pes_total = append(pes_total, a[i].Pe.Pe_total)
			pes_sh    = append(pes_sh, a[i].Pe.Pe_sh)
			pes_sz    = append(pes_sz, a[i].Pe.Pe_sz)
		}
	}
	pe_sh    = SimpleAvg(pes_sh)
	pe_sz    = SimpleAvg(pes_sz)
	pe_total = SimpleAvg(pes_total)

	return
}


// another func to reserve for further modification
func volr_proc(a []T_A)(volr_total, volr_sh, volr_sz float64){
        volrs_sh, volrs_sz, volrs_total := []float64{}, []float64{}, []float64{}
        for i:=0; i<len(a);i++{
        	if a[i].Volr.Volr_sh != 0 && a[i].Volr.Volr_sz != 0{
                	volrs_total = append(volrs_total, a[i].Volr.Volr_total)
                	volrs_sh    = append(volrs_sh, a[i].Volr.Volr_sh)
                	volrs_sz    = append(volrs_sz, a[i].Volr.Volr_sz)
		}
	}
        volr_total  = SimpleAvg(volrs_total)
        volr_sh     = SimpleAvg(volrs_sh)
        volr_sz     = SimpleAvg(volrs_sz)
	return
}


// another func to reserve for further modification
func mtsr_proc(a []T_A)(mtsr_total, mtsr_sh, mtsr_sz float64){
        mtsrs_sh, mtsrs_sz, mtsrs_total := []float64{}, []float64{}, []float64{}
        for i:=0; i<len(a);i++{
        	if a[i].Mtsr.Mtsr_sh != 0 && a[i].Mtsr.Mtsr_sz != 0{
                	mtsrs_total = append(mtsrs_total, a[i].Mtsr.Mtsr_total)
                	mtsrs_sh    = append(mtsrs_sh, a[i].Mtsr.Mtsr_sh)
                	mtsrs_sz    = append(mtsrs_sz, a[i].Mtsr.Mtsr_sz)
		}
	}
        mtsr_total  = SimpleAvg(mtsrs_total)
        mtsr_sh     = SimpleAvg(mtsrs_sh)
        mtsr_sz     = SimpleAvg(mtsrs_sz)
	return
}


//
func pb_Proc(pb []float64)(o float64){
	o = SimpleAvg(pb)
	return
}


/**********************************************************************************
/*  qif get bottom data from. return these data
/* a_p: data struct to save A punch bottom characters.
/* a_r: data struct to save A relax bottom characters
/**********************************************************************************/
func GetBotsData(fn_bot string, a []T_A )(bool){
        botsdate := GetBotsDate(fn_bot)
	fmt.Println("##(2)##: botsdate:", botsdate)

	i_bot_valid, j_day_valid := 0, 0
	lwsize := 0                                        // last valid window size
        for i_bot, win := range botsdate{
                for j_day, day := range win{
                	fmt.Println("##(3)##, i, j, day:", i_bot, j_day, day)
			time.Sleep(QIF_ACCESS_INTVL * time.Millisecond)
			dicmkt := qif.GetMarket(day)
                        if len(dicmkt ) != 0{
                                qif.FilDicToA(dicmkt, &a[(i_bot_valid)*lwsize + j_day_valid])
                                j_day_valid++;
                                Print("##(4)## i_bot_valid/i_bot_valid++:",i_bot_valid, j_day_valid )
                        }
                        if j_day_valid == PRE_SMP_NUM{
                        	Print("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`")
                        	break
                        }
                }
                i_bot_valid++;
                lwsize, j_day_valid = j_day_valid, 0
        }
        return true
}


// return all bottoms window date according the bottom record file in ../data/
func GetBotsDate(fnbotdate string)(o [][]string){
        botsDate := ReadBtDate(fnbotdate)
	fmt.Println("###@#$ bostDate:", botsDate)
        for _, date := range botsDate{
                bw := GetBotWindow(date)
                if len(bw) !=0 {
                	o = append(o, bw)
                }
        }
        return
}



/**********************  Get 1 bottom's windows ************************
/*  presmpnum: pre-sample number. Not use data after bottom point.
/***********************************************************************/

// use qif to get valid trade days
func GetBotWindow(date string)(validTradeDays []string){
	return qif.GetTradeDays(date)
}


// doesn't use qif to get valid trade days
func GetBotWindow_raw(date string, prenum int)(bw []string){
        var lastdaytmp time.Time

        if dateTime, err := time.ParseInLocation(TIME_LAYOUT_SHORT, date, time.Local); err != nil{   // (layout, value string)
                Print("<GetBotWindow> error: time.Parse error. Maybe wrong date input.")
        }else{
                lastdaytmp = dateTime
        }

        for i:=0; i < prenum; i++{
                lastdaytmp = LastDay(lastdaytmp)
                dayStr := strings.SplitAfter(lastdaytmp.Format(TIME_LAYOUT_STR), " ")   // func (t Time) Format(layout string)(string)
                bw = append(bw, strings.TrimSpace(dayStr[0]))
        }
        return
}


func LastDay(day time.Time)(lastday time.Time){
        lastday = day.AddDate(0, 0, -1)
        return
}


func OperateTime()(bool){
        now := time.Now()
        fmt.Println("now is:", now)

        d, _ := time.ParseDuration("-24h")
        d1 := now.Add(d)
        fmt.Println("d is:->",d, ", d1 is:->",d1)

        year, month, day := now.Date()    //func (t Time)Date()(year int, month Month, day int)
        fmt.Println("year, month, day are:",year, month, day)

       return true
}


