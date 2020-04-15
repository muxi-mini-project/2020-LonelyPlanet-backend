package model

import (
	"bytes"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

//type1: 1 -> 1分钟内白天的发布需求, 2 -> 1分钟内黑夜的发布秘密, 3 -> 一天的白天的发布需求, 4 -> 一天的黑夜发布需求
//            白天发布需求               黑夜发布秘密
func InspectNum(userId string, type1 int) (int, error) {
	var result int
	var bt bytes.Buffer
	bt.WriteString(userId)
	switch type1 {
	case 1:
		bt.WriteString("daySingle"); bt.WriteString("*")
	case 2:
		bt.WriteString("nightSingle"); bt.WriteString("*")
	case 3:
		bt.WriteString("dayTotal"); bt.WriteString(now())
	case 4:
		bt.WriteString("nightTotal"); bt.WriteString(now())
	}
	tmp := bt.String()
	if type1 == 1 || type1 == 2 {
		tmp2, err := redis.Strings(RedisDb.Self.Do("keys", tmp))
		if err != nil {
			return result, err
		}
		result = len(tmp2)
		return result, nil
	}
	result, err := redis.Int(RedisDb.Self.Do("get", tmp))
	if err != nil {
		return result, err
	}
	return result, nil
}

//调用此函数前一定先调用Num, redis效率较高, 所以多次调用后手动判断较好, 避免多次重置计时
func NewRecode(userId string, type1 int) error {
	var bt bytes.Buffer
	var bt2 bytes.Buffer
	bt.WriteString(userId)
	bt2.WriteString(userId)
	switch type1 {
	case 1:
		bt.WriteString("daySingle"); bt.WriteString(nowShort()); bt2.WriteString("dayTotal"); bt2.WriteString(now())
	case 2:
		bt.WriteString("nightSingle"); bt.WriteString(nowShort()); bt2.WriteString("nightTotal"); bt2.WriteString(now())
	}
	tmp := bt.String()
	tmp2 := bt2.String()
	_, err := RedisDb.Self.Do("set", tmp, 1)
	if err != nil {
		return err
	}
	_, err = RedisDb.Self.Do("expire", tmp, 60)
	if err != nil {
		return err
	}
	exists, err := redis.Bool(RedisDb.Self.Do("exists", tmp2))
	if err != nil {
		return err
	}
	_, err = RedisDb.Self.Do("incr", tmp2)
	if err != nil {
		return err
	}
	if !exists {
		_, err = RedisDb.Self.Do("expireat", tmp2, nextDay())
		if err != nil {
			return err
		}
	}

	return nil
}


func nowShort() string {
	tmp := strconv.FormatInt(time.Now().Unix(),10)
	return tmp[6:]
}

func now() string {
	now := time.Now().Format("01-02")
	//fmt.Println(now)
	return now
}

func nextDay() string {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	result := strconv.FormatInt(next.Unix(), 10)
	return result
}
