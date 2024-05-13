package main

import (
	"english/control"
	"english/view"
)

func InIt() {

	wc := &control.WordControl{}

	wordlist := wc.GetWord()
	view := &view.View{
		Wordlist: wordlist,
	}
	view.ViewInit()
}
func main() {
	InIt()
}
