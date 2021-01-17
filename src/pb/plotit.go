
/*********************************************************
/* ----
/*  License: BSD
/* ----
/* v0.01  init version --- 2019.7.12
/*********************************************************/

package pb


import (
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

//------------- Kline --------------

func PlotKl(kline []float64, fn string) bool{
	kl := RmvDc(kline)
	PlotSlicef(kl, "time", "Amp", "Kl")
}


// title: eg."price.png"
func PlotSlicef(slc []float64, xLable, yLable, fn string) bool{
	p, _ := plot.New()
	points := makeXY_Kline(slc)
	plotutil.AddLinePoints(p, points)

	p.Title.Text   = fn
	p.X.Label.Text = xLable
	p.Y.Label.Text = yLable
	p.Save(4*vg.Inch, 4*vg.Inch, fn)
	return true
}


func makeXY_Kline(kl []float64)(pts plotter.XYs){
	for i, v := range kl{
		pts.X, pts.Y = i, v
	}
	return
}

//--------------- fft ------------------
// complex128 as input
func PlotFa(fa []complex128, string fn) bool{
	Afrq := GetCmplxAmp(fa)
	PlotSlicef(Afreq, "freq", "Amp", fn)
	return true
}


func PlotSliceXY(x, y []float64, fn string)bool{

	return true
}


//-------------- ifft ------------------
