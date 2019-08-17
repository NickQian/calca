
package xxmain


import (
        "fmt"
        "time"
        "log"
        "errors"
       )


type CalErr struct{
        QifErr string
        CptErr string
        IoErr  string
        time time.Time
}

func main(){

        fmt.Printf("hello, go world.  \n")
        fmt.Println("my package name is main.")

        _, err := ftest()
        fmt.Println(err)


        fmt.Print("fmt.Printf show you the time:", time.Now(), "\n" )
        log.Print("log show you the time:", time.Now() )


}

func ftest()(int, error){
        return -1, errors.New(" ## this is msg error test")

}
