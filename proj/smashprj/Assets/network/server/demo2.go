package main

import(
//	"fmt"
//	"time"
	"sync"
)

var once sync.Once
var l sync.Mutex
var a string
var b int

func f(){
	 once.Do(testdo)
	 a = "hello,world"
	 l.Unlock()
}

func testdo(){
	b++
	print("testdo\n")
}

func main(){
	l.Lock()
	go f()
	l.Lock()
	go f()	
	print(a+"\n")
	print(b)
}