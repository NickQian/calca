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


var kl []float64



func init(){
	kline := qif.GetKline("000001.SH", "20210101", "20210112")
	kl = KlinePreprc(kline)
	fmt.Printf("<init>:kl is: %v \n", kl)
}


func TestKlinePlot(t *testing.T){
	suc := PlotKl(kl, "test_Kline_org.png")
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
	suc := PlotSlicef(kfilted, "time","Amp","testFltK_filtedKline.png")
	if !suc {
                t.Logf("Error: Someting error during <TestFltK>. suc is: %v", suc)
        }
}
