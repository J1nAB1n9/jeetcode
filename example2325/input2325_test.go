package example2325

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRangeString(t *testing.T) {
	ss := "abcdef"

	for _, c := range ss { //for _, c := range []byte(ss) {
		typec := reflect.TypeOf(c)
		fmt.Println(typec.Name(), typec.Kind())
		fmt.Println(c)
	}
}

//Input: key = "the quick brown fox jumps over the lazy dog", message = "vkbs bs t suepuv"
//Output: "this is a secret"

//Input: key = "eljuxhpwnyrdgtqkviszcfmabo", message = "zwx hnfx lqantp mnoeius ycgk vcnjrdb"
//Output: "the five boxing wizards jump quickly"

func TestDemo1(t *testing.T) {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	fmt.Println(decodeMessage(key, message))
}

func TestDemo2(t *testing.T) {
	key := "eljuxhpwnyrdgtqkviszcfmabo"
	message := "zwx hnfx lqantp mnoeius ycgk vcnjrdb"
	fmt.Println(decodeMessage(key, message))
}
