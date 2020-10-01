package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	link := "https://res.cloudinary.com/infokom/image/upload/v1568261329/sidowi-users/1568261323.jpg"
	split := strings.Split(link, "/")
	filename := strings.Join(split[8:], "")
	filenameSplit := strings.Split(filename, ".")
	publicid := filenameSplit[0]
	fmt.Println(publicid)
	//fmt.Println(strings.Join(split[7:], "/"))

	//v := 5
	//p := &v
	//
	//fmt.Println(*p)
	//
	//changePointer(p)
	//fmt.Println(*p)
	//
	//println("==========================================================================")
	//timeStart := time.Now()
	//
	//<-worker()
	//<-worker()
	//<-worker()
	//println(int(time.Since(timeStart).Seconds()))
	//
	//println("==========================================================================")
	//var counter int
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		counter++
	//	}()
	//}
	//fmt.Println(counter)
}

func changePointer(p *int) {
	v := 3
	*p = v
}

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 42
	}()

	return ch
}
