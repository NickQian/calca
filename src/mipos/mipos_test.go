/*********************************************************
/* test mipos.go
/* ----
/*  License: BSD
/* ----
/* v0.1  init version --- 2021.2.27
/*********************************************************/

package mipos

import (
	"testing"
	dfn "define"
)


func TestMipos(t *testing.T){
	crzSlc, hotSlc, pucSlc, rlxSlc, cur_pos, bi, ti := Mipos("A")
        t.Logf("<TestMipos> res: %v \n", kline)
}

