package main

func main(){
	cards := deck{"First one", "Second one"}
	cards = append(cards, "Third one")

	cards.print()
}