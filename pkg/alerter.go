package pkg

import "fmt"

type Alerter interface {
    Send(msg string)
}

// DingtalkAlert 钉钉通知
type DingtalkAlert struct {
}

func NewDingtalkAlert() *DingtalkAlert {
    return &DingtalkAlert{}
}

func (alert *DingtalkAlert) Send(msg string) {
    fmt.Printf("钉钉通知: %v\n", msg)
}

// EmailAlert 邮件通知
type EmailAlert struct {
}

func NewEmailAlert() *EmailAlert {
    return &EmailAlert{}
}

func (alert *EmailAlert) Send(msg string) {
    fmt.Printf("邮件通知: %v\n", msg)
}
