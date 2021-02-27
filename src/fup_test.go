/**************************************************************
/* ----
/*  License: BSD
/* ----
/* v0.1  . --- 2019.10.21
/**************************************************************/

package main


import (
        "testing"
        dfn "define" 
)


func TestDistilEvt(t *testing.T){
        res := DistilEvt(dfn.FN_DATE_TOP_CRZ)        // FN.DATE_BOT_RLX/PUC  TOP_CRZ/HOT
        t.Logf(" %t ", res)
}


 
