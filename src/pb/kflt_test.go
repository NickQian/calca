/*********************************************************************************
/*--------------
/* License : BSD
/*--------------
/* v0.1: init version --- 2020.10.15
/*********************************************************************************/

package pb

import (
	"testing"
	"cmn"
	"qif"
	)


var kl []float64



func init()( bool ){
	klineRvs := qif.GetKline("2001-01-01", "2020-12-30")
	kl = KlinePreprc(klineRvs)
	return true
}


func TestKlineFa(t *testing.T){
	suc := PlotKl(kl)
	KlineFa(kl)

	if !suc {
		t.Logf("Error: Someting error during <TestKlineFa>. suc is: %v", suc)
	}
}

//
func TestFltK(t *testing.T){
	kfc := FltK(kl)         // k line filtered, complex number
	suc := PlotFa(kfc)
	if !suc {
                t.Logf("Error: Someting error during <TestFltK>. suc is: %v", suc)
        }
}
