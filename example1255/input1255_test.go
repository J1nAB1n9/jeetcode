package example1255

import (
	"fmt"
	"testing"
)

//示例 1：
//
//输入：words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"], score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
//输出：23
//
//示例 2：
//
//输入：words = ["xxxz","ax","bx","cx"], letters = ["z","a","b","c","x","x","x"], score = [4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]
//输出：27
//
//示例 3：
//
//输入：words = ["leetcode"], letters = ["l","e","t","c","o","d"], score = [0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
//输出：0

func TestDemo1(t *testing.T) {
	words := []string{"xxxz", "ax", "bx", "cx"}
	letters := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(maxScoreWords(words, letters, score))
}

func TestDemo2(t *testing.T) {
	words := []string{"xxxz", "ax", "bx", "cx"}
	letters := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(maxScoreWords(words, letters, score))
}

func TestDemo3(t *testing.T) {
	words := []string{"xxxz", "ax", "bx", "cx"}
	letters := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(maxScoreWords(words, letters, score))
}

func TestDemo4(t *testing.T) {
	words := []string{"xxxz", "ax", "bx", "cx"}
	letters := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(maxScoreWords(words, letters, score))
}

func TestDemo5(t *testing.T) {
	words := []string{"xxxz", "ax", "bx", "cx"}
	letters := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(maxScoreWords(words, letters, score))
}
