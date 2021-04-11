/*********************************************************
/* web top file
/* ----
/*  License: BSD
/* ----
/* v0.1 init version  --- 2019.1.15
/*********************************************************/


package main

import(
        "net/http"
	"html/template"
	//"io/ioutil"
	//"encoding/json"
        "fmt"
        //"strings"
	"log"
	. "define"
	"time"
        //"simtrade"
	"cmn"
	//"github.com/robfig/cron"
)


const (
	contrb_site = "github.com/NickQian/calca"
	footer1 = "about"
	footer2 = "donate"
)


var calRes  = T_CalRes{ Slogan : "  * Calca For Freedom * ",
			Bi  : 50,
			Ti  : 0,
	     	       }


//------------------- cron read calca --------------------------
/*func Cron_ReadCalRes()(){
	calRes.CalDate = cmn.TodayStr

	c := cron.New( )
	spec := "00 30 22 * * 1-5"        // sec minute h day_of_month month day_of_week
	c.AddFunc(spec, ReadCalRes()error  )
	fmt.Println("-->>>>> [Cron_ReadCalRes] added. ")
	c.Start()
}
*/

func ReadCalRes() (err error) {
	// read mipos res

	// read scan res

	// read trade res
	calRes.TrdRes, err = cmn.ReadResTrd(FN_RES_TRD)

	return
}


//-------------------------------- main ------------------------------------------
// use "DefaultServeMux"
func main(){
	//--- static ---
	http.Handle("/css/",     http.FileServer(http.Dir("static")) )     // note the folder hierachy
	http.Handle("/images/",  http.FileServer(http.Dir("static")) )

	//--- dynamic ---
	http.HandleFunc("/index/",  indexHandler    )
	http.HandleFunc("/",        NotFoundHandler )
	http.HandleFunc("/about/",  aboutHandler )


	//--- srv create & listen ---
	log.Println("Webca is Listening..." )
	err := http.ListenAndServe(":80", nil)  //"nil" means use DefaultServeMux ("127.0.0.0:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	//--- calca cron ---
	//calca_cron()
	//select {}

}




//-------------------- home ------------------------

func indexHandler(w http.ResponseWriter, r *http.Request)(){
	fmt.Printf("-------> Info: in <indexHandler>, request.Header: %v   \n", r.Header)
	fmt.Println("URL:", r.URL)
	fmt.Println(" ### Will ParseFiles header.html... ")

	// parser files
	t, err := template.ParseFiles("template/header.html", "template/index.html", "template/footer.html" )
	checkError(err)

	// replace var
	ReadCalRes()                // update "calres" through read res files
	err = t.Execute(w, calRes)
	checkError(err)
}



//------------------- about -------------------
func aboutHandler(w http.ResponseWriter, r *http.Request)(){
        fmt.Printf("-------> Info: in <aboutHandler>,  r.URL: %v  \n" , r.URL)

        // parser files
        t, err := template.ParseFiles("template/about.html", "template/footer.html" )
        checkError(err)

        // replace var
        err = t.Execute(w, calRes)
        checkError(err)
}


//-------------------- 404 --------------------------

func NotFoundHandler(w http.ResponseWriter, r *http.Request)(){
	if r.URL.Path == "/"{
		log.Printf("Info: <notFoundHandler> URL.path ==/ . status: %v, Will redirect...  \n", http.StatusFound )
		http.Redirect( w, r, "/index", http.StatusFound)
	}

	// parse template
	t, err := template.ParseFiles("template/404.html")
	if (err != nil){
		fmt.Printf("ERROR: ParseFile->404.html error: %v ...      \n", err)
        	log.Println(err)
	}

	// write & t.execute
	t.Execute(w, nil)
	log.Println("Info: <notFoundHandler> go till end. Finished. " )
}





//-------------------- https --------------------------------------------
func SrvCreate () (){
	srv := &http.Server{
		ReadTimeout :     5 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout :   120 * time.Second,
		//TLSConfig   :   tlsConfig,
		//Handler     :   serveMux, 
		}
	log.Println(srv.ListenAndServeTLS("", "" ) )
}



//------------------------------------------------------------
func checkError(err error){
    if err != nil{
        fmt.Println("Fatal error", err.Error() )
        // panic(err)
    }
}







