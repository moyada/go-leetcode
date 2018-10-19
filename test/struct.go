package main

import (
	"fmt"
)

type Id uint

func (id *Id) Increase(num int) {
	*id += Id(num)
}

type Person struct {
	name string
	age  uint8
}

func (person *Person) Say() string {
	return "Hello " + person.name
}

type Student struct {
	Person
	name  string
	class uint8
}

func (student *Student) Upgrade() {
	student.class++
}

func main() {
	id := Id(3592728412)
	id.Increase(1)
	(*Id).Increase(&id, 1)
	fmt.Println(id)

	p := &Person{"name", 12}
	rename(p)
	s := &Student{Person: Person{name: "s", age: 20}, name: "haha", class: 3}
	s.Upgrade()
	fmt.Println(p)
	fmt.Println(s)

	fmt.Println(p.Say())
}

func rename(p *Person) {
	p.name = "null"
}
