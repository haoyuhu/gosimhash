package gosimhash

import (
	"strconv"
	"fmt"
	"github.com/HaoyuHu/gosimhash/utils"
	jieba "github.com/yanyiwu/gojieba"
)

const (
	BitsLength           = 64
	Binary               = 2
	DefaultHashKey       = "8b6555d0c9cff7a9"
	DefaultThresholdDist = 3
)

type Simhasher struct {
	extractor *jieba.Jieba
	hasher    utils.Hasher
}

type HashWeight struct {
	hash   uint64
	weight float64
}

func NewSimpleSimhasher() *Simhasher {
	var jenkinsHasher utils.Hasher = utils.NewJenkinsHasher()
	return NewSimhasher(jenkinsHasher, "", "", "", "", "")
}

func NewSimhasher(hasher utils.Hasher, dict string, hmm string, userDict string, idf string, stopWords string) *Simhasher {
	getDictPath(&dict, &hmm, &userDict, &idf, &stopWords)
	return &Simhasher{
		extractor: jieba.NewJieba(dict, hmm, userDict, idf, stopWords),
		hasher:    hasher}
}

func (simhasher *Simhasher) MakeSimhash(doc *string, topN int) uint64 {
	wws := simhasher.extractor.ExtractWithWeight(*doc, topN)
	size := len(wws)
	hws := make([]HashWeight, size, size)
	simhasher.convertWordWeights2HashWeights(wws, hws)
	var vector [BitsLength]float64
	var one uint64 = 1
	for _, hw := range hws {

		for i := 0; i < BitsLength; i++ {
			if ((one << uint(i)) & hw.hash) > 0 {
				vector[i] += hw.weight
			} else {
				vector[i] += -hw.weight
			}
		}
	}
	var ret uint64 = 0
	for i := 0; i < BitsLength; i++ {
		if vector[i] > 0.0 {
			ret |= one << uint(i)
		}
	}
	return ret
}

func (simhasher *Simhasher) MakeSimhashBinString(doc *string, topN int) string {
	simhash := simhasher.MakeSimhash(doc, topN)
	return strconv.FormatUint(simhash, Binary)
}

func (simhasher *Simhasher) Free() {
	simhasher.extractor.Free()
	simhasher.hasher = nil
}

func CalculateDistanceBySimhash(simhash uint64, another uint64) int {
	xor := simhash ^ another
	counter := 0
	for ; xor != 0; {
		xor &= xor - 1
		counter ++
	}
	return counter
}

func IsSimhashDuplicated(simhash uint64, another uint64, limit int) bool {
	xor := simhash ^ another
	counter := 0
	for ; xor != 0 && counter <= limit; {
		xor &= xor - 1
		counter ++
	}
	return counter <= limit
}

func CalculateDistanceBySimhashBinString(simhashStr string, anotherStr string) (int, error) {
	simhash, err := strconv.ParseUint(simhashStr, Binary, BitsLength)
	if err != nil {
		fmt.Printf("Cannot convert simHashStr(%s) to uint64 simhash: %s\n", simhashStr, err.Error())
		return 0, err
	}
	another, err := strconv.ParseUint(anotherStr, Binary, BitsLength)
	if err != nil {
		fmt.Printf("Cannot convert anotherStr(%s) to uint64 simhash: %s\n", anotherStr, err.Error())
		return 0, err
	}
	return CalculateDistanceBySimhash(simhash, another), nil
}

func IsSimhashBinStringDuplicated(simhashStr string, anotherStr string, limit int) (bool, error) {
	simhash, err := strconv.ParseUint(simhashStr, Binary, BitsLength)
	if err != nil {
		fmt.Printf("Cannot convert simHashStr(%s) to uint64 simhash: %s\n", simhashStr, err.Error())
		return false, err
	}
	another, err := strconv.ParseUint(anotherStr, Binary, BitsLength)
	if err != nil {
		fmt.Printf("Cannot convert anotherStr(%s) to uint64 simhash: %s\n", anotherStr, err.Error())
		return false, err
	}
	return IsSimhashDuplicated(simhash, another, limit), nil
}

func (simhasher *Simhasher) convertWordWeights2HashWeights(wws []jieba.WordWeight, hws []HashWeight) {
	for index, ww := range wws {
		hws[index].hash = simhasher.hasher.Hash64(ww.Word)
		hws[index].weight = ww.Weight
	}
}

func getDictPath(dict *string, hmm *string, userDict *string, idf *string, stopWords *string) {
	if *dict == "" {
		*dict = jieba.DICT_PATH
	}
	if *hmm == "" {
		*hmm = jieba.HMM_PATH
	}
	if *userDict == "" {
		*userDict = jieba.USER_DICT_PATH
	}
	if *idf == "" {
		*idf = jieba.IDF_PATH
	}
	if *stopWords == "" {
		*stopWords = jieba.STOP_WORDS_PATH
	}
}
