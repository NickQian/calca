// 背景：
// 目前市场上有4种策略：1)价值投资策略；2)量化分析策略；3)技术分析策略； 4) 行为金融策略
// 3)本质是讲究概率经验；基本认为是无效策略，被2）鄙视为“占星术”，
// 但是4)+3)是市场上“高人”们在用；
// 2)目前基本还不挣钱原因是”没有方程“或者说没有正确的方程；等待数理分析突破的一天。
// 本策略A:
//   策略A基于 1) & 4)。基础模型为“称赌模型"(称重机混合赌场模型). (基本理念来源于涛哥的“1-理念交流群”，致谢！)
// -------------------------------------------------------------------------------------------------------------------
// v0.01 - policy A得出的mix model rate% & index position  --- 2018.8.30

package pa

import (
        //""
        //""
        "fmt"
        )

// --------------------------估值函数 ----------------------------
// 0: 全仓买入;  100: 卖出（清空）
func pa_mipos()(){

}


//(*-----------------------底、顶典型特征(状态)-----------------------*)
bot = f1(bot_market, bot_emo , bot_policy);        // 市场底，政策底，情绪底，
bot_trade = f3(pe_avg, pe_extreme);
top = f2(pe_avg, vol_trade, index_emo, pe_50, pe_300,pe_ms, pe_gm);


//(*-------------------- cindex ---------------------*)
//计算cindex的原因是，因为情绪的原因，单只股票正负泡沫高点也是composite index的正负泡沫高点。 
//      在其它时刻，针对篮子里的三支票，以及真正有价值的票，composite index有失真现象，因而参考价值不大。
cindex =  f3(gdp, roi, policy, );  // composite index 


/******************************** cindex should be *************************************
/* 计算cindex的原因是，因为情绪的原因，单只股票正负泡沫高点也是composite index的正负泡沫高点。 
/*      在其它时刻，针对篮子里的三支票，以及真正有价值的票，composite index有失真现象，因而参考价值不大。
/***************************************************************************************/

cindex_should =  f3(gdp, roi, policy, );  // composite index 


//(*----------------- Bubble计算 --------------------*)
bubble = price - value;
price = pe_current;
value = pe_avg;



// ====================== 三碗面 ================================
// 基本面决定趋势，技术面决定位置，情绪面决定日常波动区间

