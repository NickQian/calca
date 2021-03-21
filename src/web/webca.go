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
)


const (
	contrb_site = "contribute:  www.github.com"
	footer1 = "about"
	footer2 = "donate"
)


var calRes  = T_CalRes{ Title : "  *** Calca For Free *** ",
			Bi  : 0.0,
			Ti : 0.5,
	     	       }



// use "DefaultServeMux"
func main(){
	//--- static ---
	http.Handle("/about/", http.FileServer(http.Dir("static")) )

	//--- dynamic ---
	http.HandleFunc("/index/",  indexHandler    )
	http.HandleFunc("/",        notFoundHandler )


	//--- srv create & listen ---
	log.Println("Webca is Listening..." )
	err := http.ListenAndServe(":9090", nil)  //"nil" means use DefaultServeMux ("127.0.0.0:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}




//-------------------- home ------------------------

// home page
func indexHandler(w http.ResponseWriter, r *http.Request)(){
	fmt.Printf("-------> Info: in <indexHandler>, request.Header: %v   \n", r.Header)
	fmt.Println("URL:", r.URL)
	fmt.Println(" ### Will ParseFiles header.html... ")

	// parser files
	t, err := template.ParseFiles("template/header.html", "template/index.html", "template/footer.html",)
	checkError(err)

	// replace var
	err = t.Execute(w, calRes)
	checkError(err)
}



//-------------------- 404 --------------------------

func notFoundHandler(w http.ResponseWriter, r *http.Request)(){
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










// 
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



func checkError(err error){
    if err != nil{
        fmt.Println("Fatal error", err.Error() )
        // panic(err)
    }
}








/*
func main(){

        //--- mux ---
        mux := http.NewServeMux()

        //- static -
        hdl_css := http.Handle("/css/", http.FileServer(http.Dir("static")))
        hdl_js  := http.Handle("/js/",  http.FileServer(http.Dir("static")))

        //- dynamic-
        //hdl_home := http.HandleFunc("/index/", indexHandler)  
        hdl_home := http.HandleFunc("/index/", indexHandler)    
        hdl_info := http.HandleFunc("/about/", aboutHandler)
        hdl_404  := http.HandleFunc("/",       notFoundHandler)
        mux.Handle("/", hdl_home)
        mux.Handle("/about/", aboutHandler)

        //--- srv create & listen ---
        log.Println("Listening..."
        err := http.ListenAndServe(":9090", mux)  //("127.0.0.0:8000", nil)
        if err != nil {
                log.Fatal("ListenAndServe:", err)
        }
}*/

