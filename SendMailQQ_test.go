package Common

import (
	"testing"
)

//func (this *SendMailQQ) SmtpSendMail(to []string, subject, body string) error
func TestSmtpSendMail(t *testing.T) {
	m := NewSendMailQQ("test", "test@mmm.com", "sdfaasdfdf", "smtp.qq.com", "smtp.qq.com:25")
	var to []string
	to = append(to, "00000000@qq.com")
	m.SmtpSendMail(to, "test Mail", "this is test mail. from golang.")
}
