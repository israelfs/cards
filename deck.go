package main

type deck []string

func (d deck) print() {
	for _, card := range d {
		println(card)
	}
}