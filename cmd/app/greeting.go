package main

import (
	"time"
)

type GreetingAnswer interface {
	GetGreeting() (uint, string)
	greet() (uint, string)
	greet_H_M() (uint, uint, string)
}

type Greeting struct {
	// use to show current time by just hour
	presentTime uint
	// use to show current time by hour and minute
	pHour   uint
	pMinute uint

	greeting string
}

func (g *Greeting) SetGreeting(presentTime, pHour, pMinute uint, greeting string) {
	g.presentTime = presentTime
	g.pHour = pHour
	g.pMinute = pMinute
	g.greeting = greeting
	g.greetHour()
	g.greet_H_M()
}

func (g *Greeting) GetGreeting() (uint, string) {
	return g.presentTime, g.greeting
}

// greet by hour
func (g *Greeting) greetHour() (uint, string) {
	g.presentTime = uint(time.Now().Hour())

	switch {
	case g.presentTime >= 6 || g.presentTime <= 9:
		g.greeting = "good moring"
	case g.presentTime > 12 || g.presentTime <= 15:
		g.greeting = "good afternoon"
	case g.presentTime > 17 || g.presentTime < 6:
		g.greeting = "good evening"
	default:
		g.greeting = "hello"
	}
	return g.presentTime, g.greeting
}

// greet by hour and minute
func (g *Greeting) greet_H_M() (uint, uint, string) {
	g.pHour = uint(time.Now().Hour())
	g.pMinute = uint(time.Now().Minute())

	switch {
	case g.presentTime >= 6 || g.presentTime <= 9:
		g.greeting = "good moring"
	case g.presentTime > 12 || g.presentTime <= 15:
		g.greeting = "good afternoon"
	case g.presentTime > 17 || g.presentTime < 6:
		g.greeting = "good evening"
	default:
		g.greeting = "hello"
	}

	return g.pHour, g.pMinute, g.greeting
}
