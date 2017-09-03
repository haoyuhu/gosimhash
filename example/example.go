package main

import (
	"flag"
	"fmt"
	"gosimhash"
)

var sentence = flag.String("sentence", "我来到北京清华大学", "Sentence used to be hash")
var top_n = flag.Int("top_n", 5, "Top n of the words separated by jieba")

func main() {
	flag.Parse()

	hasher := gosimhash.NewSimpleSimhasher()
	defer hasher.Free()
	fingerprint := hasher.MakeSimhash(sentence, *top_n)
	fmt.Printf("%s simhash: %x\n", *sentence, fingerprint)
}
