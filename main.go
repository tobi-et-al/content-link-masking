package main

import "os"

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	input := (args[1])
	space := " "
	b := []byte(space)

	spaceLiteral := (b[0])
	var bwords []byte = nil
	var sentense [][]byte = nil

	for _, val := range input {

		if byte(val) == spaceLiteral {
			sentense = append(sentense, bwords)
			bwords = nil
			continue
		}

		bwords = append(bwords, byte(val))
	}

	var spam []byte

	for k, word := range sentense {
		if k != 0 {
			spam = append(spam, []byte(" ")...)
		}

		if string(word[0:7]) == "http://" {
			var hash []byte
			var protocol = []byte("http://")

			hash = append(hash, protocol...)

			for i := 0; i < len(word)-1; i++ {
				b := []byte("*")
				hash = append(hash, b[0])
			}

			word = hash
		}

		spam = append(spam, []byte(word)...)

	}
	println(string(spam))

}
