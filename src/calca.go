/**********************************************************************/
//v 0.01 top exe.    execute once every morning. --- 2018.8.28

package main

import (
        "cmn"
        "mipos"
	"scans"
        "fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main(){
    fmt.Println("hello,,,calca start run once...")
 
	//var logger *log.Logger   
    logfile, err := os.OpenFile("run.log",os.O_APPEND|O_CREATE, 666)
    logger := log.New(logfile, "", log.Ldate | log.Ltime)
    
    res := &CalRes{}	
    res.bi, res.ti, res.mix_cw = bt_calc()	       // mix_cw: model mix of (Casino) - (weighing machine)

    if err = write_res(res){
		logger.Println(err)
	}
        
}

// logger.SetFlags   logger.Panicln   logger.Fatal()






//(*--------------------最终推荐---------------------*)
ca_sug = top * 0.6;




//---------------- write json result ------------------
// write to json result
func write_res(res *CalRes) (err strings) {
	t := time.Now()
	
	if jsonRes, err := json.Marshall(res); err == nil{
		fmt.Println("res转json成功")
	}
    if ioutil.WriteFile("res.json", jsonRes, 0644) != nil {
		return "写文件res.json出错."
	}	
	
}


// --------------- Run Sim once every day---------------
simtrade



