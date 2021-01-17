/*********************************************************
/* ----
/*  License: BSD
/* ----
/* v0.01  init version --- 2019.7.12
/*********************************************************/

package cmn


import (
         "testing"
         "cmn"
         )

/*************************************************************************************************/
/* this init value is for Normalization test. the result should be has 0 and 1 and another value */

/*var Adata = [][]float64{ {100.0001, 100.00001,  99.9999  },
                          {100,      100.001,    100.0001 },
                                                  //  {7100.0, 7300.0, 8012},
                          }
*/
var Adata = [][]float64{ {12,     11.5,   11.8,     11 },
                          {1.1,    1.01,   1.1001,   1.10001 },
                          {49,     5102,   7103,     88104},
                         }


/*
var Adata = [][]float64{ {12.6,   11.8,   12.2, 14.2, 12.2,  11.7},
                          {0.9,    1.3,    1.2,  1.6,  1.15,  1.21},
                          {7100.0, 7300.0, 8012, 7123, 6234,  8124},
                        }
*/

/*
func init ()(){
        Adata = make([][]float64, 20){ {9.2,    10.0,   18.1, 13.4, 12.2,  11.8},
                                        {0.9,    1.3,    1.2,  1.6,  1.15,  1.21},
                                        {7100.0, 7300.0, 8012, 7123, 6234,  8124},
                                       }
}
*/

func TestWeightSwc(t *testing.T){

        r := []float64{BotA.ParaMap["SWC_pe"],
                       BotA.ParaMap["SWC_volr"],
                       BotA.ParaMap["SWC_tnr"],
                      }

        w_ent := WeightEnt(Adata)

        w_scr := WeightSwc(w_ent, r)

        if len(w_scr)==0 || len(w_ent)==0{
                t.Fatalf("Error: len(w_scr)==0 or len(w_ent)==0 ")
        }else{
                t.Logf("w_ent is: %f, w_scr is: %f ", w_ent, w_scr)
        }
}

func TestWeightEnt(t *testing.T){
        w := cmn.WeightEnt(Adata)
        if len(w)==0{
                t.Fatalf("Error: len(w)==0")
        }else{
                t.Logf("w is: %f", w)
        }
}



func TestCalRowEntropy(t *testing.T){
        e := cmn.CalRowEntropy(Adata)
        if cap(e)==0 {
                t.Errorf("Error: len(e)==0")
                t.Fatalf("Error: len(e)==0")
        }else{
                t.Logf("TestCalRowEntropy: e is: %f \n", e)
        }
}


func TestSmpWeitInRow(t *testing.T){
        p, _ := cmn.Norm_ent(Adata)
        if len(p) == 0{
                t.Errorf("Error: len(p)==0")
        }

        for _, row := range p{
                xweit := cmn.SmpWeitInRow(row)
                t.Logf("<TEstsmpWeitInRow>: xweit is: %f", xweit)
        }
}


func TestNorm_ent(t *testing.T){
        xp, xn := cmn.Norm_ent(Adata)
        t.Logf("<T<TestNorm_ent> xp: %f, \n xn: %f. ", xp, xn)
}


func TestGetMinMax(t *testing.T){
        min, max := cmn.GetMinMax(Adata)
        if min == nil{
                t.Error("Err:<TestGetMinMax>: min is nil ")
        }
        if max == nil{
                t.Error("Err:<TestGetMinMax>: max is nil")
        }
        cmn.Print("<TestGetMinMax>: min, max:", min ,max)
}



func TestSimpleAvg(t *testing.T){
        f := []float64{111.11, 222.32, 333.33}
        avg := cmn.SimpleAvg(f)
        cmn.Print("<TestSimpleAvg>: avg is:", avg)
}
