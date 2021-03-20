
/*********************************************************
/* ----
/*  License: BSD
/* ----
/* v0.1  init version --- 2019.7.12
/*********************************************************/

package plotit


import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	. "cmn"
	 "fmt"
	dfn "define"
)

//-------------------------- Kline -------------------------------

func PlotKl(kl []float64, fn string) bool{
	PlotSlicef(kl, "time", "Amp", fn)
	return true
}


// title: eg."price.png"
func PlotSlicef(slc []float64, xLable, yLable, fn string) bool{
	if len(slc) == 0{ panic(ErrEmptyNoItem)   }
	
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

//---------------------- fft --------------------------------
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
        if len(x) == 0{ panic(ErrEmptyNoItem)   }
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


//------------------ plot mipos result -----------------------
func PlotMipos(cw dfn.T_Mipos) bool {
        fmt.Println("Info: <Plot Mipos> starting...")

        p, _ := plot.New()
        pts_crz, pts_hot, pts_puc, pts_rlx, pts_pos := makeXY_Evts(cw)          // the difference
        plotutil.AddLinePoints(p, pts_crz, pts_hot, pts_puc, pts_rlx, pts_pos )
        fmt.Printf("#### pts_crz: %v, pts_hot:%v, pts_puc:%v, pts_rlx:%v, pts_pos:%v  \n", pts_crz, pts_hot, pts_puc, pts_rlx, pts_pos )

        p.Title.Text   = "mipos_result"
        p.X.Label.Text = "Day"
        p.Y.Label.Text = "cw result"
        p.Save(8*vg.Inch, 8*vg.Inch, "res_mipos.png")
        fmt.Println("Info: <Plot Mipos> Finished...")
        return true
}


func makeXY_Evts(cw dfn.T_Mipos)(pts_crz, pts_hot, pts_puc, pts_rlx, pts_pos plotter.XYs){
        pts_crz = make(plotter.XYs, len(cw.TopCrz) )  //dfn.LEN_MIPOS_K)
        pts_hot = make(plotter.XYs, len(cw.TopHot) )
        pts_puc = make(plotter.XYs, len(cw.BotPuc) )
        pts_rlx = make(plotter.XYs, len(cw.BotRlx) )
        pts_pos = make(plotter.XYs, len(cw.Poslc)  )

        for i, _ := range cw.Poslc{
                if ( i < len(cw.TopCrz) ){
                        pts_crz[i].X, pts_crz[i].Y = float64(i), cw.TopCrz[i]
                        fmt.Printf("#### cw.TopCrz: %v, i:%v, pts_crz:%v  \n", cw.TopCrz, i, pts_crz )
                }

		if ( i < len(cw.TopHot) ){
                        pts_hot[i].X, pts_hot[i].Y = float64(i), cw.TopHot[i]
                }

		if ( i < len(cw.BotPuc) ){
                        pts_puc[i].X, pts_puc[i].Y = float64(i), cw.BotPuc[i]
                }


		if ( i < len(cw.BotRlx) ){
                        pts_rlx[i].X, pts_rlx[i].Y = float64(i), cw.BotRlx[i]
                }


		if ( i < len(cw.Poslc) ){
                        pts_pos[i].X, pts_pos[i].Y = float64(i), cw.Poslc[i]
                }

        }
        return
}

