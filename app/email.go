package app

import (
	"context"
	"math/rand"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("0123456789")

func randomCode(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (m module) SendEmail(ctx context.Context, sender, recipient, subject, template string, vars map[string]string) error {
	mg := mailgun.NewMailgun(m.MgDomain, m.MgKey)
	mg.SetAPIBase(mailgun.APIBaseEU)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, "", recipient)
	message.SetTemplate(template)
	for k, v := range vars {
		message.AddTemplateVariable(k, v)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)

	return err
}

func (m module) SendWelcomeEmail(ctx context.Context, recipient, name string) error {
	mg := mailgun.NewMailgun(m.MgDomain, m.MgKey)
	mg.SetAPIBase(mailgun.APIBaseEU)

	message := mg.NewMessage("noreply@defipay.tech", "Welcome To Defipay", "", recipient)
	message.SetTemplate("defipay-welcome")
	message.AddTemplateVariable("name", name)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)

	return err
}
