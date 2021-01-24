
/*********************************************************
/* ----
/*  License: BSD
/* ----
/* v0.01  init version --- 2019.7.12
/*********************************************************/

package pb


import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"cmn"
	"fmt"
)

//------------- Kline --------------

func PlotKl(kl []float64, fn string) bool{
	//kl := RmvDc(kline)
	PlotSlicef(kl, "time", "Amp", "Kl")
	return true
}


// title: eg."price.png"
func PlotSlicef(slc []float64, xLable, yLable, fn string) bool{
	if len(slc) == 0{ panic(cmn.ErrEmptyNoItem)   }
	
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
	pts = make(plotter.XYs, len(kl))
	for i, v := range kl{
	
		pts[i].X, pts[i].Y = float64(i), float64(v)
	}
	return
}

//--------------- fft ------------------
// complex128 as input
func PlotFa(fa []complex128, fn string) bool{
	Afreq := GetCmplxAmp(fa)
	PlotSlicef(Afreq, "freq", "Amp", fn)
	return true
}


func PlotSliceXY(x, y []float64, fn string)bool{
	fmt.Print("<PlotSliceXY>")
	return true
}


//-------------- ifft ------------------
