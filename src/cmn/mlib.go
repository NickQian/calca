/*******************************************************************
/* math lib for this project
/* ----
/*  License: BSD
/* ----
/* 0.1   math libs for policy B  --- 2019.1.26
/******************************************************************/

package cmn

import(
        "math"
	//"testing"
        "fmt"
)

//var T *testing.T = &testing.T{}

// simple average a slice
func SimpleAvg(din []float64)(avg_out float64){
	var avg float64 = 0.0

	for _, val := range din{
		avg += val
        }
	avg = float64(avg)/float64(len(din))

	return avg
}


// input eg.  ["2010-05-20"](["pe_sh"] 2.548, ["pb_sh"] 1.2 )
func SimpleAvg_Eggmap(mapIn map[string](map[string]float64) )(avg_mout map[string]float64 ){
	var avgPe, avgPb, avgTnr float64 = 0.0, 0.0, 0.0

	for tag, evtMap := range mapIn{
		fmt.Printf("@ SimpleAvg_map: tag:%v, evt[pe_total]:%v,[pb_total]:%v,[tnr_total]: %v  \n", tag, evtMap)
		avgPe  += evtMap["pe_total"]
		avgPb  += evtMap["pb_total"]
		avgTnr += evtMap["tnr_total"]
	}
	avgPe  = avgPe/float64(len(mapIn) )
	avgPb  = avgPe/float64(len(mapIn) )
	avgTnr = avgPe/float64(len(mapIn) )

	avg_mout["pe_total_avg"]  = avgPe
	avg_mout["pb_total_avg"]  = avgPb
	avg_mout["tnr_total_avg"] = avgTnr
	return
}


/**************************************************************************
/* 熵值法:
/* Note: "Entropy weight" is "Degree of distinction". Not weight itself.
/* if X_ij are same, then their e =1 and the row's w =0. means you can remove these
/**************************************************************************/

//Subjective Weight Correction on Weight result
// din: data in;  r: subjective correction revise; w_scr: weight after revise
func WeightSwc(w_in []float64, r []float64)(w_scr []float64){
        //w_ent := WeightEnt(din)
        wxr := []float64{}

	//assert.Equal(T, len(w_in), len(r))
	//assert.NotEqual(T, len(w_in), len(r))
	if (len(w_in) != len(r)){
		fmt.Printf(" len(w_in) is: %v, len(r) is: %v \n", len(w_in), len(r))
		Log.Panicln("<WeightSwc>: len(w_in) != len(r).")
		Log.Fatal("Fatal_Error:<WeightSwc>: len(w_in) != len(r).")
	}

        for i, _ := range w_in{
               wxr = append(wxr, w_in[i]*r[i])
        }
        sum_wxr := sum1d(wxr)
        for i, _ := range w_in{
                w_scr = append(w_scr, (w_in[i]*r[i]) / sum_wxr)
        }
        return
}

// every row is a indicator
func WeightEnt(din [][]float64)(w []float64){
        diff := []float64{}
        //3)计算信息熵冗余度(差异)
        e := CalRowEntropy(din)
        for _, v := range e{
                diff = append(diff, 1 - v)
        }
        //4)计算各项指标的权重
        sum := sum1d(diff)
        for _, v := range diff{
                w = append(w, v/sum)
        }
        return
}


/*******************************************************************
/* step 2: calculate (row) Entropy                                       *
/* Get entropy of every row( every row is a indicator)             *
/*  0 < e < 1(if all element are same, then w is 0).                                                       *
/******************************************************************/
func CalRowEntropy(din [][]float64)(e []float64){
        xweit := []float64{}
        k := 0.0

        if len(din) != 0 {
                k = 1/math.Log(float64(len(din[0]) ))
        }

        _, Y := Norm_2d(din)
        //Y, _ := Norm_ent(din)

        for _, row := range Y{
                //1)计算第j项指标下第i个样本值占该指标的比重：
                xweit = SmpWeitInRow(row)   // xweit is {p_ij, ...}, x weight in row(a indicator)
                mul_p_ln, sum_p_ln := []float64{}, 0.0
                for _, p_ij := range xweit{
                        mul_p_ln = append(mul_p_ln, p_ij * math.Log(p_ij) )
                }
                sum_p_ln = sum1d(mul_p_ln)
                //2)计算第j项指标的熵值
                e = append(e, -k * sum_p_ln)
        }
        return
}


