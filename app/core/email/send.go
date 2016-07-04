//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package email

import (
	"fmt"
	"time"

	"github.com/revel/revel"
)

var (
	sendChannel chan *Email
	config      string
)

var mate = `
========================
=== 开启mail workpool  %d
========================
`

func InitEmail() {
	conf, err := revel.LoadConfig("app.conf")
	if err != nil {
		revel.INFO.Panicln("Email Load app.conf failed")
	}

	config = fmt.Sprintf(
		`{"username":"%s","password":"%s","host":"%s","port":%d,"from":"%s"}`,
		conf.StringDefault("mail.loginUsername", ""),
		conf.StringDefault("mail.loginPassword", ""),
		conf.StringDefault("mail.host", ""),
		conf.IntDefault("mail.port", 25),
		conf.StringDefault("mail.from", ""),
	)
	queueSize := conf.IntDefault("mail.queueSize", 20)
	revel.INFO.Printf(mate, queueSize)
	sendChannel = make(chan *Email, queueSize)

	go func() {
		for {
			select {
			case m, ok := <-sendChannel:
				if !ok {
					revel.INFO.Println("email send successful", m)
					return
				}
				if err := m.Send(); err != nil {
					revel.ERROR.Println("SendMail:", err.Error())
				}
			}
		}
	}()
}

func SendHtmlMail(receiverAddress, subject, content string) bool {
	mail := NewEMail(config)
	mail.To = []string{receiverAddress}
	mail.Subject = subject
	mail.HTML = content
	mail.Identity = ""

	select {
	case sendChannel <- mail:
		return true
	case <-time.After(time.Second * 3):
		return false
	}
}

func SendTextMail(receiverAddress, subject, content string) bool {
	mail := NewEMail(config)
	mail.To = []string{receiverAddress}
	mail.Subject = subject
	mail.Text = content
	mail.Identity = ""

	select {
	case sendChannel <- mail:
		return true
	case <-time.After(time.Second * 3):
		return false
	}
}
