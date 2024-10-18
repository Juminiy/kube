package email

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestNewSMTPSender(t *testing.T) {
	inst := NewSMTPSender(SMTPAuthConfig{
		Username: "chinome@126.com",
		Password: "RDThvz99FeVHjmKG",
		Host:     "smtp.126.com",
		Port:     25,
	})

	util.Must(inst.WithMessage(util.String2BytesNoCopy("no-message")).
		WithReceiver("chisato-x@qq.com").
		Alone())

	util.Must(inst.WithMessage(util.String2BytesNoCopy("no-message")).
		WithReceiver("chisato-x@qq.com").
		Alone())
}
