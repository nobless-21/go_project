package main

import (
	"fmt"
	"os"
	"github.com/beevik/ntp"
)

func main(){
	server := "pool.ntp.oag"

	time, err := ntp.Time(server)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Ошибка получения времени: %v", err)
		os.Exit(1)
	}

	fmt.Println(time)
}