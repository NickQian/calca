/**********************************************************
/* signal process lib used by calca
/*--------------
/* License : BSD
/*--------------
/* v0.1: init version --- 2020.10.15
/*********************************************************/

package pb

import (
	"testing"
	"github.com/mjibson/go-dsp/window"
	
	)


func testGenWin(t *testing.T){
	win := GenWin(10, 3, window.Rectangular)     //L int, l_win int, winFunc func(int)
	PlotKline(win, "This_is_the_WindowGen")

	t.Logf("@@ win generated is: %v ", win)
}


func testMultYW(t *testing.T){
	A := []complex128{2+2i, 2.2+3.3i, 4+4i, 5+5i, 6+6i}
	B := []float64   {1.0,  1.0,      1.0,  0.0,  0.0}
	C := MultYW(A, B)

	t.Logf("<testMultYW> result: %v ", C)
}
