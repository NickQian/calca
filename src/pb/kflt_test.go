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
	_ "fmt"
	"cmn"
	)


var kl []float64



func init(){
	//kline := qif.GetKline("000001.SH", "20190101", "20210112")
	kline := qif.GetKline("000001.SH", "20000101", "20210120")
	kl = KlinePreprc(kline)
	//fmt.Printf("<init>:kl is: %v \n", kl)
}


func TestKlinePlot(t *testing.T){
	suc := KlinePlot("000001.SH", "20140601", "20210116")
	if !suc{ t.Logf("<TestKlinePlot>: %v \n", cmn.ErrPlotFail)   }
}

func TestKlineFa(t *testing.T){
	fftRes := KlineFa(kl)
	suc := PlotFa(fftRes, "FreqAmp.png")

	if !suc {
		t.Logf("Error: Someting error during <TestKlineFa>. suc is: %v \n", suc)
	}
}

//
func TestFltK(t *testing.T){
	kfilted := FltK(kl)         // k line filtered, complex number
	suc1 := PlotSlicef(kfilted, "time","Amp","testFltK_filtedKline.png")
	suc2 := PlotKl(kl, "testFltK_orgKline.png")

	if !(suc1 && suc2 ) {
                t.Logf("Error: Someting error during <TestFltK>. suc1&suc2 are: %v, %v ", suc1, suc2)
        }
}
