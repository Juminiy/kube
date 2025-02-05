package main

import (
	"log"
	"os"
)

func main() {
	//arr := []int{1, 2}
	arr := make([]int, 0, 8)
	arr = append(arr, 1, 2)
	doSlice(arr)
	lg(arr)
}

func doSlice(arr []int) {
	arr = append(arr, 3, 4, 5, 6, 7, 8)
	lg(arr)
	arr[0] = 200
}

func myCall() {
	arr := []int{1, 2, 3}
	lg(arr) // 1,2,3
	callMain(arr)
	lg(arr) // 1,2,3
}

func callMain(arr []int) {
	lg(arr) // 1,2,3
	callOther(arr)
	lg(arr) // 1,2,3
}

func callOther(arr []int) {
	lg(arr) // 1,2,3
	arr = append(arr, 4, 5, 6)
	arr[0] = 666
	lg(arr) // 666,2,3,4,5,6
}

var l = log.New(os.Stdout, "TEST ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

func lg(v ...any) {
	l.Println(v...)
}
