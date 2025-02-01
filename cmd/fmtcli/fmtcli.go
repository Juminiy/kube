package main

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"time"
	"unsafe"
)

func main() {
	scanv()
}

func printv() {
	fmt.Println("print normal")
	fmt.Print(true, false, uint(1), -1, " I am print ", time.Now(), "\n")
	fmt.Println("print bool")
	fmt.Printf("%T:%t\t%T:%t\n", true, true, false, false)
	fmt.Println("print int")
	fmt.Printf("%b %c %d %o %O %q %x %X %U\n", ival0, 'c', ival0, ival0, ival0, 'c', ival0, ival0, ival0)
	fmt.Println("print float")
	fmt.Printf("%b %e %E %f %F %g %G %x %X\n", fval0, fval0, fval0, fval0, fval0, fval0, fval0, fval0, fval0)
	fmt.Println("print string")
	fmt.Printf("%s %q %x %X\n", sval0, sval0, sval0, sval0)
	fmt.Println("print []byte")
	fmt.Printf("%v\n", []byte(sval0+sval1))
	fmt.Println("print slice")
	fmt.Printf("%v\n", []int{ival0, ival1})
	fmt.Println("print struct")
	fmt.Printf("%v\t%+v\t%#v\t%T\t%%\n", stval, stval, stval, stval)
	fmt.Println("print map")
	fmt.Printf("%v\n", map[string]int{"k1": 1, "k2": 2})
	fmt.Println("print array")
	fmt.Printf("%v\n", [2]int{ival0, ival1})
	fmt.Println("print chan")
	fmt.Printf("%v\n", make(chan int))
	fmt.Println("print func")
	fmt.Printf("%p\n", fval)
	fmt.Println("print pointer")
	fmt.Printf("%v\n", util.New(ival0))
	fmt.Println("print unsafePointer")
	fmt.Printf("%v\n", unsafe.Pointer(util.New(ival0)))

	/*fmt.Fprint()
	fmt.Fprintf()
	fmt.Fprintln()*/

	/*fmt.Sprint()
	fmt.Sprintf()
	fmt.Sprintln()*/

	/*fmt.Append()
	fmt.Appendf()
	fmt.Appendln()*/
}

func scanv() {
	var (
		bv  bool
		iv  int
		uiv uint
		fv  float64
		sv  string
		bsv []byte

		sfmt = "%t %d %d %f %s %s"
		err  error
	)
	/*
		fmt.Fscan() 	// \n isSpace, \n notEnd
		fmt.Fscanln() 	// \n notSpace, \n isEnd
		fmt.Fscanf() 	// \n notSpace, \n notEnd
	*/

	_, err = fmt.Scan(&bv, &iv, &uiv, &fv, &sv, &bsv)
	util.Must(err)
	fmt.Println(bv, iv, uiv, fv, sv, bsv)

	_, err = fmt.Scanln(&bv, &iv, &uiv, &fv, &sv, &bsv)
	util.Must(err)
	fmt.Println(bv, iv, uiv, fv, sv, bsv)

	_, err = fmt.Scanf(sfmt, &bv, &iv, &uiv, &fv, &sv, &bsv)
	util.Must(err)
	fmt.Println(bv, iv, uiv, fv, sv, bsv)

	_, err = fmt.Sscan("true -1919180 114514 -114.514 vvmike vvarcw", &bv, &iv, &uiv, &fv, &sv, &bsv)
	util.Must(err)
	fmt.Println(bv, iv, uiv, fv, sv, bsv)

	_, err = fmt.Sscanln("false -114514 1919810 114.514 rrtext rrbusy", &bv, &iv, &uiv, &fv, &sv, &bsv)
	util.Must(err)
	fmt.Println(bv, iv, uiv, fv, sv, bsv)

	_, err = fmt.Sscanf("true -95270 8888 1919.810 yyhex yysoft", sfmt, &bv, &iv, &uiv, &fv, &sv, &bsv)
	util.Must(err)
	fmt.Println(bv, iv, uiv, fv, sv, bsv)

}

func otherv() {
	//fmt.FormatString()
	//fmt.Errorf()
}

type vstruct struct {
	Name string
	Desc string
	LLim int64
	RLim int64
}

const ival0 = -114514
const ival1 = 1919810
const fval0 = 114.514
const fval1 = 1919.180
const sval0 = "te amo lyy"
const sval1 = "я люблю тебя, Лий"

var stval = vstruct{
	Name: sval0,
	Desc: sval1,
	LLim: ival0,
	RLim: ival1,
}

var fval = func() {}
