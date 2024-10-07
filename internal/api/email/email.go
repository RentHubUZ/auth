package email

import (
	"auth/internal/config"
	"auth/internal/storage/redis"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"

	"github.com/pkg/errors"
)

func SendEmail(ctx context.Context, cfg *config.Config, rdb *redis.RedisDB, email string) error {
	code := strconv.Itoa(rand.Intn(900000) + 100000)

	err := rdb.StoreCode(ctx, email, code)
	if err != nil {
		return err
	}

	c, err := rdb.GetCode(ctx, email)
	if err != nil {
		return err
	}

	if c != code {
		return errors.New("internal error: invalid code")
	}

	err = sendCode(cfg, email, c)
	if err != nil {
		return err
	}

	return nil
}

func sendCode(cfg *config.Config, email string, code string) error {
	from := cfg.EMAIL
	password := cfg.APP_KEY

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("internal/api/email/template.html")
	if err != nil {
		return errors.Wrap(err, "failed to parse template")
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Your verification code \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Passwd string
	}{

		Passwd: code,
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}

	log.Println("Email sended to:", email)
	return nil
}
