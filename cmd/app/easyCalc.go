package main

import (
	"fmt"
)

type EasyCalculator interface {
	getEasyCalc() *EasyCalc
	do_addCalc() (int, int, int)
	do_subCalc() (int, int, int)
	do_mulCalc() (int, int, int)
	do_divCalc() (int, int, int)
	choiceCalc()
}

type EasyCalc struct {
	x, y, ansAdd, ansSub, ansMul, ansDiv int
	chooseNum                            string
}

func (e *EasyCalc) SetEasyCalc(x, y, ansAdd, ansSub, ansMul, ansDiv int, iNum string) {
	e.x = x
	e.y = y
	e.ansAdd = ansAdd
	e.ansSub = ansSub
	e.ansMul = ansMul
	e.ansDiv = ansDiv
	e.chooseNum = iNum
}

func (e *EasyCalc) getEasyCalc() *EasyCalc {
	return e
}

func (e *EasyCalc) do_addCalc() (int, int, int) {
	e.ansAdd = e.x + e.y
	return e.x, e.y, e.ansAdd
}

func (e *EasyCalc) do_subCalc() (int, int, int) {
	e.ansSub = e.x - e.y
	return e.x, e.y, e.ansSub
}

func (e *EasyCalc) do_mulCalc() (int, int, int) {
	e.ansMul = e.x * e.y
	return e.x, e.y, e.ansMul
}

func (e *EasyCalc) do_divCalc() (int, int, int) {
	e.ansDiv = e.x / e.y
	return e.x, e.y, e.ansDiv
}

func (e *EasyCalc) choiceCalc() {
	for {
		fmt.Println("\nplease input a digit. 1: x + y, 2: x - y, 3: x * y, 4: x / y. 5: exit ")
		fmt.Scan(&e.chooseNum)
		switch e.chooseNum {
		case "1":
			fmt.Println("you choose 1 and show  x + y ")
			e.do_addCalc()
			fmt.Printf("\n%d + %d = %d", e.x, e.y, e.ansAdd)
		case "2":
			fmt.Println("you choose 2 and show  x - y ")
			e.do_subCalc()
			fmt.Printf("\n%d - %d = %d", e.x, e.y, e.ansSub)
		case "3":
			fmt.Println("you choose 3 and show  x * y ")
			e.do_mulCalc()
			fmt.Printf("\n%d * %d = %d", e.x, e.y, e.ansMul)
		case "4":
			fmt.Println("you choose 4 and show  x / y ")
			e.do_divCalc()
			fmt.Printf("\n%d / %d = %d", e.x, e.y, e.ansDiv)
		case "5":
			fmt.Println("bye")
			return
		default:
			fmt.Println("error. please input a digit between 1 and 4 again.")
		}

	}

}
