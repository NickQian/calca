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
	"sort"
        )


//var T *testing.T = &testing.T{}
var Log *log.Logger
var T_Now time.Time
var Today string
var A = new(T_A)               // Inst A parameters

var (
        ErrPasswd       = errors.New("ERR: password or user name not correct \n")
        ErrNoDataReturn = errors.New("ERR: no data return from qif \n")
        ErrEmptyNoItem  = errors.New("ERR: slice or map is empty.no item to process \n")
	ErrPlotFail     = errors.New("ERR: some error happens during plot.") 
)

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
//------------------------------------------------------------------------

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


/* ----------------------- Read json -------------------------------------------------------
/* use interface to facilitate future modification
/* JSON 值可以是：数字,字符串(双引号),bool(true/false,) 数组（方括号）, 对象（花括号）, null
/*  Bool<-json Bool; float64<-Json number; string<-json string;
/* []interface{}<-Json array; map[string]interface <- json obj; nil<- Json nul
/*-----------------------------------------------------------------------------------------*/
func ReadRunData(fn_rundata string)(rundata T_Rundata, err error){
        rdIf, err := JsonReadUnmash(fn_rundata)
      	rundata = JsonExtr_Rundata(rdIf)
	return
}


func ReadCalRes(fn_calRes string) (calRes T_CalRes, err error){
        rdIf, err := JsonReadUnmash(fn_calRes)
       	calRes = JsonExtr_CalRes(rdIf)
	return
}


// read real run result
func ReadRrunRes(fn_RrunRes string) (rrunRes T_SimRes, err error){
        rdIf, err := JsonReadUnmash(fn_RrunRes)
        rrunRes = JsonExtr_RrunRes(rdIf)
        return
}


// read file, output "map"
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


// RrunRes uses the same data struct of simulation
func JsonExtr_RrunRes(rrunResIf map[string]interface{})(rrunRes T_SimRes){ // assert 1: map[string]interface{}
	for k, v := range rrunResIf{
        	switch k{
		case "curValue":
			rrunRes.CurValue = int(v.(float64))
		case "curState":
			rrunRes.CurState = v.(string)
		case "date":
			rrunRes.CurDate  = v.(string)
                case "curCode":
       	        	v_codes_if, _ := v.([]interface{})                 // assert 2:  []interface{}
			fmt.Println("Info: v_codes_if type:",  reflect.TypeOf(v_codes_if).String())
                       	for i_arr, arr := range v_codes_if{
                               	v_code_if, _ := arr.([]interface{})        // assert 3:  []interface{}
                                for i_cinfo, cinfo := range v_code_if{     // see dfn.T_CodeInfo
                                       if i_cinfo == 0{ rrunRes.CurCode[i_arr].Code      = cinfo.(string) }
                                       if i_cinfo == 1{ rrunRes.CurCode[i_arr].OrgValue  = int(cinfo.(float64)) }
                                       if i_cinfo == 2{ rrunRes.CurCode[i_arr].CurValue  = int(cinfo.(float64)) }
                                }
                        }//for
		default:
                	fmt.Println("#Info: <JsonExtr_RrunRes>:other json key that not processed.")
		}// switch
	}//for
	return
}


/************************************ write json result *********************************
/* write json file  (0777/0644/0640)
/* Bool<-json Bool;           float64<-Json number;             string<-json string;
/* []interface{}<-Json array; map[string]interface <- json obj; nil<- Json nul
/****************************************************************************************/
//func WriteEvtdata(i_evt int, A_evt []T_A, eggMap map[string]interface{} )(  ){
func WriteEvtdata(i_evt int, A_evt []T_A, eggMap map[string](map[string]float64) )(  ){
      	if i_evt > 0{
      		WriteEvtWins(A_evt, FN_BOT_PUC_REC_DAT, true)      // if not first event, append
		WriteEvtAvg(eggMap, FN_BOT_PUC_AVG_DAT, true)
	}else{
		WriteEvtWins(A_evt, FN_BOT_PUC_REC_DAT, false)     // if first event, new write
		WriteEvtAvg(eggMap, FN_BOT_PUC_AVG_DAT, false)
	}
}

// write T_A to file directly
func WriteEvtWins(evtWin []T_A, fn string, add bool)(err error){
        jsonBytes, err := AtoJson(evtWin)
        if add{
                err = Wr_Json_(jsonBytes, fn)
        }else{
        	err = Wr_Json(jsonBytes, fn)
        }
        return
}


