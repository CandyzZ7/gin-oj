package helper

import (
	"crypto/tls"
	"gin_oj.cn/define"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <tqj13110708227@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "tqj13110708227@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
