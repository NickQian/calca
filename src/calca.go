/*******************************************************************************************
/* top. contains 3 funcs: CalcaToday/Scan/Simtrade
/* ----
/*  License: BSD
/* ----
/* v0.1 top exe.    execute once every morning. --- 2018.8.28
/*******************************************************************************************/

package main

import (
        "fmt"
	//"encoding/json"
	//"io"
	//"io/ioutil"
	//"scans"
        . "cmn"
	"pa"
	"pb"
	. "fup"
	//"trdsim"
	. "plotit"
	. "define"
)



//func Calca()(){
func main(){
   	fmt.Println("---- hello, init done!  calca start run once ... ------ ")

	// (1) how is today
	CalcaToday()

	// (2) scan value
	//Scanv()

	// (3) trade simulation
    	//Simtrade()

    	fmt.Println("----> calca finish (once). ------ ")
}


//------- calculate today postion & plot --------
func CalcaToday()(bi, ti int, mix_cw T_Mipos){
	fmt.Println("--------- <CalcaToday> --------- ")

	//(1) // fetch data from qif
	Fup(true)

        //(2) current position
        mix_cw = pa.Mipos("A")
        PlotMipos(mix_cw)

	//(3) the filter
	pb.FltK(use_m_file=true)


	return
}


//------- Scan ------------------

func Scanv()(){

        sug_1, sug_2, sug_3 := pa.Scan("data/vlist/vlist.json")
        err := WriteScanRes(sug_1, sug_2, sug_3, FN_SCAN_RES)
        CheckErr(err, "Scanv")

}

//------- Run Sim once every day ----------
func Simtrade()(acc int, code [3]string){

	return
}



/*func init(){

	//var logger *log.Logger
    	logfile, err := os.OpenFile("run.log",os.O_APPEND|O_CREATE, 666)
    	logger := log.New(logfile, "", log.Ldate | log.Ltime)

        // fetch today data once to update A status
        //if _, e := GetCurPE(); e != nil{
        //	Log.Fatalln("fatal Err: Qif not get PE data ")
        //	panic("Error: Qif not get PE data")
        //}

	return
}*/



