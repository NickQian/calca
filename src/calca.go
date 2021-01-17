/**********************************************************************/
//v 0.1 top exe.    execute once every morning. --- 2018.8.28

package main

import (
        "mipos"
	"scans"
        . "cmn"
        "fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)


func init(){

	//var logger *log.Logger
    	logfile, err := os.OpenFile("run.log",os.O_APPEND|O_CREATE, 666)
    	logger := log.New(logfile, "", log.Ldate | log.Ltime)

        // fetch today data once to update A status
        if _, e := GetCurPE(); e != nil{
        	Log.Fatalln("fatal Err: Qif not get PE data ")
                panic("Error: Qif not get PE data")
        }
}


func main(){
    	fmt.Println("hello,,,calca start run once...")

	if HasBtEvent(){
		UpdateCW_Model()
		UpdateRunData()
	}

	CalcaToday()

    	RunSimtrade()

    	fmt.Println("--> calca finish (once).")
}




// --------------- Run Sim once every day---------------
func RunSimTrade()(){
	
}

