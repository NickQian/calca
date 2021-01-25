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
	"cmn"
	"fmt"
	"qif"
	)


func KlinePlot(indexCode, startDay, endDay string)bool{
        kline := qif.GetKline( indexCode, startDay, endDay )
        suc := PlotKl(kline, "Kline_org.png")
        kl = KlinePreprc(kline)
        suc = PlotKl(kl, "Kline_DcRmved.png")
        if !suc{fmt.Printf("<KlinePlot>: %v \n", cmn.ErrPlotFail)   }
	return suc
}

// Frequency domain Analyze
func KlineFa(kl []float64)(Y_w []complex128){
	fftRes := fft.FFTReal(kl)
	//PlotFa(fftRes, "FreqAmp"
	return fftRes
}


// Kline filter, output the filted result
func FltK(kl []float64)( kf []float64){
	// 1) (y_t)- FFT -> Y_w
	Y_w := KlineFa(kl)
	PlotFa(Y_w[:30], "testFltK_Fa.png")

	// 2) Gen win -> W_w
	win := GenWin(len(kl), FFT_INTEREST_POINTS, window.Rectangular)  // Hamming/Hann/Bartlett/Rectangular/FlatTop/Blackman
	//fmt.Printf("### 1: win###: %v  \n", win)

	// 3) Multi -> YL_w
	YL_w := MultYW(Y_w, win)
	fmt.Printf("### 2 YL_w###:  %v   \n", YL_w)

	// 4) iFFT -> yL_t: complex out
	yL_t := fft.IFFT(YL_w)          // IFFT(x []complex128)[]complex128
	//fmt.Printf("### 3: yL_t###: v%   \n", yL_t)

	// 5) last process
	kf = GetCmplxAmp(yL_t)

	return
}


