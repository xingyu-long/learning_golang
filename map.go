package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

// name has the tag information.
type Animal struct {
	name string `required max:"100"`
	origin string
}

type Bird struct {
	Animal
	speedKPH float32
	canFly bool
}



func main() {
	populations := make(map[string]int)
	populations = map[string]int{
		"A": 1,
		"B": 2,
	}
	populations["C"] = 3
	fmt.Println(populations)
	delete(populations, "C")
	fmt.Println(populations)
	fmt.Println(populations["C"]) // will return 0 if the key doesn't exist
	_, isExist := populations["C"]
	fmt.Println(isExist)
	sp := populations // it's also reference of `populations`
	delete(sp, "B")
	fmt.Println(sp)
	fmt.Println(populations)

	doctor := Doctor{
		number: 3,
		actorName: "Test",
		companions: []string {
			"Test",
			"test",
		},
	}
	fmt.Println(doctor)
	fmt.Println(doctor.number)
	fmt.Println(doctor.companions)


	anotherDoctor := struct{name string}{name: "Clark Long"}
	thirdDoctor := anotherDoctor
	thirdDoctor.name = "Joe"
	fmt.Println(anotherDoctor)
	fmt.Println(thirdDoctor)

	b := Bird{}
	b.name = "Emu"
	b.origin = "USA"
	b.speedKPH = 48
	b.canFly = false
	fmt.Println(b)

	c := Bird{
		Animal: Animal{name: "emu", origin: "UK"},
		speedKPH: 100,
		canFly: false,
	}
	fmt.Println(c)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}