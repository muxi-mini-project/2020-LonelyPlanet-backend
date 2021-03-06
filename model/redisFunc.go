package model

import (
	"bytes"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
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
		bt.WriteString("daySingle")
		bt.WriteString("*")
	case 2:
		bt.WriteString("nightSingle")
		bt.WriteString("*")
	case 3:
		bt.WriteString("dayTotal")
		bt.WriteString(now())
	case 4:
		bt.WriteString("nightTotal")
		bt.WriteString(now())
	}
	tmp := bt.String()
	newRedis := RedisDb.Self.Get()
	defer newRedis.Close()
	if type1 == 1 || type1 == 2 {
		tmp2, err := redis.Strings(newRedis.Do("keys", tmp))
		if err != nil {
			log.Println("keys", err)
			return result, err
		}
		result = len(tmp2)
		return result, nil
	}
	exists, err := redis.Bool(newRedis.Do("exists", tmp))
	if err != nil {
		log.Println("exits", err)
		return result, err
	}
	result = 0
	if exists {
		result, err := redis.Int(newRedis.Do("get", tmp))
		if err != nil {
			log.Println("get", err)
			return result, err
		}
		return result, nil
	}
	return result, nil
}

//调用此函数前一定先调用Num, redis效率较高, 所以多次调用后手动判断较好, 避免多次重置计时
//sec 自定义过期时间
func NewRecode(userId string, type1 int, sec int) error {
	var bt bytes.Buffer
	var bt2 bytes.Buffer
	bt.WriteString(userId)
	bt2.WriteString(userId)
	switch type1 {
	case 1:
		bt.WriteString("daySingle")
		bt.WriteString(nowShort())
		bt2.WriteString("dayTotal")
		bt2.WriteString(now())
	case 2:
		bt.WriteString("nightSingle")
		bt.WriteString(nowShort())
		bt2.WriteString("nightTotal")
		bt2.WriteString(now())
	}
	tmp := bt.String()
	tmp2 := bt2.String()
	newRedis := RedisDb.Self.Get()
	defer newRedis.Close()
	_, err := newRedis.Do("set", tmp, 1)
	if err != nil {
		log.Println("newrecode1", err)
		return err
	}
	_, err = newRedis.Do("expire", tmp, sec)
	if err != nil {
		log.Println("newrecode2", err)
		return err
	}
	exists, err := redis.Bool(newRedis.Do("exists", tmp2))
	if err != nil {
		log.Println("newrecode3", err)
		return err
	}
	_, err = newRedis.Do("incr", tmp2)
	if err != nil {
		log.Println("newrecode4", err)
		return err
	}
	if !exists {
		_, err = newRedis.Do("expireat", tmp2, nextDay())
		if err != nil {
			log.Println("newrecode5", err)
			return err
		}
	}

	return nil
}

func nowShort() string {
	tmp := strconv.FormatInt(time.Now().Unix(), 10)
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
