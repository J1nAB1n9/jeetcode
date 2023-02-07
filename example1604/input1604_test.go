package example1604

import (
	"fmt"
	"testing"
)

//示例 1：
//
//输入：keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"], keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
//输出：["daniel"]
//解释："daniel" 在一小时内使用了 3 次员工卡（"10:00"，"10:40"，"11:00"）。
//
//示例 2：
//
//输入：keyName = ["alice","alice","alice","bob","bob","bob","bob"], keyTime = ["12:01","12:00","18:00","21:00","21:20","21:30","23:00"]
//输出：["bob"]
//解释："bob" 在一小时内使用了 3 次员工卡（"21:00"，"21:20"，"21:30"）。
//
//示例 3：
//
//输入：keyName = ["john","john","john"], keyTime = ["23:58","23:59","00:01"]
//输出：[]
//
//示例 4：
//
//输入：keyName = ["leslie","leslie","leslie","clare","clare","clare","clare"], keyTime = ["13:00","13:20","14:00","18:00","18:51","19:30","19:49"]
//输出：["clare","leslie"]

func TestDemo1(t *testing.T) {
	keyName := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	keyTime := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
	fmt.Println(alertNames(keyName, keyTime))
}

func TestDemo2(t *testing.T) {
	keyName := []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"}
	keyTime := []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}
	fmt.Println(alertNames(keyName, keyTime))
}

func TestDemo3(t *testing.T) {
	keyName := []string{"john", "john", "john"}
	keyTime := []string{"23:58", "23:59", "00:01"}
	fmt.Println(alertNames(keyName, keyTime))
}

func TestDemo4(t *testing.T) {
	keyName := []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"}
	keyTime := []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"}
	fmt.Println(alertNames(keyName, keyTime))
}
