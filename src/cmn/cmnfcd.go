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
        "log"
        . "define"
        "strings"
        "bytes"
        )


var Log *log.Logger
var T_Now time.Time
var Today string
var A = new(T_A)               // Inst A parameters

// inst the Bottom parameters container with 0
var Botpara map[string]float64 = make(map[string]float64)


//type Ftp func(a ...interface{})(n int, err error)
//var  Print Ftp = fmt.Println
var Print func(a ...interface{}) (n int, err error)  = fmt.Println



func init(){
        if logfile, err := os.OpenFile("../run.log",os.O_APPEND|os.O_CREATE, 666); err == nil{
	        Log = log.New(logfile, "", log.Ldate | log.Ltime)

        }else{
                panic("<cmnfd-init> Open logfile error")
        }

        initBotpara(&Botpara)

        T_Now = time.Now()
        TodaySlice := strings.SplitAfter(T_Now.Format(TIME_LAYOUT_STR), " ")
        Today = strings.TrimSpace(TodaySlice[0])

        Print("<cmn-init>: Today:", Today )
        Print("<cmn-init> done!", TodaySlice)
        //Log.Println("<cmn-init> done. T_Now, Today:", T_Now, Today )
}

func initBotpara(p *map[string]float64) (bool) {
        (*p)["SWC_pe"]   = 0.22       // Subjective Weight Correction for entropy weight method
        (*p)["SWC_pb"]   = 0.22
        (*p)["SWC_vol"]  = 0.22
        (*p)["SWC_tnr"]  = 0.11
        (*p)["SWC_mtss"] = 0.23

        return true
}

/*----------------------------------------------------------------------
/* JSON 值可以是：数字（整数或浮点数）, 字符串（在双引号中）, 逻辑值（true 或 false）
/*   数组（在方括号中）, 对象（在花括号中）, null
/************************************************************************/
func ReadCalcaRes(fn string) (o CalRes){

    //var jsonRes string
    var result string

    if jsonRes, err := ioutil.ReadFile(fn); err == nil{
        result := strings.Replace(string(jsonRes), "\n", "", 1)
	fmt.Println("read res.json success:", result)
    }else{
        fmt.Println("读取res.json错误")
    }

    //var res CalRes
    if err := json.Unmarshal([]byte(result), &o); err != nil{
            fmt.Println("res.json转struct错误")
    }

    return
}


// read real run result
func ReadRrunRes()(){

}


/*** Bottom date will be set manually in the "date/bot_date" file ***/
func ReadBotDate(fn string)(o []string){
        var botdate []byte

        if fbyte, err := ioutil.ReadFile(FN_BOT_DATE); err == nil{
                botdate = fbyte
        }else{
                Log.Fatal("<Get_Bot_Date>: read date file error.")
        }

        for _, line := range bytes.Fields(botdate){
                if len(line) > 0{                                      // avoid manually blank lines
                        o = append(o, strings.TrimSpace(string(line)) )
                }
        }
        return
}



/**********************************************************************************
/*
/*********************************************************************************/
func ProcBotsData()(){

}




/**********************************************************************************
/*  qif get bottom data from. return these data
/**********************************************************************************/
func GetBotsData(a []T_A )(bool){
        botsdate := GetBotsDate()
	fmt.Println("##(1)##: botsdate:", botsdate)
	i_bot_valid, j_day_valid := 0, 0
        for i_bot, win := range botsdate{
        	fmt.Println("##(2)##, i_bot, win", i_bot, win)
                for j_day, day := range win{
                	fmt.Println("##(3)##, j_day, day:", j_day, day)
			time.Sleep(1000 * time.Millisecond)
			dicmkt := qif.GetMarket(day)
                        if len(dicmkt ) != 0{
                                qif.FilDicToA(dicmkt, &a[(i_bot_valid)*(j_day_valid)])
                                j_day_valid++;
                                Print("##(4)## i_bot_valid/i_bot_valid++:",i_bot_valid, j_day_valid )
                        }
                }
                i_bot_valid++;
                j_day_valid = 0;
        }
        return true
}


// return all bottoms window date according the bottom record file in ../data/
func GetBotsDate()(o [][]string){
        botsDate := ReadBotDate(FN_BOT_DATE)

        for _, date := range botsDate{
                bw := GetBotWindow(date, PRE_SMP_NUM)
                o = append(o, bw)
        }
        return
}



/********  return 1 bottom window *********
/*  presmpnum: pre-sample number. Don't use data after bottom point. */
func GetBotWindow(date string, prenum int)(bw []string){
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
        fmt.Println("--now is`````>", now)

        d, _ := time.ParseDuration("-24h")
        d1 := now.Add(d)
        fmt.Println("--d is:->",d, "--d1 is:->",d1)

        year, month, day := now.Date()    //func (t Time)Date()(year int, month Month, day int)
        fmt.Println("--->",year, month, day)

       return true
}



func WriteBotInd()(bool){
        return true
}