//func WriteEvtAvg(evtAvg map[string]interface{}, fn string, add bool)(err error){
func WriteEvtAvg(evtAvg map[string](map[string]float64), fn string, add bool)(err error){
	jsonBytes, err := mapToJson(evtAvg)
	if add{
		err = Wr_Json_(jsonBytes, fn)
	}else{
		err = Wr_Json(jsonBytes, fn)
	}
	return
}


// As is "A" struct slice
func AtoJson(As []T_A)(jsonBytes []byte, err error){
        //jsbyte := []byte
        for _, v := range As{
        	if jsbyte, err := json.MarshalIndent(v, "", "\t"); err == nil{        //json.MarshalIndent(struct, "", "    ")
        		fmt.Println("struct转json成功" )
        		for _, e := range jsbyte{
            			jsonBytes = append(jsonBytes,e)
            		}
        	}else{
                	Print("ERROR: Struct to Json 出错.")
        	}
        }
        return
}


//func mapToJson(m map[string]interface{} )(jsonBytes []byte, err error){
func mapToJson(m map[string](map[string]float64) )(jsonBytes []byte, err error){
	if jsonBytes, err = json.MarshalIndent(m, "", "\t"); err == nil{        //func MarshalIndent(v interface{}, prefix, indent string)([]byte, error)
	        fmt.Println("map转json成功")
	}else{
	        Print("ERROR: map to Json 出错.")
	}
	return
}


// ioutil new write
func Wr_Json(jsonByte []byte, fn string) (err error){
        //        WriteFile(filename string, data[]byte, perm os.FileMode) error
     	if ioutil.WriteFile(fn, jsonByte, os.ModeAppend|0644) != nil {         // 0644
        	Print("ERROR: 写文件出错.")
       		return errors.New("写文件出错.")
        }else{
                Print("Success.写文件成功:", fn)
        }
        return
}


// append write
func Wr_Json_(jsonByte []byte, fn string)(err error){
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644) //os.O_CREATE
	if err != nil{
    		return errors.New("ERROR: Open 文件出错")
    		log.Fatalf("Open 文件出错: %s", err)
    	}else{
    		Print("Open 文件成功, f,fn:", f, fn)
    	}

        defer f.Close()

        if _, err := f.Write(jsonByte); err != nil{
        //if _, err := f.WriteString(string(jsonByte)); err != nil{
        	log.Fatalf("写文件出错:%s", err)
		return errors.New("写文件出错.")
        }
        return
}




/*********************************************************************************************
/* prepare bot/top data(pe_total/volr_total/mtsr_total) to feed trainig(Bottom/top Weighter)
/*
/**********************************************************************************************/

// output the data matrix for WeightEntropy
func GetEigDm(fn string)(dm_eig [][]float64, suc bool){

        eggs, _, suc := GetBtsData(FN_BOT_PUC_DATE)
        dm_eig = eggs2eig(eggs)

	return dm_eig, suc
}

//func keystrip()(){
//}

/**********************************************************************************
/*  1)qif get bot/top windows data.
/*  2)average it and return map
/**********************************************************************************/
//func GetBtsData(fn_bt string)(eggs map[string]interface{}, A_evt []T_A, suc bool){
func GetBtsData(fn_bt string)(eggs map[string](map[string]float64), A_evt []T_A, suc bool){
        A_evt = make([]T_A, WIN_SIZE)                 // eg. 5 T_A per event   panic: runtime error: index out of range
        //var eggmap map[string]interface{}
        var egg_map map[string](map[string]float64 )
        //eggs = make(map[string]interface{})
        eggs = make( map[string](map[string]float64) )

        btsdate := GetBtsDate(fn_bt)
	fmt.Println("#(2): btsdate:", btsdate)

	i_evt_valid, j_day_valid := 0, 0
	lwsize := 0                                   // last valid window size

        for i_evt, win := range btsdate{
                for j_day, day := range win{    //--> In a event
                	fmt.Println("#(3), i_evt, j_win, day:", i_evt, j_day, day)
		        time.Sleep(QIF_ACCESS_INTVL * time.Millisecond)
			dicmkt := qif.GetMarket(day)

                        if len(dicmkt ) != 0{
                              //qif.FilDicToA(dicmkt, &a[(i_evt_valid)*lwsize + j_day_valid])
                                qif.FilDicToA(dicmkt, &A_evt[j_day_valid])
                                j_day_valid++;
                                Print("#(4) i_evt_valid/j_day_valid are:",i_evt_valid, j_day_valid)
                        }
                        if j_day_valid == PRE_SMP_NUM{
   			        egg_map = avgEvt(day, A_evt)        //eggmap map[string]interface{}
   			        eggAppend(&eggs, egg_map)
				WriteEvtdata(i_evt, A_evt, egg_map )
				Print("#(5) write done.")
                        	break
                        }
                } // end a event
                i_evt_valid++;
                lwsize, j_day_valid = j_day_valid, 0
                log.Printf("~~~~~~~~~~~~~~~~~~~~ a event finish, window size: %v ~~~~~~~~~~~~~~~~~~~~~~", lwsize)
        } // all events

        return eggs, A_evt, true
}



