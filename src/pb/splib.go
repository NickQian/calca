/**********************************************************
/* signal process lib used by calca
/*--------------
/* License : BSD
/*--------------
/* v0.1: init version --- 2020.10.15
/*********************************************************/

package pb


import (
	"math"
        "math/cmplx"
        _ "github.com/mjibson/go-dsp/window"
	"github.com/mjibson/go-dsp/dsputils"
        _ "cmn"
        "fmt"
        dfn "define"
        )


// return K line which has been pre-processed
func KlinePreprc(kline []float64)(kp []float64){
        fmt.Printf("Info: <KlinePreprc> len of kline: %v, PAD_LEN: %v \n", len(kline), dfn.FFT_FLT_PAD_LEN )

        //kln     := GetLnK(kline)
        //kl, _   := RmvIniDc(kln)
        kl, _   := RmvIniDc(kline)

	//kp = kl
	kp = PadHeadEnd(kl, 0,  //int(dfn.FFT_FLT_PAD_LEN * float64(len(kline)) ),
	                    int(dfn.FFT_FLT_PAD_LEN * float64(len(kline)) )  )
        return
}





// remove kline[0] or kline[-1]
func RmvIniDc (kline []float64)(kl, kle []float64){
	for _, v := range kline{
		kl  = append(kl, v - kline[0]  )
		kle = append(kl, v - kline[len(kline)-1]  )
	}
	return
}

// remove the "channel" DC: kl with "ac" only
func RmvChDc(kl []float64) (klAc []float64){
	return
}

func GetLnK(kl []float64) (kln []float64){
	for _, v := range kl{
		kln = append(kln, math.Log(v) )
	}
	return
}


func PadHeadEnd(k []float64, L_pad_Head, L_pad_End int)(kpad []float64){
	padH := make([]float64, L_pad_Head)
	padE := make([]float64, L_pad_End )
	for i, _ := range padH{
		padH[i] = k[0];
	}
	for i, _ := range padE{
		padE[i] = k[len(k) -1]
	}

	H_k  := append(padH, k...)
	kpad =  append(H_k, padE...)
	fmt.Println("Info: k padded: ", kpad)

	return
}

/*
// abs, then div Len()
func GetCmplxAmp(cin []complex128)(fo [] float64){
	for _, v := range cin{
		fo = append(fo, cmplx.Abs(v) / (float64(len(cin)) /2.0) )
	}
	return
}


func GetFreqAxis(cin []complex128)(faxis []float64){
	Fs := 1
	for i, _ := range cin{
		faxis = append(faxis, float64(i*Fs)/float64(len(cin))  )
	}
	return
}
*/


// generate the window to extract Low Freq
func GenWin(L int, l_win int, winFunc func(int)[]float64 )(win []float64){
	win_org := winFunc(l_win)
	pad := make([]float64, L - l_win)
	win = append(win_org, pad...)

	fmt.Printf("@@@ win generated is: %v  \n ", win)
	PlotSlicef(win_org, "time", "window", "window_gen.png")

	return
}


// note: (comlex128, float64)
func MultYW(Y_w []complex128, W_w []float64)(YL_w []complex128){
	//assert.Equal(t, len(Y_w), len(W_w))
	if len(Y_w) != len(W_w){
		panic("<splib>: len(Y_w) != len(W_w). Plese check the inputs.")
	}

	W_w_c := dsputils.ToComplex(W_w)
	for i:=0; i<len(Y_w);i++{
		YL_w = append(YL_w, Y_w[i]*W_w_c[i] )
	}
	return
}
