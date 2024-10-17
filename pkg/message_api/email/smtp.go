package email

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"net/smtp"
)

type (
	SMTPAuthConfig struct {
		Username   string
		Password   string
		Host       string
		Port       int
		ClientAddr string
	}
	SMTPSender struct {
		receivers []string
		message   []byte

		config *SMTPAuthConfig
		addr   string
		auth   smtp.Auth
		err    []error
	}
)

var (
	NoIdentity = ""
)

func NewSMTPSender(config SMTPAuthConfig) *SMTPSender {
	return &SMTPSender{
		addr:   fmt.Sprintf("%s:%d", config.Host, config.Port),
		auth:   smtp.PlainAuth(NoIdentity, config.Username, config.Password, config.Host),
		config: &config,
	}
}

func (s *SMTPSender) WithReceiver(addr ...string) *SMTPSender {
	s.receivers = addr
	s.err = make([]error, 0, len(addr))
	return s
}

func (s *SMTPSender) WithMessage(msg []byte) *SMTPSender {
	s.message = msg
	return s
}

// Group
// send group email by goroutine
func (s *SMTPSender) Group() error {
	return util.MergeError(s.err...)
}

// Alone
// send single email for distinct service
func (s *SMTPSender) Alone() error {
	if len(s.message) == 0 || len(s.receivers) == 0 {
		return nil
	}
	s.receivers = s.receivers[:1]
	s.sendMail()
	return util.MergeError(s.err...)
}

func (s *SMTPSender) sendMail() {
	s.err = append(s.err,
		smtp.SendMail(s.addr, s.auth, s.config.ClientAddr, s.receivers, s.message))
}
