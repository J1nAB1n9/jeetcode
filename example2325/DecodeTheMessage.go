package example2325

//////////////////////////////////////////////////////
//	执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户	//
//	内存消耗：2.3 MB, 在所有 Go 提交中击败了70.00% 的用户	//
//////////////////////////////////////////////////////

func decodeMessage(key string, message string) string {
	msgArr := []byte(message)
	keyMap := make(map[int32]int32, 26)
	num := 97
	for _, v := range key {
		if v < 97 || v > 122 {
			continue
		}
		if _, ok := keyMap[v]; !ok {
			keyMap[v] = int32(num)
			num++
		}
	}

	for i, v := range msgArr {
		if v >= 97 && v <= 122 {
			msgArr[i] = byte(keyMap[int32(v)])
		}
	}
	return string(msgArr)
}
