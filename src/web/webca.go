# v0.01 web top --- 2019.1.15  


package main

import(
       "simtrade"
       "fmt"
       "net/http"
	   "io/ioutil"
	   "encoding/json"
       "strings"
       "log"
	   "defines"
)
 
 
constant (
    contrb_site = "contribute:  www.github.com"
    footer1 = "about"
	footer2 = "donate"
)
    

calRes := CalRes{bi:,  ti:50%, }



func indexHandler(w http.ResponseWriter, r *http.Request){
    //index, err := ioutil.ReadFile("/static.html")htmlinclude
    
    t, err := template.ParseFiles("header.html","index.html", "footer.html",)
    checkError(err)
    
    err = t.Execute(w, calRes)
	checkError(err)
}


func main(){
    // static
	http.Handle("/css/", http.FileServer(http.Dir("static")))
    http.Handle("/js/",  http.FileServer(http.Dir("static")))
    
    // dynamic
    http.HandleFunc("/index/", indexHandler)	
    http.HandleFunc("/about/", aboutHandler)
    http.HandleFunc("/",       notFoundHandler)
	
    // listen
    err := http.ListenAndServe(":9090", nil)  //("127.0.0.0:8000", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
		
}

func checkError(err error){
    if err != nil{
        fmt.Println("Fatal error", err.Error() )
        // panic(err)
    }
}


	

//-------------------- 404 --------------------------
func notFoundHandler(w, http.ResponseWriter, r *http.Request){
    if r.URL.Path == "/"{
        http.Redirect(w,r,"/index", http.StatusFound)
    }

    t, err := template.ParseFiles("template/404.html")
    if (err != nil){
        log.Println(err)
    }
    t.Execute(w, nil)
}


