package main

type Ninja struct {
	Id string
	Name string
	Level int
	Status NinjaStatus
}

type Dojo struct {
	Id string
	Name string
	AssociatedNinjas []Ninja
	Status DojoStatus
}
