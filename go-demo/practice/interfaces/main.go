package main

import (
	"fmt"
)

type Notifier interface {
	Message(m string) string
}

type Email struct {
	email_addres string
}

func (e Email) Message(m string) string {
	return fmt.Sprintf("Your message: `%s` has been sent to email %s", m, e.email_addres)
}

type SMS struct {
	phone_number string
}

func (s SMS) Message(m string) string {
	return fmt.Sprintf("Your message: `%s` has been sent to number %s", m, s.phone_number)
}

type Push struct {
	device_token string
}

func (p Push) Message(m string) string {
	return fmt.Sprintf("Your message: `%s` has been sent to specific destination %s", m, p.device_token)
}

type MultiNotifier struct {
	notifiers []Notifier
}

func (mn MultiNotifier) Message(m string) string {
	for _, n := range mn.notifiers {
		fmt.Println(n.Message(m))
	}
	return ""
}

func SendNotification(n Notifier, m string) {
	fmt.Println(n.Message(m))
}

func main() {
	multi := MultiNotifier{notifiers: []Notifier{
		Email{"something@gmail.com"},
		SMS{"+777777777"},
		Push{"1absd43vas"},
	}}

	SendNotification(multi, "Hello")
}
