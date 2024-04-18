package models

import (
	"fmt"
	"sort"
	"time"
)

type RequestSoal1Model struct {
	Nums   []int `json:"nums"`
	Target int   `json:"target"`
}

type ResponseSoal1Model struct {
	Result []int `json:"result"`
	Time   time.Duration
}

type RequestSoal2Model struct {
	Nums []int `json:"nums"`
}

type ResponseSoal2Model struct {
	Result [][]int `json:"result"`
	Time   time.Duration
}

type RequestSoal3Model struct {
	Substring string   `json:"substring"`
	Words     []string `json:"words"`
}

type ResponseSoal3Model struct {
	Result []int
	Time   time.Duration
}

func Soal3Model(input RequestSoal3Model) *ResponseSoal3Model {
	start := time.Now()
	if len(input.Words) == 0 || len(input.Substring) == 0 {
		fmt.Println("result nil")
		return &ResponseSoal3Model{
			Result: nil,
			Time:   time.Since(start),
		}
	}

	wordLen := len(input.Words[0])
	wordsCount := len(input.Words)
	totalLen := wordLen * wordsCount
	result := []int{}

	wordsFreq := make(map[string]int)
	for {
		for _, word := range input.Words {
			wordsFreq[word]++
		}

		for i := 0; i <= len(input.Substring)-totalLen; i++ {
			substrFreq := make(map[string]int)
			j := i
			substr := input.Substring[j : j+totalLen]

			for len(substr) >= wordLen {
				word := substr[:wordLen]
				substrFreq[word]++
				if substrFreq[word] > wordsFreq[word] {
					break
				}
				j += wordLen
				substr = substr[wordLen:]
			}
			if len(substr) == 0 {
				result = append(result, i)
			}
		}
		if len(result) > 0 {
			return &ResponseSoal3Model{
				Result: result,
				Time:   time.Since(start),
			}
		}

	}
	
}

func Soal1Model(input RequestSoal1Model) *ResponseSoal1Model {
	start := time.Now()
	numMap := make(map[int]int)
	for i, num := range input.Nums {
		if index, ok := numMap[input.Target-num]; ok {
			return &ResponseSoal1Model{
				Result: []int{index, i},
				Time:   time.Since(start),
			}
		}
		numMap[num] = i
	}
	return &ResponseSoal1Model{
		Result: []int{},
		Time:   time.Since(start),
	}
}

func Soal2Model(input RequestSoal2Model) *ResponseSoal2Model {
	start := time.Now()
	sort.Ints(input.Nums)
	var result [][]int

	n := len(input.Nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && input.Nums[i] == input.Nums[i-1] {
			continue
		}
		left, right := i+1, n-1

		for left < right {
			sum := input.Nums[i] + input.Nums[left] + input.Nums[right]
			if sum == 0 {
				result = append(result, []int{input.Nums[i], input.Nums[left], input.Nums[right]})
				for left < right && input.Nums[left] == input.Nums[left+1] {
					left++
				}
				for left < right && input.Nums[right] == input.Nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return &ResponseSoal2Model{
		Result: result,
		Time:   time.Since(start),
	}
}
