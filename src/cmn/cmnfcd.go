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


//-------------------------- global var ------------------------------------
//var T *testing.T = &testing.T{}
var Log *log.Logger
var T_Now time.Time
var timeLocation *time.Location

var TodayStr string

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


//------------------------------ init ------------------------------------
// init funcs
//
//------------------------------------------------------------------------
func init(){
        if logfile, err := os.OpenFile(RUN_DIR+"run.log",os.O_APPEND|os.O_CREATE, 666); err == nil{
	        Log = log.New(logfile, "", log.Ldate | log.Ltime)

        }else{
                panic("<cmnfd-init> Open logfile error")
        }


	timeLocation, _ = time.LoadLocation("Asia/Shanghai")

	initToday()
        initBotpara(BotA)                // point input
	initLogger()
        //Log.Println("<cmn-init> done. T_Now, Today:", T_Now, Today )
}


func initToday(){
        T_Now = time.Now()
        TodaySlice := strings.SplitAfter(T_Now.Format(TIME_LAYOUT_STR), " ")
        TodayStr = strings.TrimSpace(TodaySlice[0])


        Print("<cmn-init>: TodayStr:", TodayStr )
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


//---------------------------- READ --------------------------------------
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


// read calca(analyze data) res
func ReadCalRes(fn_calRes string) (calRes T_CalRes, err error){
        rdIf, err := JsonReadUnmash(fn_calRes)
       	calRes = JsonExtr_CalRes(rdIf)
	return
}


// read one type events and collcet the "Cha"(character data) -> raw data martrix(before normalization)
func ReadEgg(fn_egg string)( map[string](map[string]float64)  ){
	rdIf, err := JsonReadUnmash(fn_egg)
	if err != nil{
		fmt.Printf("Err: <ReadEgg> JsonReadUnmash result is: %v, fn: %v  \n", err, fn_egg)
	}
	eggMap := JsonExtr_toMap(rdIf)
	return eggMap
}



func ReadResTrd(fn_resTrd string)(resTrd T_ResTrd, err error){
        rdIf, err := JsonReadUnmash(fn_resTrd)
      	resTrd = JsonExtr_ResTrd(rdIf)
	return
}

// read real run result
func ReadLastOp(fn_LastOp string) (lastOpDat T_LastOp, err error){
        rdIf, err := JsonReadUnmash(fn_LastOp)
        lastOpDat = JsonExtr_LastOp(rdIf)
        return
}


//------------ basic read -------------------
func JsonReadUnmash(fn_json string)(rdmapIf map[string]interface{}, err error){
        rbytes, err := ReadFile(fn_json)
	var rdIf interface{}           // left of assertion must be interface{}
        if err = json.Unmarshal(rbytes, &rdIf); err != nil{
        	fmt.Printf("Error: <JsonReadUnmash> unmarshall error: %v, fn: %v  \n", err, fn_json)
        	return
        }

        // type assertion to use interface
        runIf, ok := rdIf.(map[string]interface{})
       	fmt.Printf("Info:<JsonReadUnmash> fn:%v readed.  \n",  fn_json)

        if !ok{
        	fmt.Println("Error: <JsonReadUnmash> type assertion err, ok value:", ok)
        }
        return runIf, nil
}

//--- interface extract ----
// eig(egg) -> evts -> cha(1 evt) -> item
// jason read -> map
func JsonExtr_toMap(rdEggIf map[string]interface{} )(egg map[string](map[string]float64) ){
	egg = make(map[string](map[string]float64) )

	for tag, evtIf := range rdEggIf{
		//fmt.Printf("###1@@: tag:%v, evt:%v,  \n", tag, evtIf)
		var itemMap = make(map[string]float64 )
		if chaIf, ok := evtIf.(map[string]interface{} ); ok{           // "pe_sz": interface, "pb_sh": interface
			for key, numIf := range chaIf{                        // iterate the map[string]interface
				if num, ok := numIf.(float64); ok{       // "pe_sz": 42.05800
					itemMap[key] = num
				}
			}
		}else{
			fmt.Println("Error: <JsonExtr_toMap>???? type assertion fail.", reflect.TypeOf(evtIf).String() )
		}
		egg[tag] = itemMap
		//fmt.Printf("###2@@:  egg: %v   \n",  egg)
	} // for
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
                        	calRes.ScanRes[index]    = int(value.(float64))
                        }
		case "rrun_res":
			fmt.Println("Info: rrun_res type:", reflect.TypeOf(v).String())
			calRes.TrdRes.Acc = v.(int)
                default:
                	fmt.Println("#Info: other json key that not processed.")
                }//switch
        }//for
        return
}


