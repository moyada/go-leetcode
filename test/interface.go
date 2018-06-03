package main

import (
	"fmt"
)

type USB interface {
	Connecter

	Name() string
}

type Connecter interface {
	Connect()
}

type Phone struct {
	name string
}

func (p Phone) Name() string {
	return p.name
}

func (p Phone) Connect() {
	fmt.Println(p.name, "connect")
}

func main() {
	var p USB = Phone{"xiaomi"}
	fmt.Println(p.Name())
	p.Connect()
	Disconnect(p)
}

func Disconnect(u interface{}) {
	if p, ok := u.(Phone); ok {
		fmt.Println(p.name, "disconnect")
	}

	switch v := u.(type) {
	case Phone:
		fmt.Println(v.name, "disconnect")
	default:
		fmt.Println("Unknow decive")
	}
}
