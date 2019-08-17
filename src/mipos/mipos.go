// -----------------------------------------------------------------------------------
// v0.01 - policy A得出的mix model rate% & index position  --- 2018.8.30

package mipos



import (
        "pa"
        //"pb"
        //"pc"
        "cmn"
        "fmt"
        )

// --------------------------估值函数 ----------------------------
// 0: 全仓买入;  100: 卖出（清空）
func policya_mipos()(){

}


//(*-----------------------底、顶典型特征(状态)-----------------------*)
bot = f1(bot_market, bot_emo , bot_policy);        // 市场底，政策底，情绪底，
bot_trade = f3(pe_avg, pe_extreme);
top = f2(pe_avg, vol_trade, index_emo, pe_50, pe_300,pe_ms, pe_gm);


//(*-------------------- cindex ---------------------*)
//计算cindex的原因是，因为情绪的原因，单只股票正负泡沫高点也是composite index的正负泡沫高点。 
//      在其它时刻，针对篮子里的三支票，以及真正有价值的票，composite index有失真现象，因而参考价值不大。
cindex =  f3(gdp, roi, policy, );  // composite index 


//(*----------------- Bubble计算 --------------------*)
bubble = price - value;
price = pe_current;
value = pe_avg;



// ====================== 三碗面 ================================
// 基本面决定趋势，技术面决定位置，情绪面决定日常波动区间

