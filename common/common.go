package common

func Ternary(str string) (ret string) {
	if len(str) != 0 {
		ret = str
	} else {
		ret = "N/A"
	}

	return ret
}
