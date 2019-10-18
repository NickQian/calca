/*******************************************************************
/* math lib for this project
/* ----
/*  License: BSD
/* ----
/* v 0.01   math libs for policy B  --- 2019.1.26
/******************************************************************/

package cmn

import(
        "math"
        "fmt"
)

// simple average
func SimpleAvg(din []float64)(avg_out float64){
	var avg float64 = 0.0

	for _, val := range din{
		avg += val
        }
	avg = float64(avg)/float64(len(din))

	return avg
}


/**************************************************************************
/* 熵值法:
/* Note: "Entropy weight" is "Degree of distinction". Not weight itself.
/* if X_ij are same, then their e =1 and the row's w =0. means you can remove these
/**************************************************************************/
func WeightEntSwc(din [][]float64, r []float64)(w []float64){
        w_ent := WeightEnt(din)
        wxr := []float64{}
        for i, _ := range w_ent{
               wxr = append(wxr, w_ent[i]*r[i])
        }
        sum_wxr := sum1d(wxr)
        for i, _ := range w_ent{
                w = append(w, (w_ent[i]*r[i]) / sum_wxr)
        }
        return
}

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
/* step 2: calculate Entropy                                       *
/* Get entropy of every row( every row is a indicator)             *
/*  0 < e < 1(if all element are same, then w is 0).                                                       *
/******************************************************************/
func CalRowEntropy(din [][]float64)(e []float64){
        xweit := []float64{}
        k := 0.0

        if len(din) != 0 {
                k = 1/math.Log(float64(len(din[0]) ))
        }

        _, Y := Norm_ent(din)
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
func Norm_ent(d [][]float64)(xp, xn [][]float64){
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
