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
	//"define"
)


func TestMipos(t *testing.T){
	cw  := Mipos("A")
        t.Logf("<TestMipos> res: %v \n", cw)
}

