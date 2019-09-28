/*********************************************************
/* test qif.go
/* ----
/*  License: BSD
/* ----
/* v0.01  init version --- 2019.9.18
/*********************************************************/

package qif

import (
         "testing"
         . "define"
//	"github.com/stretchr/testify/assert"
	"fmt"
)



func TestHavaLook(t *testing.T){
	fmt.Println("-------------------- Test HaveaLook ------------------------")

        a := new(T_A)
        a.Cmv.Cmv_total = 19999.0
        t.Logf("a.Cmv.Cmv_total is: %v ", a.Cmv.Cmv_total)
	suc := HavaLook("2019-9-12", a)
	t.Logf("<TestHavalook>: %v. start assert.... the panic may caused by assert. ", suc)

}


func TestMarketUpdate(t *testing.T){
	fmt.Println("-------------------- TestMarketUpdate------------------------")


	a := new(T_A)
	a.Cmv.Cmv_total = 29999.0
	t.Logf("a.Cmv.Cmv_total is: %v ", a.Cmv.Cmv_total) 

	suc := MarketUpdate(a)
	t.Logf("<TestMarketUpdate>: %v. start assert.... the panic may caused by assert. ", suc)
//	assert.True(t, suc)
}
