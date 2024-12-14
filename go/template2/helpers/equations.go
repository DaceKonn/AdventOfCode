package helpers

import (
	"fmt"
	"strings"
)

type Calculator interface {
	Calculate() int
	Print() string
}

type Element struct {
	Value int
}

func NewElement(value int) Element {
	// fmt.Printf("\t\tcreating element: %d\n", value)
	return Element{Value: value}
}

func (e Element) Calculate() int {
	return e.Value
}

func (e Element) Print() string {
	return fmt.Sprintf("%d", e.Value)
}

type Add struct {
	elements []Calculator
}

func NewAdd() Add {
	// fmt.Println("\t\tcreating Add")
	return Add{}
}

func (a *Add) AddElement(el Calculator) {
	// fmt.Println("\t\t\t Adding element")
	a.elements = append(a.elements, el)
}

func (a Add) Calculate() int {
	var base = 0
	for _, el := range a.elements {
		base += el.Calculate()
	}
	return base
}

func (a Add) Print() string {
	var outSlice []string
	for _, el := range a.elements {
		outSlice = append(outSlice, el.Print())
	}
	switch len(outSlice) {
	case 0:
		return "emptyAdd"
	case 1:
		return fmt.Sprintf("singleAdd(%s)", outSlice[0])
	default:
		return strings.Join(outSlice, " + ")
	}
}

type Multiply struct {
	elements []Calculator
}

func NewMultiply() Multiply {
	// fmt.Println("\t\tcreating multiply")
	return Multiply{}
}

func (m *Multiply) AddElement(el Calculator) {
	// fmt.Println("\t\t\t Adding element")
	m.elements = append(m.elements, el)
}

func (m Multiply) Calculate() int {
	var base int = 1
	for _, el := range m.elements {
		base *= el.Calculate()
	}
	return base
}

func (m Multiply) Print() string {
	var outSlice []string
	for _, el := range m.elements {
		outSlice = append(outSlice, el.Print())
	}
	switch len(outSlice) {
	case 0:
		return "emptyMultiply"
	case 1:
		return fmt.Sprintf("singleMultiply(%s)", outSlice[0])
	default:
		return strings.Join(outSlice, " * ")
	}
}

type Concentrate struct {
	elements []Calculator
}

func NewConcentrate() Concentrate {
	// fmt.Println("\t\tcreating Concentrate")
	return Concentrate{}
}

func (a *Concentrate) AddElement(el Calculator) {
	// fmt.Println("\t\t\t Concentrateing element")
	a.elements = append(a.elements, el)
}

func (a Concentrate) Calculate() int {
	var tmp string = ""
	for _, el := range a.elements {
		tmp = strings.Join([]string{tmp, fmt.Sprint(el.Calculate())}, "")
	}
	return ParseInt(tmp)
}

func (a Concentrate) Print() string {
	var outSlice []string
	for _, el := range a.elements {
		outSlice = append(outSlice, el.Print())
	}
	switch len(outSlice) {
	case 0:
		return "emptyConcentrate"
	case 1:
		return fmt.Sprintf("singleConcentrate(%s)", outSlice[0])
	default:
		return strings.Join(outSlice, " || ")
	}
}
