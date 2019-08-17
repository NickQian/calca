# v0.01 web top --- 2019.1.15  

package simtrade

import (
        "fmt"
)

type StageOfMkt  int

const (
    Init  StageOfMkt = 0 +iota    //
    Idle
    NiuShiKou                     // 建仓
    NiuShiKou_                    // 建仓完毕，等
    NiuWangMiao                   // 卖点
    NiuWangMiao_                  // 卖完空，等
    DongMenDaQiao                 // 看到跳水
    DongMenDAQiao_                // 跳水中，等
    ChunXiLu                      // 落地响;
    CHunXiLu_                     // 落地后春天到，程序化筛选拿大票
)

type state struct(
    State string
    NexState string
)


func simtrade()(acc float64){    
    
    logfile, err := os.OpenFile("run.log",os.O_APPEND|O_CREATE, 666)
	logger := &log.New(logfile, "", log.Ldate | log.Ltime)
    
    calres := get_calca_res(Fn_Res_Calc)    
    stateReg := &state{State:"init", NexState:"init"}
    
    determineState(stateReg, calres, logger)
    sta2trade(stateReg)
    
}




// state machine
func determineState(stateReg *state, calres CalRes, logger *log.Logger)(res bool){
    (*stateReg).State = (*stateReg).NexState
    
    switch (*stateReg).State{
    case Init:  
        (*stateReg).State = Idle
        (*stateReg).NexState = Idle
    case Idle, ChunXiLu_:
        if cal_btm.bi == 100{
            (*stateReg).NexState = NiuShiKou
        }            
    case NiuShiKou:
        (*stateReg).NexState = NiuShiKou_
    case NiuShiKou_:
        if cal_btm.ti == 100{
            (*stateReg).NexState = NiuWangMiao
        }
    case NiuWangMiao:
        (*stateReg).NexState = NiuWangMiao_
    case NiuWangMiao_:
        if cal_btm.mix_cw[0] == 0{
            (*stateReg).NexState = DongMenDaQiao_
        } 
    case DongMenDaQiao:
        logger.
    case DongMenDaQiao_:
        if cal_btm.bi == 100{
            (*stateReg).NexState = ChunXiLu
        }      
    case ChunXiLu:  
        (*stateReg).NexState = ChunXiLu_    
    default:
        logger.Fatal("Fatal Error: DetermineState: run in default")
    }
    
   
}



// trade bases on state
func sta2trade(*stateReg state)(res bool){

    switch (*stateReg).State{
    case Init:  
        
    case Idle, ChunXiLu_:
        
    case NiuShiKou:
        
    case NiuShiKou_:
        
    case NiuWangMiao:
        
    case NiuWangMiao_:
        
    case DongMenDaQiao_:
        
    case ChunXiLu:  
        
    default:
        
    }
    
    
}










