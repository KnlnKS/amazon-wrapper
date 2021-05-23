package common

import "log"

func Ternary(str string) (ret string) {
	if len(str) != 0 {
		ret = str
	} else {
		ret = "N/A"
	}

	return ret
}

func OnError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
