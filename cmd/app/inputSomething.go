package main

import (
	"fmt"
)

type Inputer interface {
	GetInputSomething() string
	enterName() string
	enterNumber() string
}

type InputSomething struct {
	name string
	num  string
}

func (i *InputSomething) SetInputSomething(name string, num string) {
	i.name = name
	i.num = num
}

func (i *InputSomething) GetInputSomething() (string, string) {
	return i.name, i.num
}

func (i *InputSomething) enterName() string {
	fmt.Println("\nplese input name")
	fmt.Scan(&i.name)
	return i.name
}

func (i *InputSomething) enterNumber() string {
	for {
		fmt.Println("\nplese input a digit")
		fmt.Scan(&i.num)
		switch i.num {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			return i.name
		default:
			fmt.Println("error. please input a digit again.")

		}
	}
}