// average evt data to single win, then convert it to eig maps
//func eggs2eig(eggs map[string]interface{})(dm_eig [][]float64){
func eggs2eig(eggs map[string](map[string]float64) )(dm_eig [][]float64){
	//eggs = make(map[string][]float64 )

        var r_pe_total,   r_pe_sh,   r_pe_sz   []float64
        var r_volr_total, r_volr_sh, r_volr_sz []float64
        var r_mtsr_total, r_mtsr_sh, r_mtsr_sz []float64


	//------------- get ordered key slice ------------------------------
	var eggkeys []string                         // key is day string
        for ek, _ := range eggs{
        	//if _, ok := egg.(string); ok{        // do interface type assertion
              	eggkeys = append(eggkeys, ek)
              	//}
	}
	fmt.Println("<egg2eig>: eggkeys:", eggkeys)
	sort.Strings(eggkeys)
	fmt.Println("<egg2eig>: sorted keys: ", eggkeys)

        //------- iterate the key slice to append the date's eigen data ------
        for _, key_day := range eggkeys{
        	fmt.Println("@@key: , r_pe_total: ", key_day, r_pe_total)
        	egg_sorted := eggs[key_day]
        	//if egg_m, ok := egg_sorted.(map[string]float64); ok{           // do interface type assertion
        	//        for k, eig:= range egg_m{
       	        for k, eig := range egg_sorted{
       	   		if k == "pe_total"  { r_pe_total   = append(r_pe_total,   eig) }  // eig...
       	   		if k == "pe_sh"     { r_pe_sh      = append(r_pe_sh,      eig) }  // eig...
       	   		if k == "pe_sz"     { r_pe_sz      = append(r_pe_sz,      eig) }  // eig...

                        if k == "volr_total"{ r_volr_total = append(r_volr_total, eig) }
                        if k == "volr_sh"   { r_volr_sh    = append(r_volr_sh,    eig) }
                        if k == "volr_sz"   { r_volr_sz    = append(r_volr_sz,    eig) }

                        if k == "mtsr_total"{ r_mtsr_total = append(r_mtsr_total, eig) }
			if k == "volr_sh"   { r_mtsr_sh    = append(r_mtsr_sh,    eig) }
			if k == "mtsr_sz"   { r_mtsr_sz    = append(r_mtsr_sz,    eig) }
                } // end the sorted single egg element extract
  		fmt.Println("$ r_pe_total: , r_pe_sh: , r_pe_sz:", r_pe_total, r_pe_sh, r_pe_sz)
		//}
        }  // end dated eggs map extraction

        // ----------- assemble the data matrix ----------------------------
        dm_eig = append(append(append(append(append(append(append(append(append(dm_eig,  
                        r_pe_total  ), r_pe_sh  ), r_pe_sz  ),  
                        r_volr_total), r_volr_sh), r_volr_sz), 
                        r_mtsr_total), r_mtsr_sh), r_mtsr_sz)
        return
}


//
//func eggAppend(egg *map[string]interface{}), egg_in map[string]interface{})( bool ){
func eggAppend(egg *map[string](map[string]float64), egg_in map[string](map[string]float64) )( bool ){
	//Print("== <eggAppend> egg in:", egg_in)
	for key, item := range egg_in{
		  (*egg)[key] = item
	}
	//Print("=== <eggAppend> egg out:", *egg)
	return true
}


