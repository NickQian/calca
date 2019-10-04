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
	"github.com/stretchr/testify/assert"
	"fmt"
)



func TestGetMarket(t *testing.T){
	fmt.Println("-------------------- Test <GetMarket> ------------------------")

        a := new(T_A)
	dicmkt := GetMarket("2019-01-02")
	suc := FilDicToA(dicmkt, a)
	assert.True(t, suc)
	t.Logf("<TestGetMarket>: FilDicToA result: suc %v. dicmkt: %v \n ", suc, dicmkt)
}


func TestMarketUpdate(t *testing.T){
	fmt.Println("-------------------- TestMarketUpdate------------------------")

	a := new(T_A)
	suc := MarketUpdate(a)

	t.Logf("<TestMarketUpdate>: result:  %v \n", suc)

}
