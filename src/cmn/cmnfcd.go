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
        //"strconv"
        "encoding/json"
        "os"
        "log"
        "strings"
        )


var Log *log.Logger

// inst the A parameters
var A = new(T_A)

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

        Log.Println("<cmn-init> done. Date:", time.Now().Format("2006-01-02 15:04:05") )
        Print("<cmn-init> done.")
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
func Get_calca_res(fn string) (o CalRes){

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


func Get_simtrade_res()(){

}


/*** Bottom date will be set manually in the "date/bot_date" file ***/
func Get_Bot_Date(fn string)(o []string){
        var botdate string

        if fbyte, err := ioutil.ReadFile(FN_BOT_DATE); err == nil{
                botdate = string(fbyte)
        }else{
                Log.Fatal("<Get_Bot_Date>: read date file error.")
        }

        for _, line := range strings.SplitAfter(botdate, "\n"){
                if len(line) > 0{                                      // avoid manually blank lines
                        o = append(o, line)
                }
        }
        return
}



func Get_Bot_Rec()(d[][]float64){
/*
    var botrec string
    var botdata [][]float64

    if rd_contents, err := ioutil.ReadFile(FN_BOT_DATA); err == nil{
        botrec =  strings.Replace(string(rd_contents), "\n", "", 1)
        fmt.Println("read botdata success:", botdata)
    }

    //for _, line := range strings.Split(string(data), "\n") {
    for i, line := range botrec{
        if f, err := strconv.ParseInt(line, 10, 64); err == nil{
                fmt.Println("read botdata ->float64:", line)
                d.append(f)
        }else{
                panic(err)
        }
    }
*/
    return
}


/**********************************************************************************
/*  fetch bottom indicators data from qif. also return these data
/*  presmpnum: pre-sample number. Don't use data after bottom point.
/**********************************************************************************/
func FetchBotInd(date string, presmpnum int)(d [][]float64){

        return
}

func WriteBotInd()(bool){
        return
}
