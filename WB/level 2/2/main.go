package main

import "strings"

func strk(str string) string{
	var sliseOpen []string
	for i, s := range str{
		if (rune(s) < 57 && rune(s) > 47) && i != 0{
			for k:=0;k<int(s)-49;k++{
				sliseOpen = append(sliseOpen, string(str[i-1]))
			}
		}else if (rune(s) < 57 && rune(s) > 47) && i == 0{
			return "некорректная строка"
		}else{
			sliseOpen = append(sliseOpen, string(str[i]))
		}
	}
	stringOpen := strings.Join(sliseOpen, "")
	return stringOpen
}