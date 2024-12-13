package email

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestNewSMTPSender(t *testing.T) {
	inst := NewSMTPSender(SMTPAuthConfig{
		Username: "-",
		Password: "-",
		Host:     "smtp.126.com",
		Port:     25,
	})

	util.Must(inst.WithMessage(util.String2BytesNoCopy("no-message")).
		WithReceiver("-").
		Alone())

	util.Must(inst.WithMessage(util.String2BytesNoCopy("no-message")).
		WithReceiver("-").
		Alone())
}
