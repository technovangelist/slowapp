package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/didip/tollbooth"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var limitpersecond string
var responsetime string

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "limitpersecond, l",
			Usage: "Specify how many to limit per second",
			Value: "1000",
		},
		cli.StringFlag{
			Name:  "r, responsetime",
			Usage: "Specify response time",
			Value: "1000",
		},
	}
	app.Action = func(c *cli.Context) {
		responsetime = c.String("responsetime")
		limitpersecond = c.String("limitpersecond")
		runwebsite()
	}
	app.Run(os.Args)

}

func runwebsite() {
	ilimit, _ := strconv.ParseInt(limitpersecond, 10, 64)
	fmt.Println("response time is " + responsetime)
	fmt.Println("limitpersecond is " + limitpersecond)
	limiter := tollbooth.NewLimiter(ilimit, time.Second)
	http.Handle("/", tollbooth.LimitFuncHandler(limiter, slowHandler))
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func slowHandler(w http.ResponseWriter, req *http.Request) {
	iresponse, _ := strconv.ParseInt(responsetime, 10, 64)
	time.Sleep(time.Duration(iresponse) * time.Millisecond)
	w.Write([]byte("Hello, World!"))
}