// resTrd uses the same data struct of simulation
func JsonExtr_ResTrd(resTrdIf map[string]interface{})(resTrd T_ResTrd){
	for k, v := range resTrdIf{				       // assert 1: the whole map
        	switch k{
		case "acc":
			resTrd.Acc = int(v.(int))
		case "update":
			resTrd.Update  = v.(string)
                case "curCode":
       	        	v_codes_if, _ := v.(map[string]interface{})        // assert 2:  []interface{}
			fmt.Println("Info: v_codes_if type:",  reflect.TypeOf(v_codes_if).String())
			resTrd.CurCodBnk = CodeBankExtra(v_codes_if)
		default:
                	fmt.Println("#Info: <JsonExtr_resTrd>:other json key that not processed.")
		}// switch
	}//for
	return
}


func CodeBankExtra(Ifcb map[string]interface{})(cbnk  T_CodeBank){
	for key_ccode, ccode := range Ifcb{              // iterate in "Current Code"(3 or maore)
        	v_code_if, _ := ccode.(map[string]interface{})    // assert 3:  []interface{}
                switch key_ccode{
                case "curCode1":
                	cbnk.Code1 = CodeInfoExtra(v_code_if)
                case "curCode2":
                        cbnk.Code2 = CodeInfoExtra(v_code_if)
                case "curCode3":
                        cbnk.Code3 = CodeInfoExtra(v_code_if)
                default:
                        panic ("Error: <JsonExtr_ResTrd> goes in default branch.")
        	} //case
     	} //for
	return
}


func CodeInfoExtra(IfCode map[string]interface{})(cinfo T_CodeInfo){
	for key, ci := range IfCode{     // interate in code
		switch key{
		case "code":
			cinfo.Code   = ci.(string)
		case "share":
			cinfo.Share  = ci.(int)
		case "inprc":
	 		cinfo.Inprc  = ci.(float64)
		case "amount":
	 		cinfo.Amount = ci.(int)
		default:
			panic ("Error: <CodeInfoExtra> goes in default branch.")
		}
	}
	return
}



// need ...
func JsonExtr_LastOp(runIf map[string]interface{})(lastOpDat T_LastOp){
        for k, v := range runIf{
        	//type assetion to use interface
		switch k{
                case "date":
                        lastOpDat.Date        = v.(string)
		case "acc":
			lastOpDat.Acc         = v.(int)
        	case "op_evt_type":
			lastOpDat.Op_evt_type = v.(string)
		case "buy":
			v_cbnk_if, _ := v.(map[string]interface{})
            		lastOpDat.Buy_codbnk  = CodeBankExtra(v_cbnk_if)
		case "sale":
			v_cbnk_if, _ := v.(map[string]interface{})
            		lastOpDat.Sale_codbnk = CodeBankExtra(v_cbnk_if)
		default:
			fmt.Println("#Info: other json key that not processed.")
		}//switch
        }//for
        return
}



