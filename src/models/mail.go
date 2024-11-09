package models

import (
	"strings"
)

type MailSend struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    []string `json:"body"`
}

func (m *MailSend) Sender(sender string) *MailSend {
	m.From = sender
	return m
}
func (m *MailSend) Receiver(recipient string) *MailSend {
	m.To = append(m.To, recipient)
	return m
}
func (m *MailSend) Title(subject string) *MailSend {
	m.Subject = subject
	return m
}
func (m *MailSend) MsgLine(line string) *MailSend {
	m.Body = append(m.Body, line)
	return m
}
func (m *MailSend) MsgSameLine(values ...interface{}) *MailSend {
	var line []string
	for _, l := range values {
		line = append(line, l.(string))
	}

	m.Body = append(m.Body, strings.Join(line, " "))
	return m
}
func (m *MailSend) MsgWhiteLine() *MailSend {
	m.Body = append(m.Body, "")
	return m
}

func (m *MailSend) Message() []byte {
	return []byte("To: " + strings.Join(m.To, ", ") + "\r\n" +
		"From: " + m.From + "\r\n" +
		"Subject: " + m.Subject + "\r\n" +
		"\r\n" +
		strings.Join(m.Body, "\r\n"),
	)
}
