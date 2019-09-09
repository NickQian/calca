/**********************************************************************************************************
/*  Quantitative interface:
/*  可选： Wind(没找到linux接口文件), 京东JDQuant(2018.12关闭), MindGo(同花顺),优矿(通联数据Uqer.io), 米筐('高质量数据免费使用'), 
/*           JoinQuant(聚宽,已申请1年)，BigQuant, tushare等, 国外Quantopian, WorldQuant
/*  使用gRPC 还是cgo？  cgo: github.com/sbinet/go-python
/* ----
/*  License: BSD
/* ----
/*  0.1.1: 放弃使用interface, 采用struct里面放函数指针的办法 - 2019.6  Nick cKing
/*  0.1.0: init version - 2019.1 -  Nick cKing
/************************************************************************************************************/

package qif


import(
        "github.com/sbinet/go-python"
        "fmt"
        "time"
        . "define"
        "errors"
        "strings"
)


//----------------------- 定义& Var --------------------------
//type QIF interface {

type QIF struct{
        GetAstat   func(a *T_A)(bool)           //getAstat(a *T_A) (bool)            // statistics & status
        GetApara   func(a *T_A)(bool)           //getApara(a *T_A) (bool)
        GetPE      func(a *T_A)(bool)           //getPE(a *T_A) (bool)
        GetPB      func(a *T_A)(bool)           //getPB(a *T_A) (bool)
        GetVolr    func(a *T_A)(bool)           //getVol(a *T_A) (bool)
        GetTnr     func(a *T_A)(bool)           //getTnr(a *T_A) (bool)   //turnover ratio
        GetMtsr    func(a *T_A)(bool)           //getMtss(a *T_A) (bool)

        GetSingle  func(code int)(bool)         //getSingle(code int) (bool)
}


var (
        ErrPasswd       = errors.New("qif: password or user name not correct")
        ErrNoDataReturn = errors.New("qif: no data return from this qif")
        ErrNoItem       = errors.New("there's no this item in qif")
)


//var Qif QIF   // interface
//var A = nil

var now = time.Now()
//var now time.Time = time.Now()

var today string

//var PyBrg *T_go2py
var PyBrg *T_go2py = &T_go2py{Str2Py: python.PyString_FromString,
                              Py2Str: python.PyString_AsString,
                             }


//------------------------ func实现 ------------------------------------

func MarketUpdate(a *T_A) (suc bool){
        AstatDic, _  := goCallpy("getMarket", today)
        a.Cmv.Cmv_total = AstatDic["cmv_total"]
        return true
}


// have a look
func HavaLook(day string, a *T_A) (bool){
        peDic   := GetPE(day)
        pbDic   := GetPB(day)
        volrDic := GetVolr(day)
        tnrDic  := GetTnr(day)
        mtsrDic := GetMtsr(day)
        a.Pe.Pe_sh = peDic.Pe_sh
        a.Pb.Pb_sh = pbDic.Pb_sh
        a.Volr.Volr = volrDic.Volr
        a.Tnr.Tnr = tnrDic.Tnr
        a.Mtsr.Mtsr_total = mtsrDic.Mtsr_total
        return true
}


func GetPE(day string) (pe T_pe) {
        peMap, _ := goCallpy("getPE", day)
        pe.Pe_sh = peMap["pe_sh"]
        return pe
}


func GetPB(day string) (pb T_pb) {
        pbMap, _ := goCallpy("getPB", day)
        pb.Pb_total = pbMap["pb_total"]
        return
}


func GetVolr(day string)(o T_volr){
        //goCallpy("getVol", day)
        return
}


func GetTnr(day string)(o T_tnr){
        //goCallpy("getTnr", day)
        return
}


func GetMtsr(day string) (o T_mtsr) {
        //goCallpy("getMtss", day)
        return
}



//---------------------- python 桥接 ------------------------------------

type T_go2py struct{
        Str2Py func(string)(*python.PyObject)     // func PyString_FromString(v string) *PyObject
        Py2Str func(*python.PyObject)(string)
}

//PyStr := python.PyString_FromString
//GoStr := python.PyString_AS_STRING


func go2pyInit()(suc bool){     //pymodule *python.PyObject){        //, PyBrg *T_go2py){
        suc = false
        if err:= python.Initialize(); err != nil{
                panic(err.Error())
        }else{
               fmt.Println("info: python.Initialize done! ")
        }

        /* PyBrg = &T_go2py{Str2Py: python.PyString_FromString,
                         Py2Str: python.PyString_AsString,
                        }

        //--- select Qif的公司，选一家，注释掉其它家
        switch QIF_VENDOR {
                case "JQ":      PyModule = ImportModule("./", "if_jq")
                case "UQ":      PyModule = ImportModule("./", "if_uq")
                case "RQ":      PyModule = ImportModule("./", "if_rq")
                case "BQ":      PyModule = ImportModule("./", "if_bq")
                case "tushare": PyModule = ImportModule("./", "if_ts")
                default:  //Log.Fatal("Err: [qif]: no Qif configurated.")
                        panic("Err: Wrong (QIF_VENDOR) value.")
        }
        */

        suc = true
        fmt.Println("<go2pyInit> Init & import module done!  ")

        return
}




