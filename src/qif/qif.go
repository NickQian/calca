/**********************************************************************************************************
/*  Quantitative interface:
/*  可选： Wind(没找到linux接口文件), 京东JDQuant(2018.12关闭), MindGo(同花顺),优矿(通联数据Uqer.io), 米筐('高质量数据免费使用'), 
/*           JoinQuant(聚宽,已申请1年)，BigQuant, tushare等, 国外Quantopian, WorldQuant
/*  使用gRPC 还是cgo？  cgo: github.com/sbinet/go-python
/* ----
/*  License: BSD
/* ----
/*  0.2?: ?give up go-python and use gRPC? - 2019.9
/*  0.1.1: 放弃使用interface, 采用struct里面放函数指针的办法 - 2019.6  Nick cKing
/*  0.1: init version - 2019.1 -  Nick cKing
/************************************************************************************************************/

package qif


import(
        //"github.com/DataDog/go-python3"		// python3.
	"github.com/sbinet/go-python"           // python.
	"log"
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

var PyBrg *T_go2py
//var PyBrg *T_go2py = &T_go2py{Str2Py: python.PyString_FromString,
//                              Py2Str: python.PyString_AsString,
//                              }

var PyModule *python.PyObject

//------------------------ func实现 ------------------------------------

func MarketUpdate(a *T_A) (suc bool){
        AstatDic, _  := goCallpy("getMarketMap", today)
        //a.Cmv.Cmv_total = AstatDic["cmv_total"]
        fmt.Println("MarketUpdate done. Result:", AstatDic)
	return true
}


// have a(a stock) look
func HavaLook(day string, a *T_A) (bool){
        peDic  := GetPE(day)
        dicmkt := GetMarket(day)
        if dicmkt == nil{
        	return false
        }
        fmt.Println("&$^#$^##$^--peDic, dicmkt:", peDic, dicmkt)
        /*pbDic   := GetPB(day)
        volrDic := GetVolr(day)
        tnrDic  := GetTnr(day)
        mtsrDic := GetMtsr(day)
        a.Pe.Pe_sh = peDic.Pe_sh
        a.Pb.Pb_sh = pbDic.Pb_sh
        a.Volr.Volr = volrDic.Volr
        a.Tnr.Tnr = tnrDic.Tnr
        a.Mtsr.Mtsr_total = mtsrDic.Mtsr_total  */
        return true
}


func GetMarket(day string)(omap map[string]float64){
	omap = make(map[string]float64, 77)
	omap, _ = goCallpy("getMarketMap", day)
	return
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
         Str2Py func(string)(*python.PyObject)      // func PyString_FromString(v string) *PyObject
         Py2Str func(*python.PyObject)(string)
	 Num2Py func(int)(*python.PyObject)          // func PyInt_FromLong(val int) *PyObject
	 Py2Num func(*python.PyObject)(float64)      // func PyFloat_AsDouble(self *PyObject) float64
}

//PyStr := python.PyString_FromString
//GoStr := python.PyString_AS_STRING


// init python & import module & login
func go2pyInit()(suc bool){     //pymodule *python.PyObject){        //, PyBrg *T_go2py){
        suc = false

        //python.Py_Initialize()
	//if !python.Py_IsInitialized() {

	err := python.Initialize()
        if err != nil {
                //panic(err.Error())
		panic("Err: <go2pyInit> initial failed. ")
                log.Panic(err)
        }else{
		fmt.Println("info: python Py_Initialize done! ")
        }

        PyBrg = &T_go2py{Str2Py: python.PyString_FromString,
                         Py2Str: python.PyString_AsString,
			 Num2Py: python.PyInt_FromLong,        // func PyInt_FromLong(val int) *PyObject
			 Py2Num: python.PyFloat_AsDouble,      // func PyFloat_AsDouble(self *PyObject) float64
                        }

        //--- select Qif的公司，选一家，注释掉其它家
        switch QIF_VENDOR {
                case "JQ":      PyModule = ImportModule("./", "if_jq")
                case "UQ":      PyModule = ImportModule("./", "if_uq")
                case "RQ":      PyModule = ImportModule("./", "if_rq")
                case "BQ":      PyModule = ImportModule("./", "if_bq")
                case "tushare": PyModule = ImportModule("./", "if_ts")
                default:        //Log.Fatal("Err: [qif]: no Qif configurated.")
                        	panic("Err: Wrong (QIF_VENDOR) value.")
        }
        if PyModule == nil{
                panic(" #: Error: import result nil.")
        }
        fmt.Printf(" PyModule: %v, [MODULE]repr()= %s   \n", PyModule,  PyBrg.Py2Str(PyModule.Repr() ) )
        return
}


func ImportModule(dir, name string)(*python.PyObject){
        //(1) add dir into python env "sys.path"
        sysModule := python.PyImport_ImportModule("sys")                 // func PyImport_ImportModule(name string) *PyObject
        path := sysModule.GetAttrString("path")                          // path is ['', '', '']
        python.PyList_Insert(path, 0, PyBrg.Str2Py(dir) )                // func PyList_Insert(self *PyObject, index int, item_to_insert *PyObject)
        path2 := python.PySys_GetObject("path")                          // func PySys_GetObject(name string) *PyObject
        fmt.Println("sys.path is(after): ", path2)
        //(2) import *.py module in the dir which contain this *.py file
        return python.PyImport_ImportModule(name)
}


func goCallpy(defname string, args ...string)(omap map[string]float64, suc bool){
        suc = false

        f := PyModule.GetAttrString(defname)
        argv := python.PyTuple_New(len(args))           // func PyTuple_New(sz int) *PyObject
        for i, value := range args{
                python.PyTuple_SetItem(argv, i, PyBrg.Str2Py(value)) //func PyTuple_SetItem(self *PyObject, pos int, o *PyObject) error
        }
        fmt.Printf("------ (1)-----f:%v, defname: %v, argv: %v----->   \n", f, defname, argv)

	resDict := f.Call(argv, python.Py_None)     // func (self *PyObject) Call(args, kw *PyObject) *PyObject
        fmt.Printf("------ (2)----> <Call out>----- %v \n", resDict)

	if defname != "Login_JQ"  &&  resDict != nil{
		omap = DicResExtract2(resDict)
		suc = true
	}
        return
}


func DicResExtract(dicIn *python.PyObject)(omap map[string]float64){
        keyObjs := python.PyDict_Keys(dicIn)           // return a PyListObject
	var keys []string

	for i:=0; i<python.PyList_Size(keyObjs); i++{       //func PyList_Size(self *PyObject) int
		keyObj := python.PyList_GetItem(keyObjs, i) //func PyList_GetItem(self *PyObject, index int) *PyObject
                key := PyBrg.Py2Str(keyObj)
	        keys = append(keys, key)
	        fmt.Println("------Dic Extract (3)--keyObj,key,keys:--->", keyObj, key, keys)
	}
	omap = make(map[string]float64, python.PyList_Size(keyObjs))

        for _, key := range keys{
                itemValue := python.PyDict_GetItemString(dicIn, key)
       		omap[key] = python.PyFloat_AsDouble(itemValue)        //PyInt_FromLong
        }
	return
}


func DicResExtract2(dicIn *python.PyObject)(omap map[string]float64){
	dicItems := python.PyDict_Items(dicIn)  //func PyDict_Items(self *PyObject) *PyObject{Return a PyListObject }
        omap = make(map[string]float64, python.PyList_Size(dicItems))

        for i:=0; i<python.PyList_Size(dicItems); i++{
        	dicItem := python.PyList_GetItem(dicItems, i)    // return a PyListObject
		keyObj   := python.PyTuple_GetItem(dicItem, 0)
		valueObj := python.PyTuple_GetItem(dicItem, 1)
		key   := PyBrg.Py2Str(keyObj)
		omap[key] = python.PyFloat_AsDouble(valueObj)        //PyInt_FromLong
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
        QifLogin()
        fmt.Println("<go2pyInit> Init & import module done & log in successfully!  ")

/*        // fetch today data once to update A status
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
*/

}


func QifLogin( )(suc bool){
	_, CallRes := goCallpy("Login_JQ", "18602122079", "calcaapi")
        fmt.Println("QifLogin success result:", CallRes)
/*
        if suc {
                fmt.Println("Info: QifLogin success.")
        }
*/
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


