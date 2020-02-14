package model

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func ChangeString(str string, from int, end int) string {
	tmpStr := []byte(str)
	tmpStr1 := tmpStr[from-1 : end]
	return string(tmpStr1)
}

func ConvertStringToIntSlice(str string) ([]int, error) {
	var intSlice []int
	if len(str) == 0 {
		return intSlice, nil
	}
	for i := 0; i <= len(str)-1; i++ {
		b, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return intSlice, err
		}
		intSlice = append(intSlice, b)
	}
	return intSlice, nil
}

/*
func ConverType2String1(n int) string {
	if n == 1 {
		return
	}
}
*/

func NowTime() string {
	timeNow := time.Now()
	t := timeNow.Format("2006-01-02 15:04")
	return t
}

func NowTimeStampStr() string {
	str := strconv.FormatInt(time.Now().Unix(), 10)
	return str
}

func Dec2BinStr(num int) string {
	var s string
	for num > 0 {
		a := num % 2
		s = strconv.Itoa(a) + s
		num = num / 2
	}
	return s
}

func BinStr2Dec(str string) (int, error) {
	tmpResult, err := strconv.ParseInt(str, 2, 32)
	if err != nil {
		log.Print("convert err")
		fmt.Println(err)
		return 0, err
	}
	result := int(tmpResult)
	return result, nil
}
