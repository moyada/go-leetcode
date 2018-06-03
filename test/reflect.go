package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string

	Age int
}

type User struct {
	Person

	Id int
}

func (p Person) GetName() string {
	return p.Name
}

func (u *User) UpdageAge(age int) bool {
	if age < 0 {
		return false
	}
	u.Age = age
	return true
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	if t.Kind() != reflect.Struct {
		fmt.Println("Type Error")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("type: ", t.PkgPath(), t.Name(), t.Kind())
	fmt.Println("value: ", v)

	fmt.Println(t.FieldByIndex([]int{0, 0}))

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i), v.Field(i).Interface())
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i))
	}
}

func set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Type Error")
		return
	}

	v = v.Elem()
	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("Field Invalid")
		return
	}

	f.Set(reflect.ValueOf("haha"))
	f.SetString("666")
}

func callMethod(o interface{}, method string, val interface{}) {
	v := reflect.ValueOf(o)
	m := v.MethodByName(method)
	args := []reflect.Value{reflect.ValueOf(val)}
	m.Call(args)
}

func main() {
	u := User{Person: Person{Name: "ok", Age: 20}, Id: 100}
	info(u)
	set(&u)
	callMethod(&u, "UpdageAge", 10)
	fmt.Println(u)
}
