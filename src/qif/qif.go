/**********************************************************************************************************
/*  Quantitative interface:
/*  可选： Wind(没找到linux接口文件), 京东JDQuant(2018.12关闭), MindGo(同花顺),优矿(通联数据Uqer.io), 米筐('高质量数据免费使用'), 
/*           JoinQuant(聚宽,已申请1年)，BigQuant, tushare等, 国外Quantopian, WorldQuant
/*  使用gRPC 还是cgo？  cgo: github.com/sbinet/go-python
/* ----
/*  License: BSD
/* ----
/*  0.1.1: 放弃使用interface, 采用struct里面放函数指针的办法 - Nick cKing
/*  0.1.0: init version - 2019.1 -  Nick cKing
/************************************************************************************************************/

package qif


import(
    "github.com/sbinet/go-python"
    "fmt"
)

//----------------------- 定义--------------------------


//type QIF interface {
type QIF struct{
    getCurPE()      func (pa *T_A)getCurPE() (bool)
    getCurPB()      func (pa *T_A)getCurPB() (bool)
    getCurMtss()    func (pa *T_A)getCurMtss() (bool)
    getCurVol()     func (pa *T_A)getCurVol() (bool)
    getCurTnr()     func (pa *T_A)getCurTnr() (bool)   //turnover ratio
    getSingle()     func (pa *T_A)getSingle() (bool)  
}

var (
        ErrPasswd = errors.New("qif: password or user name not correct")
        ErrNoDataReturn = errors.New("qif: no data return from this qif")
        ErrNoItem = errors.New("there's no this item in qif")
)



//------ func实现 -------
func (pa *T_A)getCurPE() (bool) {
    pelist, suc := goCallpy("getCurPE") 
    pa.pe.total = pelist
    return suc 
}



func (pa *T_A) getCurPB() (bool) {
    return goCallpy("getCurTor")
}

func (pa *T_A) getCurVol()(bool){
    return goCallpy("getCurVol")
}

func (pa *T_A) getCurTnr()(bool){
    return goCallpy("getCurTor")
}


func (pa *T_A) getCurMtss() (bool) {
    return goCallpy("getCurMtss")
}




//---------------------- python 桥接 ------------------------------------

type T_go2py struct{
    Str2Py func(string)          (*python.PyObject)
    Py2Str func(*python.PyObject)(string)
}
    //PyStr := python.PyString_FromString
    //GoStr := python.PyString_AS_STRING

func go2pyInit()(*python.PyObject){
    if err:= python.Initialize(); err != nil{
        panic(err.Error())
    }
    pyBrg := &T_go2py{Str2Py: python.PyString_FromString,
                      Py2Str: python.PyString_AsString} 
    }
    
    // (1) (2) --- select Qif的公司，选一家，注释掉其它家
    switch QIF_VENDOR {  
        case "JQ":      module := pyBrg.ImportModule("./", "if_jq")
        case "UQ":      module := pyBrg.ImportModule("./", "if_uq")
        case "RQ":      module := pyBrg.ImportModule("./", "if_rq")
        case "BQ":      module := pyBrg.ImportModule("./", "if_bq")
        case "tushare": module := pyBrg.ImportModule("./", "if_ts")
        default:  Log.Fatal("Err: [qif]: no Qif configurated.")
    }    
    
    return module    
}


func (pyBrg *T_go2py)ImportModule(dir, name string)(*python.PyObject){
    //(1) 
    module := python.PyImport_ImportModule("sys")  
    path := module.GetAttrString("path")
    python.PyList_Insert(path, 0, pyBrg.Str2Py(dir))
    //(2)
    return python.PyImport_ImportModule(name)     //func PyImport_ImportModule(name string) *PyObject
}


func (pymodule *PyObject)goCallpy(defname string, args ... string){
    f := pymodule.GetAttrString(defname)
    
    //func PyTuple_New(sz int) *PyObject
    argv := python.PyTuple_New(len(args))          //Py_BuildValue，PyTuple_SetItem
    for i, value := range args{
        python.PyTuple_setItem(argv, i, pyBrg.Str2Py(value)) //func PyTuple_SetItem(self *PyObject, pos int, o *PyObject) error
    }
    
    res := f.Call(argv, python.Py_None)
    return PyBrg.Py2Str(res)
}



// ----------------- operate with python api ---------------------
var Qif QIF   // interface
var A = nil  

func init(){      
    //switch QIF_VENDOR {  
    //    case "JQ":     A := &Q_jq{}
    //}
    //Qif = A     // 接a 
    
    // init go_py bridge    
    go2pyInit()
    
    // fetch today data once to update A status
    if e := Qif.getCurPE(); e != nil{
            //log.Fatalln("fatal Err: Qif not get PE data ") 
            panic("Error: Qif not get PE data")
    }
    if e := Qif.getCurPB(); e != nil{
            panic("Error: Qif not get PB data")
    }
    if e := Qif.getCurMtss(); e != nil{
            panic("Error: Qif not get Mtss data")
    }
    if e := Qif.getCurVol(); e != nil{
            panic("Error: Qif not get Vol data")
    }
    
}












