package gosimhash

import (
	"testing"
	"github.com/HaoyuHu/gosimhash/utils"
)

func TestSimhashWithJenkins(t *testing.T) {
	hasher := NewSimpleSimhasher()
	defer hasher.Free()

	var sentence string = "我来到北京清华大学"
	var topN int = 5

	func() {
		var expected uint64
		var actual uint64

		expected = 0xfa596a42bb35f945
		actual = hasher.MakeSimhash(&sentence, topN)
		if expected != actual {
			t.Error(expected, "!=", actual)
		}
	}()

	func() {
		var expected string
		var actual string

		expected = "1111101001011001011010100100001010111011001101011111100101000101"
		actual = hasher.MakeSimhashString(&sentence, topN)
		if expected != actual {
			t.Error(expected, "!=", actual)
		}
	}()

	func() {
		var first uint64
		var second uint64

		first = hasher.MakeSimhash(&sentence, topN)
		second = hasher.MakeSimhash(&sentence, topN)
		if first != second {
			t.Error(first, "!=", second)
		}
	}()

	func() {
		distance := CalculateDistanceBySimhash(0x812e5cf1b47eb66, 0x812e5cf1b47eb61)
		if distance != 3 {
			t.Error(distance, "!= 3")
		}
	}()

	func() {
		distance := CalculateDistanceBySimhashString(
			"100000010010111001011100111100011011010001111110101101100110",
			"100000010010111001011100111100011011010001111110101101100001")
		if distance != 3 {
			t.Error(distance, "!= 3")
		}
	}()
}

func TestSimhashWithSipHash(t *testing.T) {
	sip := utils.NewSipHasher([]byte(DEFAULT_HASH_KEY))
	hasher := NewSimhasher(sip, "./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "",
		"./dict/idf.utf8", "./dict/stop_words.utf8")
	defer hasher.Free()

	var sentence string = "我来到北京清华大学"
	var topN int = 5

	func() {
		var expected uint64
		var actual uint64

		expected = 0x812e5cf1b47eb66
		actual = hasher.MakeSimhash(&sentence, topN)
		if expected != actual {
			t.Error(expected, "!=", actual)
		}
	}()

	func() {
		var expected string
		var actual string

		expected = "100000010010111001011100111100011011010001111110101101100110"
		actual = hasher.MakeSimhashString(&sentence, topN)
		if expected != actual {
			t.Error(expected, "!=", actual)
		}
	}()

	func() {
		var first uint64
		var second uint64

		first = hasher.MakeSimhash(&sentence, topN)
		second = hasher.MakeSimhash(&sentence, topN)
		if first != second {
			t.Error(first, "!=", second)
		}
	}()

	func() {
		distance := CalculateDistanceBySimhash(0x812e5cf1b47eb66, 0x812e5cf1b47eb61)
		if distance != 3 {
			t.Error(distance, "!= 3")
		}
	}()

	func() {
		distance := CalculateDistanceBySimhashString(
			"100000010010111001011100111100011011010001111110101101100110",
			"100000010010111001011100111100011011010001111110101101100001")
		if distance != 3 {
			t.Error(distance, "!= 3")
		}
	}()
}
