/**********************************************************
/* signal process lib used by calca
/*--------------
/* License : BSD
/*--------------
/* v0.1: init version --- 2020.10.15
/*********************************************************/

package pb


import (
        "math/cmplx"
        "github.com/mjibson/go-dsp/window"
	"github.com/mjibson/go-dsp/dsputils"
        "cmn"
	"github.com/stretchr/testify/assert"
        )



func KlinePreprc(klineRvs []float64)(kl []float64){
        kline := cmn.ReverseSlc(klineRvs)
        kl    = RmvDc(kline) 
        return
}


// remove kline[0]
func RmvDc (kline []float64)(kl []float64){
	for _, v := range kline{
		kl = append(kl, v - kline[0])
	}
	return
}


func GetCmplxAmp(cin []complex128)(fo [] float64){
	for _, v := range cin{
		fo = append(fo, cmplx.abs(v))
	}
	return
}


// generate the window to extract Low Freq
func GenWin(L int, l_win int, winFunc func(int)[]float64))(win []float64){
	win_org := winFunc(l_win)
	pad := make([]float64, L - l_win)
	win = append(win_org, pad...)
	return
}


// note: (comlex128, float64)
func MultYW(Y_w []complex128, W_w []float64)(YL_w []complex128){
	assert.Equal(t, len(Y_w), W_w)

	W_w_c := dsputils.ToComplex(W_w)
	for i:=0; i<len(Y_w);i++{
		YL_w = append(YL_w, Y_w[i]*W_w_c[i]
	}
	return
}
