package util

import (
	"bytes"
	"github.com/go-gomail/gomail"
	"strconv"
	"time"
)

//type1: 1 -> 白天的需求 2 -> 夜晚的吐槽  3 -> 登录页面的反馈
//content: 被举报的内容
//reason: 被举报的原因
//reporter: 举报人
//person: 被举报人
//addition: 附加信息
func SendMail(type1 int, content, reason, reporter, person, addition string) error {

	mailConn := map[string]string{
		"user": "3243837480@qq.com",
		"pass": "mywhictdshrvdbdj",
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	var bt bytes.Buffer

	bt.WriteString("主类别: ")
	tmp := convert1(type1)
	bt.WriteString(tmp)
	bt.WriteString("<br><br>")
	bt.WriteString("举报人: ")
	bt.WriteString(reporter)
	bt.WriteString("<br><br>")
	bt.WriteString("被举报人: ")
	bt.WriteString(person)
	bt.WriteString("<br><br>")
	bt.WriteString("原因: ")
	tmp = convert2(reason)
	bt.WriteString(tmp)
	bt.WriteString("<br><br>")
	bt.WriteString("附加信息: ")
	bt.WriteString(addition)
	bt.WriteString("<br><br>")
	bt.WriteString("被举报内容: ")
	bt.WriteString(content)
	bt.WriteString("<br><br>")
	bt.WriteString("举报时间: ")
	bt.WriteString(time.Now().String())
	bt.WriteString("<br><br>")
	bt.WriteString(`<a href="https://imgchr.com/i/JVB0mD"><img src="https://s1.ax1x.com/2020/04/17/JVB0mD.th.jpg" alt="JVB0mD.jpg" border="0" /></a>`)
	body := bt.String()
	//fmt.Println(body)
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "孤独星球"))
	m.SetHeader("To", "3243837480@qq.com") //发送给多个用户
	m.SetHeader("Subject", "新举报提醒")        //设置邮件主题
	m.SetBody("text/html", body)           //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func convert1(num int) string {
	switch num {
	case 1:
		return "白天的需求"
	case 2:
		return "黑夜的吐槽"
	case 3:
		return "通过首页的反馈"
	}
	return ""
}

func convert2(str string) string {
	var bt bytes.Buffer
	for _, v := range str {
		switch v {
		case '1':
			bt.WriteString("色情低俗")
		case '2':
			bt.WriteString("暴力血腥")
		case '3':
			bt.WriteString("政治敏感")
		case '4':
			bt.WriteString("欺诈骗财")
		case '5':
			bt.WriteString("人参攻击")
		case '6':
			bt.WriteString("侵犯隐私")
		case '7':
			bt.WriteString("广告骚扰")
		case '8':
			bt.WriteString("侮辱谩骂")
		case '9':
			bt.WriteString("其他")
		}
		bt.WriteString(" ")
	}
	result := bt.String()
	return result
}
