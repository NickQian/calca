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


// Ix: index
func TestGetIxKline(t *testing.T){
	kline := GetIxKline("000001.SH","20190103", "20210118")
	t.Logf("<TestGetIxKline> res: %v \n", kline)
}


// Is: single stock
func TestGetIsKline(t *testing.T){
        kline := GetIsKline("600109.SH","20190103", "20190318")
        t.Logf("<TestGetIsKline> res: %v \n", kline)
}



func TestGetMarket_raw(t *testing.T){
	fmt.Println("-------------------- Test <GetMarket> ------------------------")

        a := new(T_A)
	dicmkt := GetMarket_raw("2019-01-02")
	suc := FilDicToA(dicmkt, a, "2019-01-02")                //dicmkt map[string]float64, a *T_A, tag string)
	assert.True(t, suc)
	t.Logf("<TestGetMarket>: FilDicToA result: suc %v. dicmkt: %v \n ", suc, dicmkt)
}

/*
func TestMarketUpdate(t *testing.T){
	fmt.Println("-------------------- TestMarketUpdate------------------------")

	a := new(T_A)
	suc := MarketUpdate(a)

	t.Logf("<TestMarketUpdate>: result:  %v \n", suc)

}
*/

func TestGetTradeDays(t *testing.T){
	fmt.Println("-------------------- TestGetTradeDays------------------------")
	days := GetTradeDays("2019-09-12", 5)
	t.Logf("<TestGetTradeDays>: %v \n", days)
}


