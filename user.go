package main

import "fmt"

var exp = 0
var level = 1

func uplevel(a int){
	exp += a
	level = (exp / 10)
}

func main() {
	uplevel(100)
	fmt.Printf("Level:%d, Exp:%d",level,exp)
}