func ImportModule(dir, name string)(*python.PyObject){
        //(1) add dir into python env "sys.path"
        sysModule := python.PyImport_ImportModule("sys")                           // func PyImport_ImportModule(name string) *PyObject
        path := sysModule.GetAttrString("path")                                    // path is ['', '', '']
        python.PyList_Insert(path, 0, PyBrg.Str2Py(dir) )                          // func PyList_Insert(self *PyObject, index int, item_to_insert *PyObject)
        //(2) import *.py module in the dir which contain this *.py file
        return python.PyImport_ImportModule(name)
}


//func (pymodule *python.PyObject)goCallpy(defname string, args ... string){
func goCallpy(defname string, args ... string)(omap map[string]float64, suc bool){
        suc = false
        var PyModule *python.PyObject                                               // = new(python.PyObject)


        sysModule := python.PyImport_ImportModule("sys")                            // func PyImport_ImportModule(name string) *PyObject
        path := sysModule.GetAttrString("path")                                     // path is ['','']
        fmt.Println("sys.path is(before):", path)
        python.PyList_Insert(path, 0, PyBrg.Str2Py("/home/nk/calca/src/qif") )      // func PyList_Insert(self *PyObject, index int, item_to_insert *PyObject)
        fmt.Println("sys.path is(after):", path)

        switch QIF_VENDOR {
        case "JQ":      PyModule = python.PyImport_ImportModule("if_jq") //"./", "if_jq")
                        fmt.Print("Info: QIF selected is JQ. \n "  )
        case "UQ":      PyModule = python.PyImport_ImportModule("if_uq") //"./", "if_uq")
        case "RQ":      PyModule = python.PyImport_ImportModule("if_rq") //"./", "if_rq")
        case "BQ":      PyModule = python.PyImport_ImportModule("if_bq") //"./", "if_bq")
        case "tushare": PyModule = python.PyImport_ImportModule("if_ts") //"./", "if_ts")
        default:  //Log.Fatal("Err: [qif]: no Qif configurated.")
                  panic("Err: Wrong (QIF_VENDOR) value.")
        }

        if PyModule == nil{
                panic("##$$%%^&&&##: Error: import result nil.")
        }

        fmt.Printf(" @(1)->PyModule: %v, [MODULE]repr()= %s ", PyModule,  PyBrg.Py2Str(PyModule.Repr() ) )
        fmt.Println(" @(2), defname is:    ", defname)

        f := PyModule.GetAttrString(defname)
        fmt.Printf("@(3)->f: %v", f)

        //func PyTuple_New(sz int) *PyObject
        argv := python.PyTuple_New(len(args))                        //Py_BuildValue，PyTuple_SetItem
        for i, value := range args{
                python.PyTuple_SetItem(argv, i, PyBrg.Str2Py(value)) //func PyTuple_SetItem(self *PyObject, pos int, o *PyObject) error
        }

        resDict := f.Call(argv, python.Py_None)
        keyObj := python.PyDict_Keys(resDict)
        keys := PyBrg.Py2Str(keyObj)
        for i, key := range keys{
                //value := python.PyDict_GetItemString(resDict, key)
                //omap[key] = PyBrg.Py2Str(value)
                fmt.Println("@@@ <goCallpy> call map result :", i, key)
                suc = true
        }

        return
}



// ----------------- operate with python api ---------------------

func init(){
        todayFull := now.Format(TIME_LAYOUT_STR)
        todaySlice := strings.SplitAfter(todayFull, " ")
        today = todaySlice[0]
        fmt.Println("<qif:init>: today:", today)

        // init go_py bridge
        go2pyInit()
        //pymodule  = go2pyInit()

        QifLogin()

        // fetch today data once to update A status
        if _, e := GetCurPE(); e != nil{
                //log.Fatalln("fatal Err: Qif not get PE data ")
                panic("Error: Qif not get PE data")
        }
        if _, e := GetCurPB(); e != nil{
                panic("Error: Qif not get PB data")
        }
        if _, e := GetCurMtsr(); e != nil{
                panic("Error: Qif not get Mtss data")
        }
        if _, e := GetCurVolr(); e != nil{
                panic("Error: Qif not get Vol data")
        }
}


func QifLogin( )(bool){
        _, suc := goCallpy("Login_JQ", "18602122079", "calcaapi")
        if suc {
                fmt.Println("QifLogin success.")
        }
        return suc
}




func GetCurPE()(o T_pe, err error){
        o = GetPE(today)
        return o, nil
}


func GetCurPB()(o T_pb, err error){
        o = GetPB(today)
        return o, nil
}


func GetCurMtsr()(o T_mtsr, err error){
        o = GetMtsr(today)
        return o, nil
}


func GetCurVolr()(o T_volr, err error){
        o = GetVolr(today)
        return o, nil
}


