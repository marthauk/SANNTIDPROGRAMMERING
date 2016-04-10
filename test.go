package main

import (
	"fmt"
)

type Button struct{
	Button_type int
	Floor int
}

func main(){
	Button_Press_Chan := make(chan Button,10)
	x := Button{Button_type: 2, Floor: 2}

	Button_Press_Chan <- x
	
	fmt.Println("X is:",x)

	fmt.Println("The channel is filled with")
	for i := range Button_Press_Chan {
		fmt.Println(i)
	}
}
