package errors

import "os"

func PrintIfErr(err error) {
	if err != nil {
		println("Error: ", err.Error())
	}
}

func ExitIfErr(err error) {
	if err != nil {
		println("Error: ", err.Error())
		os.Exit(1)
	}
}