//----------------------------------- WRITE --------------------------------------------
// write json file  (0777/0644/0640)
// Bool<-json Bool;           float64<-Json number;             string<-json string;
// []interface{}<-Json array; map[string]interface <- json obj; nil<- Json nul
//--------------------------------------------------------------------------------------
func WriteEvtdata(fn_rec_data, fn_avg_data string, i_evt int, A_evt []T_A, eggMap map[string](map[string]float64) )(  ){
//func WriteEvtdata(i_evt int, A_evt []T_A, eggMap map[string]interface{} )(  ){
	fmt.Printf("===> <WriteEvtdata>:  A_evt: %v   \n",  A_evt )
	fmt.Printf("===> <WriteEvtdata>:   eggMap: %v   \n",   eggMap)
	if i_evt > 0{
      		WriteEvtWins(A_evt, fn_rec_data, true)      // if not first event, append
		WriteEvtAvg(eggMap, fn_avg_data, true)
	}else{
		WriteEvtWins(A_evt, fn_rec_data, false)     // if first event, new write
		WriteEvtAvg(eggMap, fn_avg_data, false)
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


func WriteScanRes(sug1, sug2, sug3 string, fn string) (err error) {
	sugMap := map[int]string{1:sug1, 2:sug2, 3:sug3}
	jsonBytes := []byte{}
	if jsonBytes, err = json.MarshalIndent(sugMap, "", "\t"); err == nil{
                fmt.Println("Info:<WriteScanRes> map转json成功")
        }else{
                Print("ERROR: <WriteScanRes> map to Json 出错:", err)
        }

	err = Wr_Json(jsonBytes, fn)
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



//------------- Json <-> A struct ----------
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
	        Print("ERROR: map to Json 出错:", err)
	}
	return
}



//-------- basic write ---------
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




/******************************Eig -> Dm *******************************************************
/* prepare bot/top data(pe_total/volr_total/mtsr_total) to feed trainig(Bottom/top Weighter)
/*
/**********************************************************************************************/

// read eggs, normalize it, output the data matrix( for WeightEntropy)
func GetEigDm(fn_eggBotRlx, fn_eggBotPuc, fn_eggTopCrz, fn_eggTopHot string)(dmEigRlx, dmEigPuc, dmEigHot, dmEigCrz [][]float64, suc bool){
	suc = false
	// read files in "/data/proc"
        eggRlx := ReadEgg(fn_eggBotRlx)
        eggPuc := ReadEgg(fn_eggBotPuc)
        eggHot := ReadEgg(fn_eggTopHot)
        eggCrz := ReadEgg(fn_eggTopCrz)

	//// simple average it. relsults: map
	//eggAvgRlx, _ := SimpleAvg_Eggmap(eggRlx)
	//eggAvgPuc, _ := SimpleAvg_Eggmap(eggPuc)
	//eggAvgHot, _ := SimpleAvg_Eggmap(eggHot)
	//eggAvgCrz, _ := SimpleAvg_Eggmap(eggCrz)
	// Average the PE/PB/TNR ... may be not good idea

        // maps to matrix
        dmEigRlx = Eggs2Dm(eggRlx)
        dmEigPuc = Eggs2Dm(eggPuc)
        dmEigHot = Eggs2Dm(eggHot)
        dmEigCrz = Eggs2Dm(eggCrz)

	suc = true

	return dmEigRlx, dmEigPuc, dmEigHot, dmEigCrz, suc
}



/**********************************************************************************
/*  1)qif get bot/top windows data.
/*  2)average it and return map (also write the rec file and the avg data file)
/**********************************************************************************/
// get Bottom/Top data bases on date writen in BT fn. Then write result uses <WriteEvtdata> 
func GetBtsData(fn_bt string, fn_rec_data, fn_avg_data string)(eggs map[string](map[string]float64), A_evt []T_A, suc bool){
        A_evt = make([]T_A, WIN_SIZE)                 // eg. 5 T_A per event   panic: runtime error: index out of range
        //var eggmap map[string]interface{}
        var egg_map map[string](map[string]float64 )
        eggs = make( map[string](map[string]float64) )

        btsdate := GetBtsDate(fn_bt)

	i_evt_valid, j_day_valid := 0, 0
	lwsize := 0                                   // last valid window size

        for i_evt, win := range btsdate{
                for j_day, day := range win{    //--> In a event
                	fmt.Printf("Info:(stp 1): i_evt:%v, j_win:%v, day:%v   \n", i_evt, j_day, day)
		        time.Sleep(QIF_ACCESS_INTVL * time.Millisecond)
			dicmkt_raw := qif.GetMarket_raw(day)
			dicmkt := mktRaw2dict(day, dicmkt_raw)

                        if len(dicmkt ) != 0{
                                //A_evt[j_day_valid].Evt_Tag = day
                                qif.FilDicToA(dicmkt, &A_evt[j_day_valid], DateStrAddSlash(day) )
                                //fmt.Println("@@@ A is:", A_evt )
                                j_day_valid++;
                                //Print("Info(2) i_evt_valid/j_day_valid are:",i_evt_valid, j_day_valid)
                        }
                        if j_day_valid == PRE_SMP_NUM{
   			        egg_map = avgEvt(day, A_evt)        //eggmap map[string]interface{}
   			        eggAppend(&eggs, egg_map)
				WriteEvtdata(fn_rec_data, fn_avg_data, i_evt, A_evt, egg_map )
				//Print("------#(3)#---- write done-----")
                        	break
                        }
                } // end a event
                i_evt_valid++;
                lwsize, j_day_valid = j_day_valid, 0
                log.Printf("~~~~~~~~~~~~~~ a event finish, window size: %v ~~~~~~~~~~~~~~~~~~~", lwsize)
        } // all events

        return eggs, A_evt, true
}


// use <mktRaw2dict> to generate the *_total parameters
func GetCurMarket()(map[string]float64) {
	dicRaw := qif.GetCurMarket_raw()
	dicMkt := mktRaw2dict(TodayStr, dicRaw)
	return dicMkt
}


func GetMarket(day string)(map[string]float64){
	dicRaw := qif.GetMarket_raw(day)
        dicMkt := mktRaw2dict(day, dicRaw)
	return dicMkt

}

//func GetMarketBk( num_back int)(SlcDic[](map[string]float64)){
//	
//
//}


func mktRaw2dict(dateTag string, dicRaw map[string]float64)(dicMkt map[string]float64){
	var cmc_total, cmc_sz,  pe_sz,  pb_sz, tnr_sz float64
	cmc_sh          := dicRaw["cmc_sh"]
	cmc_szm, cmc_smb, cmc_gem  := dicRaw["cmc_szm"], dicRaw["cmc_smb"], dicRaw["cmc_gem"]
	dateSmbStart, _ := time.ParseInLocation(TIME_LAYOUT_SHORT, DATE_MKT_SMB_START, time.Local )
	dateGemStart, _ := time.ParseInLocation(TIME_LAYOUT_SHORT, DATE_MKT_GEM_START, time.Local )
	date520Evt,   _ := time.ParseInLocation(TIME_LAYOUT_SHORT, DATE_2015_0520EVT,  time.Local )
	dateTag_ := DateStrAddSlash(dateTag)
        dateSmp, _ := time.ParseInLocation(TIME_LAYOUT_SHORT, dateTag_, time.Local)
	dicMkt = dicRaw                                   // init dicMkt
	fmt.Printf("Info: init dicMkt with raw qif data: %v   \n", dicMkt)

	//--- get pe_sz,pb_sz ...etc
	if dateSmp.Before(dateSmbStart){                  //2004->2006.2.10.  use qif raw data
		cmc_sz, pe_sz, pb_sz, tnr_sz = dicRaw["cmc_szm"] , dicRaw["pe_szm"] , dicRaw["pb_szm"] , dicRaw["tnr_szm"] 
		fmt.Println("@: case 1: cmc_sz=cmc_szm, pe:", pe_sz)
	}else if dateSmp.Before(dateGemStart){            //2006.2.10->2010.6.18. smb started
		cmc_sz = cmc_szm + cmc_smb
		pe_sz  = dicRaw["pe_szm"] *(cmc_szm/cmc_sz) + dicRaw["pe_smb"] * (cmc_smb/cmc_sz)
		pb_sz  = dicRaw["pb_szm"] *(cmc_szm/cmc_sz) + dicRaw["pb_smb"] * (cmc_smb/cmc_sz)
		tnr_sz = dicRaw["tnr_szm"]*(cmc_szm/cmc_sz) + dicRaw["tnr_smb"]* (cmc_smb/cmc_sz)
		fmt.Printf("@: case 2: cmc_sz=cmc_szm(%v) + cmc_smb(%v), pe_szm: %v, pe_sz: %v  \n", cmc_szm, cmc_smb, dicRaw["pe_szm"], pe_sz)
	}else if dateSmp.Before(date520Evt){              //2010.6.18->2015.5.20. smb + gem started, but not included in szm.
                cmc_sz = cmc_szm + cmc_smb + cmc_gem
                pe_sz  = dicRaw["pe_szm"] *(cmc_szm/cmc_sz) + dicRaw["pe_smb"] * (cmc_smb/cmc_sz) + dicRaw["pe_gem"] * (cmc_gem/cmc_sz)
                pb_sz  = dicRaw["pb_szm"] *(cmc_szm/cmc_sz) + dicRaw["pb_smb"] * (cmc_smb/cmc_sz) + dicRaw["pb_gem"] * (cmc_gem/cmc_sz)
                tnr_sz = dicRaw["tnr_szm"]*(cmc_szm/cmc_sz) + dicRaw["tnr_smb"]* (cmc_smb/cmc_sz) + dicRaw["tnr_gem"]* (cmc_gem/cmc_sz)
                fmt.Printf("@: case 3: cmc_sz=cmc_szm+ cmc_smb + cmc_gem:%v, pe_szm:%v, pe_sz:%v  \n", cmc_sz, dicRaw["pe_szm"], pe_sz)
        }else{                                            // after 2015.5.20.
		cmc_sz  = dicRaw["cmc_szm"]  // Cmc_szm already include Cmc_gem, cmc_smb, or partially ?
		pe_sz, pb_sz, tnr_sz = dicRaw["pe_szm"] , dicRaw["pb_szm"] , dicRaw["tnr_szm"]  // but 399001.SZ can be used as the market representation
                fmt.Println("@: case 4. sames as 1: cmc_sz=cmc_szm: pe_sz:", cmc_sz, pe_sz)
	}

	//--- get pe_total, pb_total ...etc
	cmc_total = cmc_sh + cmc_sz
	wei_sh,  wei_sz  := cmc_sh/cmc_total,  cmc_sz/cmc_total
	pe_total  := dicRaw["pe_sh"]  * wei_sh + pe_sz  * wei_sz
	pb_total  := dicRaw["pb_sh"]  * wei_sh + pb_sz  * wei_sz
	tnr_total := dicRaw["tnr_sh"] * wei_sh + tnr_sz * wei_sz
	dicMkt["cmc_total"],  dicMkt["cmc_sz"]  = cmc_total,  cmc_sz           // these 2 need process
	dicMkt["pe_total"],  dicMkt["pe_sz"]  = pe_total,  pe_sz
	dicMkt["pb_total"],  dicMkt["pb_sz"]  = pb_total,  pb_sz
	dicMkt["tnr_total"], dicMkt["tnr_sz"] = tnr_total, tnr_sz
	fmt.Printf("Info: <mktRaw2dict> cmc_sh: %v,cmc_sz: %v,cmc_szm:%v, cmc_smb:%v, cmc_gem:%v  \n", cmc_sh, cmc_sz, cmc_szm, cmc_smb, cmc_gem )
	fmt.Printf("Info: <mktRaw2dict> pe_tatal: %v, wei_sh: %v, wei_sz:%v  \n", pe_total, wei_sh,  wei_sz)
	return
}

// convert egg map to to dm_eig
//func Eggs2Dm(eggs map[string]interface{})(dm_eig [][]float64){
func Eggs2Dm(eggs map[string](map[string]float64) )(dm_eig [][]float64){
	// row_pe_total etc...
        var r_pe_total,   r_pe_sh,   r_pe_sz   []float64
        var r_pb_total,   r_pb_sh,   r_pb_sz   []float64
        var r_tnr_total,  r_tnr_sh,  r_tnr_sz  []float64
        var r_volr_total, r_volr_sh, r_volr_sz []float64
        var r_mtsr_total, r_mtsr_sh, r_mtsr_sz []float64

	//------------- get ordered key slice ------------------------------
	var eggkeys []string                            // key is day string
        for ek, _ := range eggs{
              	eggkeys = append(eggkeys, ek)
	}
	sort.Strings(eggkeys)

        //------- iterate the key slice to append the date's eigen data ------
        for _, key_day := range eggkeys{
        	egg_sorted := eggs[key_day]
        	//if egg_m, ok := egg_sorted.(map[string]float64); ok{           // do interface type assertion
        	//        for k, eig:= range egg_m{
       	        for k, eig := range egg_sorted{
       	   		if k == "pe_total"  { r_pe_total   = append(r_pe_total,   eig) }
       	   		if k == "pe_sh"     { r_pe_sh      = append(r_pe_sh,      eig) }
       	   		if k == "pe_sz"     { r_pe_sz      = append(r_pe_sz,      eig) }

                        if k == "pb_total"  { r_pb_total   = append(r_pb_total,   eig) }
                        if k == "pb_sh"     { r_pb_sh      = append(r_pb_sh,      eig) }
                        if k == "pb_sz"     { r_pb_sz      = append(r_pb_sz,      eig) }

                        if k == "tnr_total" { r_tnr_total  = append(r_tnr_total,   eig) }
                        if k == "tnr_sh"    { r_tnr_sh     = append(r_tnr_sh,      eig) }
                        if k == "tnr_sz"    { r_tnr_sz     = append(r_tnr_sz,      eig) }

                        if k == "volr_total"{ r_volr_total = append(r_volr_total, eig) }
                        if k == "volr_sh"   { r_volr_sh    = append(r_volr_sh,    eig) }
                        if k == "volr_sz"   { r_volr_sz    = append(r_volr_sz,    eig) }

                        if k == "mtsr_total"{ r_mtsr_total = append(r_mtsr_total, eig) }
			if k == "volr_sh"   { r_mtsr_sh    = append(r_mtsr_sh,    eig) }
			if k == "mtsr_sz"   { r_mtsr_sz    = append(r_mtsr_sz,    eig) }
                } // end the sorted single egg element extract
		//}
        }  // end dated eggs map extraction

        // ----------- assemble the data matrix ----------------------------
	// note the matrix row/col sequence
	dm_eig = append(append(append(append(append(append(append(append(append(append(append(append(append(append(append(dm_eig,
        		r_pe_total  ), r_pe_sh  ), r_pe_sz  ),
                        r_pb_total  ), r_pb_sh  ), r_pb_sz  ),
                        r_tnr_total ), r_tnr_sh ), r_tnr_sz ),
                        r_volr_total), r_volr_sh), r_volr_sz),
                        r_mtsr_total), r_mtsr_sh), r_mtsr_sz)
	fmt.Println("Info:<Eggs2Dm> dm_eig:", dm_eig)

        return
}


func DmAppend(dm *[][]float64, dmIn [][]float64) bool {
	//fmt.Printf("<DmApend>@@@: dmIn: %v, *dm:%v  \n", dmIn, *dm)
	for i, row := range dmIn{
		(*dm)[i] = append((*dm)[i], row...)
		//fmt.Printf(" row: %v, dmIn[i]:%v, *dm:%v  \n", row, dmIn[i], *dm)
	}
	return true
}



//func eggAppend(egg *map[string]interface{}), egg_in map[string]interface{})( bool ){
func eggAppend(egg *map[string](map[string]float64), egg_in map[string](map[string]float64) )( bool ){
	for key, item := range egg_in{
		  (*egg)[key] = item
	}
	return true
}


// remove the first empty element and rows behind that in DM(the empty elements are for the future)
func DmClean(dm *[][]float64)(int){
	for i, row := range *dm{
		if len(row) == 0{
			*dm = (*dm)[:i]
			return i
		}
	}
	return len(*dm)
}

//------------------------------------------------------------------------------------------
// average maps in a Event (windows data), output the "eggmap" which is 1 event eigen map
// eggmap: map[ ["day"]: "2019-01-05"
//              ["item"]: [map[pe] : 14.2 ]
func avgEvt(dayStr string, a []T_A)(eggmap map[string](map[string]float64) ){ // eggmap map[string](map[string]float64) ){
	pes_sh,   pes_sz,   pes_total   := []float64{}, []float64{}, []float64{}
	pbs_sh,   pbs_sz,   pbs_total   := []float64{}, []float64{}, []float64{}
	tnrs_sh,  tnrs_sz,  tnrs_total  := []float64{}, []float64{}, []float64{}
	//mtsrs_sh, mtsrs_sz, mtsrs_total := []float64{}, []float64{}, []float64{}
	//volrs_sh, volrs_sz, volrs_total := []float64{}, []float64{}, []float64{}

        //eggmap = make(map[string]interface{} )
	eggmap = make(map[string](map[string]float64) )

	for i:=0; i<len(a);i++{
		if a[i].Pe.Pe_sh     != 0 &&  a[i].Pe.Pe_sz != 0{
			pes_total    = append(pes_total,   a[i].Pe.Pe_total)
			pes_sh       = append(pes_sh,      a[i].Pe.Pe_sh)
			pes_sz       = append(pes_sz,      a[i].Pe.Pe_sz)
		}
		if a[i].Pb.Pb_sh     != 0 &&  a[i].Pb.Pb_sz != 0{
                        pbs_total    = append(pbs_total,   a[i].Pb.Pb_total)
                        pbs_sh       = append(pbs_sh,      a[i].Pb.Pb_sh)
                        pbs_sz       = append(pbs_sz,      a[i].Pb.Pb_sz)
                }
                if a[i].Tnr.Tnr_sh   != 0 &&  a[i].Tnr.Tnr_sz != 0{
                        tnrs_total   = append(tnrs_total,  a[i].Tnr.Tnr_total)
                        tnrs_sh      = append(tnrs_sh,     a[i].Tnr.Tnr_sh)
                        tnrs_sz      = append(tnrs_sz,     a[i].Tnr.Tnr_sz)
                }
        	/*
        	if a[i].Mtsr.Mtsr_sh != 0 && a[i].Mtsr.Mtsr_sz != 0{
                	mtsrs_total  = append(mtsrs_total, a[i].Mtsr.Mtsr_total)
                	mtsrs_sh     = append(mtsrs_sh,    a[i].Mtsr.Mtsr_sh)
                	mtsrs_sz     = append(mtsrs_sz,    a[i].Mtsr.Mtsr_sz)

		if a[i].Volr.Volr_sh != 0 && a[i].Volr.Volr_sz != 0{
                	volrs_total  = append(volrs_total, a[i].Volr.Volr_total)
                	volrs_sh     = append(volrs_sh,    a[i].Volr.Volr_sh)
                	volrs_sz     = append(volrs_sz,    a[i].Volr.Volr_sz)
		}
		}*/
	}
        // ---------- prepare the map for wr & dm ----------------
	//eggmap["day"] = dayStr
        var eggc = make(map[string] float64)

        pe_sh       := SimpleAvg(pes_sh)
	pe_sz       := SimpleAvg(pes_sz)
	pe_total    := SimpleAvg(pes_total)
        eggc["pe_total"], eggc["pe_sh"], eggc["pe_sz"] = pe_total,  pe_sh,   pe_sz

        pb_sh       := SimpleAvg(pbs_sh)
        pb_sz       := SimpleAvg(pbs_sz)
        pb_total    := SimpleAvg(pbs_total)
        eggc["pb_total"], eggc["pb_sh"], eggc["pb_sz"] = pb_total,  pb_sh,   pb_sz

        tnr_sh      := SimpleAvg(tnrs_sh)
        tnr_sz      := SimpleAvg(tnrs_sz)
        tnr_total   := SimpleAvg(tnrs_total)
        eggc["tnr_total"], eggc["tnr_sh"], eggc["tnr_sz"] = tnr_total, tnr_sh, tnr_sz

        /*volr_total  := SimpleAvg(volrs_total)
        volr_sh     := SimpleAvg(volrs_sh)
        volr_sz     := SimpleAvg(volrs_sz)
        eggc["volr_total"], eggc["volr_sh"], eggc["volr_sz"] = volr_total, volr_sh, volr_sz

        mtsr_total  := SimpleAvg(mtsrs_total)
        mtsr_sh     := SimpleAvg(mtsrs_sh)
        mtsr_sz     := SimpleAvg(mtsrs_sz)
        eggc["mtsr_total"], eggc["mtsr_total"], eggc["mtsr_sz"] = mtsr_total, mtsr_sh, mtsr_sz
	*/

	dayStr_ := DateStrAddSlash(dayStr)
	eggmap[dayStr_] = eggc

	return
}


// return all bottoms window date according the bottom record file in ../data/
func GetBtsDate(fnbtdate string)(o [][]string){
        btsDate := ReadBtDate(fnbtdate)
	fmt.Println("Info: botsDate are:", btsDate)
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
	return qif.GetTradeDays(date, PRE_SMP_NUM)
}


// doesn't use qif to get valid trade days. don't use this.
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


// eg. input: 20150520, output: 2010-05-20
func DateStrAddSlash (dateStrShort string)(dateStr string){
	dateByte := []byte(dateStrShort)
	byteTmp  := [][]byte{ dateByte[0:4], []byte("-"), dateByte[4:6], []byte("-"), dateByte[6:8] }
	for _, v := range byteTmp{
        	strTmp := string(v)
		//fmt.Printf("##, v: %v, strTmp: %v   \n ", v, strTmp)
		dateStr = dateStr + strTmp
	}
	//fmt.Print("## dateStr:", dateStr,)
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

func Delay_Qif_Intvl()(){

        time.Sleep(QIF_ACCESS_INTVL * time.Millisecond)
	return
}
