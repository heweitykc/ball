package main

import(
	"fmt"
//	"time"
//	"log"
)

var c = make(chan int, 10)
var a string

func f(){
	 a = "hello,world"
	 list0 := make([]int, 1,3)
	 println(list0[0])
	 list0 = append(list0,100)
	 println("len:",len(list0),"cap:",cap(list0))
	 list0 = append(list0,101)	 
	 list0 = append(list0,102)
	 println("len:",len(list0),"cap:",cap(list0))
	 list0 = append(list0,103)
	 println("len:",len(list0),"cap:",cap(list0))
	 println(list0[4])
	 c <- 0
}

func g(){
	list0 := new([]int)
	
	println(list0)
}

func h(){
	done := make(chan bool)
	
	values := []string{"a","b","c"}
	for _,v := range values {
		go func(h string){
			fmt.Println("cal:",h)
			done <- true
		}(v)
	}
	
	for _,k := range values {
		d := <-done
		fmt.Println(k,"=",d)
	}
}

func main(){
	h()
	//g()
	//go f()
	//<-c
	println(a)
}