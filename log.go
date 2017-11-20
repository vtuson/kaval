package main

import (
	"fmt"
	"os"
)

var report_fmt = `----------------------------------------
		test pass: %d
		test fail: %d

		Overall: %s
----------------------------------------
`

//TODO use logger
func KInfo(text ...interface{}) {
	if verbose {
		fmt.Print("INF:")
		fmt.Println(text...)
	}
}

func KError(text ...interface{}) {
	fmt.Print("ERR:")
	fmt.Println(text...)
	os.Exit(2)
}

func KWarning(text ...interface{}) {
	fmt.Print("WRN:")
	fmt.Println(text...)
}

func KPass(text ...interface{}) {
	fmt.Print("## PASS:")
	fmt.Println(text...)
	passcount++
}
func KFail(text ...interface{}) {
	fmt.Print("## FAIL:")
	fmt.Println(text...)
	kpass_val = false
	passfail++
}
func Report() {
	overall := "PASS"
	if !kpass_val {
		overall = "FAIL"
	}
	fmt.Printf(report_fmt, passcount, passfail, overall)
	if !kpass_val {
		os.Exit(2)
	}
}
