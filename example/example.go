package main

import (
	"flag"
	"fmt"
	"github.com/HaoyuHu/gosimhash"
)

var sentence = flag.String("sentence", "今天的天气确实适合户外运动", "Sentence for simhash")
var topN = flag.Int("top_n", 5, "Top n of the words separated by jieba")

func main() {
	flag.Parse()

	hasher := gosimhash.NewSimpleSimhasher()
	defer hasher.Free()
	fingerprint := hasher.MakeSimhash(sentence, *topN)
	binary := hasher.MakeSimhashBinString(sentence, *topN)
	fmt.Printf("sentence: %s, simhash in uint64: %x\n", *sentence, fingerprint)
	fmt.Printf("sentence: %s, simhash in binary: %s\n", *sentence, binary)

}
