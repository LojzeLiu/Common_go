package Common

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SendMailQQ struct {
	_auth        smtp.Auth
	_user        string
	_nickname    string
	_contentType string
	_addr        string
}

func NewSendMailQQ(nickname, user, password, server, addr string) *SendMailQQ {
	tmp := &SendMailQQ{}
	tmp._auth = smtp.PlainAuth("", user, password, server)
	tmp._user = user
	tmp._nickname = nickname
	tmp._contentType = "Content-Type: text/plain; charset=UTF-8"
	tmp._addr = addr
	return tmp
}

func (this *SendMailQQ) SmtpSendMail(to []string, subject, body string) error {
	msg := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
		strings.Join(to, ","), this._nickname, this._user, subject, this._contentType, body)

	return smtp.SendMail(this._addr, this._auth, this._user, to, []byte(msg))
}
