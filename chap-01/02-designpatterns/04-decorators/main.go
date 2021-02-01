package main

import "fmt"

type IProcess interface {
	process()
}

type ProcessClass struct{}

func (pc *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

type ProcessDecorator struct {
	processClass *ProcessClass
}

func (pd *ProcessDecorator) process() {
	if pd.processClass == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Println("ProcessDecorator process")
		pd.processClass.process()
	}
}

func main() {
	processClass := &ProcessClass{}
	processDeco := &ProcessDecorator{
		processClass: processClass,
	}

	processDeco.process()
}
