/*********************************************************
/* filter for the K line (see http://zhuanlan.zhihu.com/p/63386524)
/* ----
/*  License: BSD
/* ----
/* v0.1  init version --- 2020.8.19
/*********************************************************/

package pb

import (
	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/window"
	_ "github.com/mjibson/go-dsp/dsputils"
	. "define"
	_ "cmn"
	)

// Frequency domain Analyze
func KlineFa(kl []float64)(Y_w []complex128){
	fftRes := fft.FFTReal(kl)
	//PlotFa(fftRes, "FreqAmp")

	return fftRes
}


// Kline filter, output the filted result
func FltK(kl []float64)( kf []float64){
	// 1) (y_t)- FFT -> Y_w
	Y_w := KlineFa(kl)

	// 2) Gen win -> W_w
	win := GenWin(len(kl), FFT_INTEREST_POINTS, window.Rectangular)  // Hamming/Hann/Bartlett/Rectangular/FlatTop/Blackman

	// 3) Multi -> YL_w
	YL_w := MultYW(Y_w, win)

	// 4) iFFT -> yL_t
	yL_t := fft.IFFT(YL_w)          // IFFT(x []complex128)[]complex128

	// 5) last process
	kf = GetCmplxAmp(yL_t)

	return
}


