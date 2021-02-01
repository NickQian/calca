
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
	dfn "define"
)

//------------- Kline --------------

func PlotKl(kl []float64, fn string) bool{
	PlotSlicef(kl, "time", "Amp", fn)
	return true
}


// title: eg."price.png"
func PlotSlicef(slc []float64, xLable, yLable, fn string) bool{
	if len(slc) == 0{ panic(cmn.ErrEmptyNoItem)   }
	
	p, _ := plot.New()
	points := makeXY_slcIn(slc)
	plotutil.AddLinePoints(p, points)

	p.Title.Text   = fn
	p.X.Label.Text = xLable
	p.Y.Label.Text = yLable
	p.Save(8*vg.Inch, 8*vg.Inch, fn)
	return true
}


func makeXY_slcIn(slc []float64)(pts plotter.XYs){
	pts = make(plotter.XYs, len(slc))
	for i, v := range slc{
		//fmt.Printf("##### len(kl): %v, i:%v, v:%v  \n", len(slc), i, v)
		pts[i].X, pts[i].Y = float64(i), float64(v)
	}
	return
}

//--------------- fft ------------------
// complex128 as input
func PlotFa(fa []complex128, fn string) bool{
	fmt.Println("Info: <PlotFa> starting...")
	A_freq := GetCmplxAmp(fa)
	F_axis := GetFreqAxis(fa)
	PlotSliceXY(F_axis, A_freq, "freq", "Amp", fn)
	fmt.Println("Info: <PlotFa> Finished...")

	return true
}


func PlotSliceXY(x, y []float64, xLable, yLable, fn string) bool{
        if len(x) == 0{ panic(cmn.ErrEmptyNoItem)   }
        p, _ := plot.New()
        points := makeXY_XYIn(x, y)          // the difference
        plotutil.AddLinePoints(p, points)

        p.Title.Text   = fn
        p.X.Label.Text = xLable
        p.Y.Label.Text = yLable
        p.Save(8*vg.Inch, 8*vg.Inch, fn)
        return true
}


func makeXY_XYIn(x, y []float64)(pts plotter.XYs){
        pts = make(plotter.XYs, len(x))
        for i, _ := range x{
        	if ( i < dfn.FFT_FA_INTEREST_PTS ){
                        //fmt.Printf("#### len(kl): %v, x:%v, y:%v  \n", len(x), x[i], y[i] )
			pts[i].X, pts[i].Y = x[i], y[i]
		}
        }
        return
}

//-------------- ifft ------------------
