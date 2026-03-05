package main

type Calctulator interface {
	GetCalc() *Calc
	CalcAdd()
	CalcMul()
	CalcDiv()
	CalcMod()
}

type Calc struct {
	formAdd, formMul, formDiv, formMod               string
	x, y, addResult, mulResult, divResult, modResult int
}

func (c *Calc) SetCalc(formAdd, formMul, formDiv, formMod string, x, y, addResult, mulResult, divResult, modResult int) {
	c.formAdd = formAdd
	c.formMul = formMul
	c.formDiv = formDiv
	c.formMod = formMod
	c.x = x
	c.y = y
	c.addResult = addResult
	c.mulResult = mulResult
	c.divResult = divResult
	c.modResult = modResult
}

func (c *Calc) GetCalc() *Calc {
	return c
}

func (c *Calc) CalcAdd() (string, int, int, int) {
	c.addResult = c.x + c.y
	c.formAdd = "Addition:"
	return c.formAdd, c.x, c.y, c.addResult
}

func (c *Calc) CalcMul() (string, int, int, int) {
	c.mulResult = c.x * c.y
	c.formMul = "Multiplication:"
	return c.formMul, c.x, c.y, c.mulResult
}

func (c *Calc) CalcDiv() (string, int, int, int) {
	c.divResult = c.x / c.y
	c.formDiv = "Division:"
	return c.formDiv, c.x, c.y, c.divResult
}

func (c *Calc) CalcMod() (string, int, int, int) {
	c.modResult = c.x % c.y
	c.formMod = "Modulo:"
	return c.formMod, c.x, c.y, c.modResult
}
