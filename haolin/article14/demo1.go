package main

import "fmt"

type Pet interface {
	setName(name string)
	Name() string
	Category() string
}
type Dog struct {
	name string //名字
}

func (dog *Dog) setName(name string) {
	dog.name = name
}
func (dog Dog) Name() string {
	return dog.name
}
func (dog Dog) Category() string {
	return "dog"
}
func main() {
	//示例1
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	//dog2 := Dog{name: "little pig"}
	fmt.Println()
	//示例2
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}
