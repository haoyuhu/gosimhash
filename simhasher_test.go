package gosimhash

import (
	"testing"
	"github.com/HaoyuHu/gosimhash/utils"
)

func TestSimhashWithJenkins(t *testing.T) {
	hasher := NewSimpleSimhasher()
	defer hasher.Free()

	var sentence = "我来到北京清华大学"
	var topN = 5

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
		actual = hasher.MakeSimhashBinString(&sentence, topN)
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
		distance, err := CalculateDistanceBySimhashBinString(
			"100000010010111001011100111100011011010001111110101101100110",
			"100000010010111001011100111100011011010001111110101101100001")
		if err != nil {
			t.Error(err.Error())
			return
		}
		if distance != 3 {
			t.Error(distance, "!= 3")
		}
	}()
}

func TestSimhashWithSipHash(t *testing.T) {
	sip := utils.NewSipHasher([]byte(DefaultHashKey))
	hasher := NewSimhasher(sip, "./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "",
		"./dict/idf.utf8", "./dict/stop_words.utf8")
	defer hasher.Free()

	var sentence = "我来到北京清华大学"
	var topN = 5

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
		actual = hasher.MakeSimhashBinString(&sentence, topN)
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
		distance, err := CalculateDistanceBySimhashBinString(
			"100000010010111001011100111100011011010001111110101101100110",
			"100000010010111001011100111100011011010001111110101101100001")
		if err != nil {
			t.Error(err.Error())
			return
		} else if distance != 3 {
			t.Error(distance, "!= 3")
		}

		distance, err = CalculateDistanceBySimhashBinString(
			"100000010010111001011100111100011011010001111110101101100113",
			"100000010010111001011100111100011011010001111110101101100001")
		if err == nil {
			t.Error("Should throw error at CalculateDistanceBySimhashBinString")
		}

		distance, err = CalculateDistanceBySimhashBinString(
			"100000010010111001011100111100011011010001111110101101100110",
			"100000010010111001011100111100011011010001111110101101100003")
		if err == nil {
			t.Error("Should throw error at CalculateDistanceBySimhashBinString")
		}
	}()

	func() {
		duplicated := IsSimhashDuplicated(0x812e5cf1b47eb66, 0x812e5cf1b47eb61, 3)
		if !duplicated {
			t.Error("Should be duplicated at IsSimhashDuplicated")
		}
		duplicated, err := IsSimhashBinStringDuplicated(
			"100000010010111001011100111100011011010001111110101101100110",
			"100000010010111001011100111100011011010001111110101101100001", 3)
		if err != nil {
			t.Error(err.Error())
		} else if !duplicated {
			t.Error("Should be duplicated at IsSimhashBinStringDuplicated")
		}
		duplicated, err = IsSimhashBinStringDuplicated(
			"100000010010111001011100111100011011010001111110101101100113",
			"100000010010111001011100111100011011010001111110101101100001", 3)
		if err == nil {
			t.Error("Should throw error at IsSimhashBinStringDuplicated")
		}
		duplicated, err = IsSimhashBinStringDuplicated(
			"100000010010111001011100111100011011010001111110101101100110",
			"100000010010111001011100111100011011010001111110101101100003", 3)
		if err == nil {
			t.Error("Should throw error at IsSimhashBinStringDuplicated")
		}
	}()
}