// normalized input 0< v< 1. output: weight in row(a indicator)
func SmpWeitInRow(row []float64)(xweit []float64){
        sumx := sum1d(row)
        for _, v := range row{
                xweit = append(xweit, v/sumx)
        }
        return
}


// do sum for 2-d data, return a slice(colume) of sum
func sum2d(d [][]float64)(sum []float64){
        for _, row := range d{
                sum = append(sum, sum1d(row))
        }
        return
}


func sum1d(d []float64)(sum float64){
        for _, v := range d{
                sum += v
        }
        return
}




/****************************
/* step 1: normalization    *
/****************************/
// Normalize 1d slice data in(every data is a indicator) bases on a input min-max.
func Norm_1d_In(d []float64, min float64, max float64)(d_p, d_n []float64){
	for _, v  := range d{
		//YMIN := 0.0000001
		if v < min  || v > max{
			panic("ERROR: <Norm_1d_In>: Some thing goes wrong. v<min or v > max during do normlization.")
		}
		d_p = append(d_p, (v - min)/(max - min) )
		d_n = append(d_n, (max - v)/(max - min) )
	}
	return
}


func Norm_2d(d [][]float64)(xp, xn [][]float64){
        max_min := []float64{}
        min, max := GetMinMax(d)
        YMAX, YMIN := 0.9999, 0.0001

        for idx, row := range d{
                prow, nrow, max_min_i := []float64{}, []float64{}, 0.0000001  // set 0.0000001 to avoid /0
                if min[idx] != max[idx] {
                        max_min_i = max[idx] - min[idx]
                }else{
                        fmt.Errorf("<Norm_ent>: Note: max[idx]==max[idx]")
                }

                //max_min = append(max_min, max[idx]-min[idx])
                max_min = append(max_min, max_min_i)
                for _, v := range row{
                        prow = append(prow, ((YMAX-YMIN)*(v - min[idx]) )/max_min[idx] + YMIN )
                        nrow = append(nrow, ((YMAX-YMIN)*(max[idx] - v) )/max_min[idx] + YMIN )
                }
                xp, xn = append(xp, prow), append(xn, nrow)
                //Print("@2:xp:", xp, ", xn:", xn)
        }
        return
}


func Norm_EvtsDm(dmEig [][]float64, minCha, maxCha []float64)(dmNorm[][]float64) {
        //fmt.Printf("### <Norm_EvtsDm> dmEig in: %v \n", dmEig )
        for i, eig := range dmEig{
                d_p, _ := Norm_1d_In(eig, minCha[i], maxCha[i])
                dmNorm = append(dmNorm, d_p  )
        }
        //fmt.Printf("### <Norm_EvtsDm> dmNorm out: %v  \n", dmNorm)
        return
}

//----------------------- Get Min/Max ----------------------
// 1d out
func GetMinMax_1d(din []float64)(min, max float64){
	min, max = din[0], din[0]
	for _, v := range din{
		if v > max { max = v}
		if v < min { min = v}
	}
	return
}


// 2d out
func GetMinMax(din [][]float64)(min, max []float64){
        min, max = varinit_2d(din), varinit_2d(din)

        for idx, row := range din{
                for _, e := range row{
                        if e > max[idx]{
                                max[idx] = e
                        }

                        if e < min[idx]{
                                min[idx] = e
                        }
                }
        }
        return
}


// to init var for 2d data process
func varinit_2d(din [][]float64)(o []float64){
        for _,d := range din{
                o = append(o,d[0])
        }
        return
}


// T the matrix
func TranposeDm(dm [][]float64)(dmT [][]float64){
	for n := 0; n < len(dm[0]); n++{
		var rowT []float64
		for _, row := range dm{
			for j, vInRow := range row{
				if j == n{
					rowT = append(rowT, vInRow)
				}
			}
		}
		dmT = append(dmT, rowT)
	}
	return
}
