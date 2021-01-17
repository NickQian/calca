/**********************************************************************************************************
/*  Quantitative interface:
/*  可选： Wind(没找到linux接口文件), 京东JDQuant(2018.12关闭), MindGo(同花顺),优矿(通联数据Uqer.io), 米筐('高质量数据免费使用'), 
/*           JoinQuant(聚宽,已申请1年)，BigQuant, tushare等, 国外Quantopian, WorldQuant
/*  使用gRPC 还是cgo？  cgo: github.com/sbinet/go-python
/* ----
/*  License: BSD
/* ----
/*  0.2?:  ?give up go-python and use gRPC? - 2019.9
/*  0.1.1: 放弃使用interface, 采用struct里面放函数指针的办法 - 2019.6  Nick cKing
/*  0.1: init version - 2019.1 -  Nick cKing
/************************************************************************************************************/

package qif


import(
        //"github.com/DataDog/go-python3"		// python3.
	"github.com/sbinet/go-python"           	// python.
	"log"
        "fmt"
        "time"
        . "define"
        "errors"
        "strings"
        "strconv"
        "runtime"
        "path"
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

var now = time.Now()
//var now time.Time = time.Now()

var Today string

var PyBrg *T_go2py
//var PyBrg *T_go2py = &T_go2py{Str2Py: python.PyString_FromString,
//                              Py2Str: python.PyString_AsString,
//                              }

var PyModule *python.PyObject


//------------------------ func实现 ------------------------------------

func  MarketUpdate(a *T_A) (suc bool){
        resDic := GetMarket(Today)
        if len(resDic) == 0{
        	fmt.Println("Error: <MarketUpdate> result is empty. Maybe internet access problem or not trade day.")
        	return false
        }
        FilDicToA(resDic, a)
        fmt.Println("<MarketUpdate> done. Result:", suc)
	return true
}


func GetKline(dayStart, dayEnd string)(kline []float64){
	
	return
}

func GetMarket(day string)(dicmkt map[string]float64){
	//dicmkt = make(map[string]float64, 100)
	dicmkt = make(map[string]float64)
	dicmkt, _, _ = goCallpy("getMarketMap", day)
        if len(dicmkt) == 0{
		fmt.Println("Error: <GetMarket> result dicmkt is empty. Maybe internet problem or not trade day. ")
        	//return false
        }
        //fmt.Println("<GetMarket> result: day, dicmkt:", day, dicmkt)

	return 	dicmkt
}


func FilDicToA(dicmkt map[string]float64, a *T_A)(bool){
        if len(dicmkt) >0 {
	        a.Cmc.Cmc_sh, a.Cmc.Cmc_sz, a.Cmc.Cmc_gem = dicmkt["cmc_sh"], dicmkt["cmc_sz"], dicmkt["cmc_gem"]
        	a.Cmc.Cmc_total = a.Cmc.Cmc_sh + a.Cmc.Cmc_sz

        	a.Pe.Pe_sh,   a.Pe.Pe_sz,   a.Pe.Pe_gem   = dicmkt["pe_sh"],  dicmkt["pe_sz"],  dicmkt["pe_szm"]
		a.Pe.Pe_total = a.Pe.Pe_sh * (a.Cmc.Cmc_sh/a.Cmc.Cmc_total) + a.Pe.Pe_sz * (a.Cmc.Cmc_sz/a.Cmc.Cmc_total)

	        a.Tnr.Tnr_sh, a.Tnr.Tnr_sz = dicmkt["tnr_sh"], dicmkt["pe_sz"]

        	vol_sh,  vol_sz,  vol_gem := dicmkt["vol_sh"], dicmkt["vol_sz"], dicmkt["vol_gem"]
		a.Volr.Volr_total = 100*(vol_sh + vol_sz)/a.Cmc.Cmc_total
		a.Volr.Volr_gem, a.Volr.Volr_sh, a.Volr.Volr_sz = 100*vol_gem/a.Cmc.Cmc_gem, 100*vol_sh/a.Cmc.Cmc_sh, 100*vol_sz/a.Cmc.Cmc_sz

		mtss_sh, mtss_sz := dicmkt["mtss_sh"], dicmkt["mtss_sz"]
        	a.Mtsr.Mtsr_total = 100*(mtss_sh + mtss_sz)/a.Cmc.Cmc_total
        	a.Mtsr.Mtsr_sh, a.Mtsr.Mtsr_sz = 100*mtss_sh/a.Cmc.Cmc_sh, 100*mtss_sz/a.Cmc.Cmc_sz
		return true
	}
	return false
}


func GetTradeDays(date string)(days []string){
	_, days, _ = goCallpy("getTradeDays", date, strconv.Itoa(PRE_SMP_NUM) )    // string to facilitate <goCallpy>
	return
}

func GetPE(day string) (pe T_pe) {
        peMap, _, _ := goCallpy("getPE", day)
        pe.Pe_sh = peMap["pe_sh"]
        return pe
}


func GetPB(day string) (pb T_pb) {
        pbMap, _, _:= goCallpy("getPB", day)
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
	 //Num2Py func(int)(*python.PyObject)          // func PyInt_FromLong(val int) *PyObject
	 Num2Py func(float64) *python.PyObject
	 Py2Num func(*python.PyObject)(float64)      // func PyFloat_AsDouble(self *PyObject) float64
}

//PyStr := python.PyString_FromString
//GoStr := python.PyString_AS_STRING


// init python & import module & login
func go2pyInit()(suc bool){     //pymodule *python.PyObject){        //, PyBrg *T_go2py){
        suc = false

        //python.Py_Initialize()
	//if !python.Py_IsInitialized() {

	_, filename, _, _ :=runtime.Caller(1)
	pyFilePath := path.Join(path.Dir(filename), "./")
	fmt.Println("Info;pyFilePath:", pyFilePath)

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
			 Num2Py: python.PyFloat_FromDouble,        // func PyFloat_FromDouble(v float64) *PyObject
			 Py2Num: python.PyFloat_AsDouble,      // func PyFloat_AsDouble(self *PyObject) float64
                        }

        //--- select Qif的公司，选一家，注释掉其它家
        switch QIF_VENDOR {
                case "JQ":      PyModule = ImportModule(pyFilePath, "if_jq")
                case "UQ":      PyModule = ImportModule(pyFilePath, "if_uq")
                case "RQ":      PyModule = ImportModule(pyFilePath, "if_rq")
                case "BQ":      PyModule = ImportModule(pyFilePath, "if_bq")
                case "tushare": PyModule = ImportModule(pyFilePath, "if_ts")
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


func goCallpy(defname string, args ...string)(omap map[string]float64, oslice[]string, suc bool){
        suc = false

        f := PyModule.GetAttrString(defname)
        argv := python.PyTuple_New(len(args))                        // func PyTuple_New(sz int) *PyObject

        for i, value := range args{
                python.PyTuple_SetItem(argv, i, PyBrg.Str2Py(value)) //func PyTuple_SetItem(self *PyObject, pos int, o *PyObject) error
        }
        //fmt.Printf("--(1)--f:%v,defname:%v, argv:%v--\n", f, defname, argv)

	switch defname{
	case "Login_JQ", "Login_RQ", "Login_BQ", "Login_UQ", "Login_TS":
		sucObj := f.Call(argv, python.Py_None)               // func (self *PyObject) Call(args, kw *PyObject) *PyObject
		fmt.Println("--(2-1)--Qif_Login: sucObj is:", sucObj)
		return nil, nil, sucObj.IsTrue()
	case "getTradeDays":
		resListObj := f.Call(argv, python.Py_None)
		if resListObj != nil{
			oslice = ListResExtract(resListObj)
			suc = true
		}
	default:   						     // default is get market dict
        	resDictObj := f.Call(argv, python.Py_None)              // func (self *PyObject) Call(args, kw *PyObject) *PyObject
		if resDictObj != nil{
        		//fmt.Printf("--(2-3)-- <Call out>: %v \n", resDict)
			omap = DicResExtract2(resDictObj)
		}
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
	}
	omap = make(map[string]float64, python.PyList_Size(keyObjs))

        for _, key := range keys{
                itemValue := python.PyDict_GetItemString(dicIn, key)
       		omap[key] = PyBrg.Py2Num(itemValue)        //PyInt_FromLong
        }
        //fmt.Println("------ (3) Dic Extract----:", omap)
	return
}


func DicResExtract2(dicIn *python.PyObject)(omap map[string]float64){
	dicItems := python.PyDict_Items(dicIn)  //func PyDict_Items(self *PyObject) *PyObject{Return a PyListObject }
        omap = make(map[string]float64, python.PyList_Size(dicItems))

        for i:=0; i<python.PyList_Size(dicItems); i++{
        	dicItem  := python.PyList_GetItem(dicItems, i)    // return a PyListObject
		keyObj   := python.PyTuple_GetItem(dicItem, 0)
		valueObj := python.PyTuple_GetItem(dicItem, 1)
		key      := PyBrg.Py2Str(keyObj)
		omap[key] = PyBrg.Py2Num(valueObj)        //PyInt_FromLong
        }
        //fmt.Println("------ (3) Dic Extract----:", omap)
        return
}


// only process string
func ListResExtract(listIn *python.PyObject)(oslice []string){
	for i:=0; i<python.PyList_Size(listIn); i++{          //func PyList_Size(self *PyObject) int
                dayObj := python.PyList_GetItem(listIn, i)    // PyList_GetItem(self *PyObject, index int) *PyObject //PyList_GetI$
                oslice = append(oslice, PyBrg.Py2Str(dayObj))
        }
	return
}

// ----------------- operate with python api ---------------------

func init(){
        todayFull := now.Format(TIME_LAYOUT_STR)
        todaySlice := strings.SplitAfter(todayFull, " ")
        Today = strings.TrimSpace(todaySlice[0])
        fmt.Printf("<qif:init>: Today:%s---%s\n", Today, todaySlice)

        go2pyInit()
        QifLogin()
	time.Sleep(1000 * time.Millisecond)  // sleep to wait remote server finish authorization
}


func QifLogin( )(suc bool){
	_, _, loginSuc := goCallpy("Login_JQ", "18602122079", "calcaapi")
        fmt.Println("QifLogin success result:", loginSuc)
        if loginSuc {
                fmt.Println("Info: QifLogin success.")
        }else{
        	fmt.Println("Error: QifLogin failed.")
        }
      return suc
}




func GetCurPE()(o T_pe, err error){
        o = GetPE(Today)
        return o, nil
}


func GetCurPB()(o T_pb, err error){
        o = GetPB(Today)
        return o, nil
}


func GetCurMtsr()(o T_mtsr, err error){
        o = GetMtsr(Today)
        return o, nil
}


func GetCurVolr()(o T_volr, err error){
        o = GetVolr(Today)
        return o, nil
}


