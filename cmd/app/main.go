package main

import "fmt"

func main() {
	cal := &Calc{}
	cal.SetCalc("", "", "", "", 198, 231, 0, 0, 0, 0)
	// addition
	formAdd, x, y, addResult := cal.CalcAdd()
	fmt.Printf("\n%s %d + %d = %d", formAdd, x, y, addResult)
	// mutiplication
	formMul, x, y, mulResult := cal.CalcMul()
	fmt.Printf("\n%s %d * %d = %d", formMul, x, y, mulResult)

	// division
	formDiv, x, y, divResult := cal.CalcDiv()
	fmt.Printf("\n%s %d / %d = %d", formDiv, x, y, divResult)
	// modulo
	formMod, x, y, modResult := cal.CalcMod()
	fmt.Printf("\n%s %d mod %d = %d", formMod, x, y, modResult)

	// greeting
	gr := &Greeting{}
	gr.SetGreeting(0, 0, 0, "")
	// show greeting and current hour
	fmt.Printf("\nit´s %d now. %s", gr.presentTime, gr.greeting)
	// show greeting, current hour and minute
	fmt.Printf("\nit´s %02d:%02d now. %s", gr.pHour, gr.pMinute, gr.greeting)
}