// average maps in a Event (windows data), output the "eggmap" which is 1 event eigen map
// eggmap: map[ ["day"]: "2019-01-05"
//              ["item"]: [map[pe] : 14.2 ]
func avgEvt(dayStr string, a []T_A)(eggmap map[string](map[string]float64) ){ // eggmap map[string](map[string]float64) ){
//func avgEvt(dayStr string, a []T_A)(eggmap map[string]interface{}){
	pes_sh,   pes_sz,   pes_total   := []float64{}, []float64{}, []float64{}
	volrs_sh, volrs_sz, volrs_total := []float64{}, []float64{}, []float64{}
	mtsrs_sh, mtsrs_sz, mtsrs_total := []float64{}, []float64{}, []float64{}
        //eggmap = make(map[string]interface{} )
	eggmap = make(map[string](map[string]float64) )

	for i:=0; i<len(a);i++{
		if a[i].Pe.Pe_sh     != 0 &&  a[i].Pe.Pe_sz != 0{
			pes_total    = append(pes_total,   a[i].Pe.Pe_total)
			pes_sh       = append(pes_sh,      a[i].Pe.Pe_sh)
			pes_sz       = append(pes_sz,      a[i].Pe.Pe_sz)
		}
        	if a[i].Volr.Volr_sh != 0 && a[i].Volr.Volr_sz != 0{
                	volrs_total  = append(volrs_total, a[i].Volr.Volr_total)
                	volrs_sh     = append(volrs_sh,    a[i].Volr.Volr_sh)
                	volrs_sz     = append(volrs_sz,    a[i].Volr.Volr_sz)
		}
        	if a[i].Mtsr.Mtsr_sh != 0 && a[i].Mtsr.Mtsr_sz != 0{
                	mtsrs_total  = append(mtsrs_total, a[i].Mtsr.Mtsr_total)
                	mtsrs_sh     = append(mtsrs_sh,    a[i].Mtsr.Mtsr_sh)
                	mtsrs_sz     = append(mtsrs_sz,    a[i].Mtsr.Mtsr_sz)
		}
	}
        // ---------- prepare the map for wr & dm ----------------
	//eggmap["day"] = dayStr
        var eggc = make(map[string] float64)

        pe_sh       := SimpleAvg(pes_sh)
	pe_sz       := SimpleAvg(pes_sz)
	pe_total    := SimpleAvg(pes_total)
        //eggmap["pe"]   = []float64{pe_total,   pe_sh,   pe_sz}
        eggc["pe_total"], eggc["pe_sh"], eggc["pe_sz"] = pe_total,  pe_sh,   pe_sz

        volr_total  := SimpleAvg(volrs_total)
        volr_sh     := SimpleAvg(volrs_sh)
        volr_sz     := SimpleAvg(volrs_sz)
        eggc["volr_total"], eggc["volr_sh"], eggc["volr_sz"] = volr_total, volr_sh, volr_sz

        mtsr_total  := SimpleAvg(mtsrs_total)
        mtsr_sh     := SimpleAvg(mtsrs_sh)
        mtsr_sz     := SimpleAvg(mtsrs_sz)
        eggc["mtsr_total"], eggc["mtsr_total"], eggc["mtsr_sz"] = mtsr_total, mtsr_sh, mtsr_sz

	eggmap[dayStr] = eggc

	return
}


// return all bottoms window date according the bottom record file in ../data/
func GetBtsDate(fnbtdate string)(o [][]string){
        btsDate := ReadBtDate(fnbtdate)
	fmt.Println("## botsDate:", btsDate)
        for _, date := range btsDate{
                bw := GetBtWindow(date)
                if len(bw) !=0 {
                	o = append(o, bw)
                }
        }
        return
}



/**********************  Get 1 bottom's windows ************************
/*  presmpnum: pre-sample number. Not use data after bottom point.
/***********************************************************************/

// use qif to get a bot/top window valid trade days
func GetBtWindow(date string)(validTradeDays []string){
	return qif.GetTradeDays(date)
}


// doesn't use qif to get valid trade days
func GetBtWindow_raw(date string, prenum int)(bw []string){
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


// reverse the input slice(float64)
func ReverseSlc(s []float64)([]float64){
	for i,j := 0,len(s)-1; i < j;  i,j = i+1,j-1{
		s[i], s[j] = s[j], s[i]
	}
	return s
}
