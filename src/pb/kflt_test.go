/*********************************************************************************
/*--------------
/* License : BSD
/*--------------
/* v0.1: init version --- 2020.10.15
/*********************************************************************************/

package pb

import (
	"testing"
	"qif"
	"fmt"
	"cmn"
	)


var kp []float64

//var indexCode, startDay, endDay string = "000001.SH","19990901","20210120" // A stock valuable data
//var indexCode, startDay, endDay string = "000001.SH","20140601","20210116" // from 2014.6. 1618 points
var indexCode, startDay, endDay string = "000001.SH","20190101","20210112" // current trend period. 494 points
//var indexCode, startDay, endDay string = "000001.SH","20200701","20200715"  // short. for test


func init(){
	fmt.Println("Info: <kflt_test-init>: starting...")

	kline := qif.GetKline(indexCode, startDay, endDay)     // short. for test

	kp = KlinePreprc(kline)
	fmt.Printf("Info: <kflt_test-init> finished! len(kp) is: %v \n", len(kp))
}


func TestKlinePlot(t *testing.T){
	suc := KlinePlot(indexCode, startDay, endDay)           // from 2014.6

	if !suc{ t.Logf("<TestKlinePlot>: %v \n", cmn.ErrPlotFail)   }
}

func TestKlineFa(t *testing.T){
	KlinePlot(indexCode, startDay, endDay)
	fftRes := KlineFa(kp)
	suc := PlotFa(fftRes, "FA_lnDcRmv.png")

	if !suc {
		t.Logf("Error: Someting error during <TestKlineFa>. suc is: %v \n", suc)
	}
}

//
func TestFltK(t *testing.T){
	kfilted := FltK(kp)         // k line filtered, complex number
	suc1 := PlotSlicef(kfilted, "time","Amp","FltK_filtedKline.png")
	suc2 := PlotKl(kp, "FltK_orgKline.png")

	if !(suc1 && suc2 ) {
                t.Logf("Error: Someting error during <TestFltK>. suc1&suc2 are: %v, %v ", suc1, suc2)
        }
}
