package model

func ChangeString(str string, from int, end int) string {
	tmpStr := []byte(str)
	tmpStr1 := tmpStr[from-1:end-1]
	return string(tmpStr1)
}
