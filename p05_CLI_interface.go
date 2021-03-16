package main

import (
	"flag"
	"fmt"
)

func main5()  {
	helpPtr :=flag.String("yes","","asdasdsadsa")
	helpPtr2 :=flag.String("asd","asdasdasd","a string")

	flag.Parse()

	fmt.Println(*helpPtr)
	fmt.Println(*helpPtr2)
}