package model

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"strconv"
	"time"
)

func mainTypeImprove(type1 int) string {
	switch type1 {
	case 1:
		return "学习"
	case 2:
		return "运动"
	case 3:
		return "娱乐"
	case 4:
		return "其他"
	}
	return ""
}

func placeImprove(place int, type1 int) string {
	if type1 == 1 {
		switch place {
		case 1:
			return "教学楼"
		case 2:
			return "图书馆"
		case 3:
			return "其他"
		}
	}
	if type1 == 2 {
		switch place {
		case 1:
			return "佑铭体育馆"
		case 2:
			return "高职体育场"
		case 3:
			return "乒羽中心"
		case 4:
			return "西区篮球场"
		case 5:
			return "高职篮球场"
		case 6:
			return "学子篮球场"
		case 7:
			return "学子网球场"
		case 8:
			return "其他"
		}
	}
	if type1 == 3 {
		switch place {
		case 1:
			return "校内"
		case 2:
			return "校外"
		case 3:
			return "其他"
		}
	}
	if type1 == 4 {
		switch place {
		case 1:
			return "校内"
		case 2:
			return "校外"
		case 3:
			return "其他"
		}
	}
	return ""
}

func tagImprove(tag int, type1 int) string {
	if type1 == 1 {
		switch tag {
		case 1:
			return "自习"
		case 2:
			return "看书"
		case 3:
			return "上课"
		case 4:
			return "其他"
		}
	}
	if type1 == 2 {
		switch tag {
		case 1:
			return "跑步"
		case 2:
			return "篮球"
		case 3:
			return "足球"
		case 4:
			return "羽毛球"
		case 5:
			return "乒乓球"
		case 6:
			return "网球"
		case 7:
			return "其他"
		}
	}
	if type1 == 3 {
		switch tag {
		case 1:
			return "游戏"
		case 2:
			return "吃饭"
		case 3:
			return "电影"
		case 4:
			return "其他"
		}
	}
	return ""
}

func dateImprove(date string) string {
	var result string
	for i, v := range date {
		if v != 48 {
			if len(result) == 0 {
				result = date2(len(date)-i)
			}else {
				result = result + " " +date2(len(date)-i)
			}
		}
	}
	return result
}

func date2(date int) string {
	switch date {
	/*	case 0:
		return "" */
	case 1:
		return "周一"
	case 2:
		return "周二"
	case 3:
		return "周三"
	case 4:
		return "周四"
	case 5:
		return "周五"
	case 6:
		return "周六"
	case 7:
		return "周日"
	}
	return ""
}

func timestamp2json(str string) string {
	tmpTime, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Print("timestamp2json err:")
		fmt.Println(err)
	}
	result := time.Unix(tmpTime, 0).Format("01-02 15:04")
	return result
}