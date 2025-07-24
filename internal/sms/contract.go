package sms

import "example.com/local/Go2part/domain"

type ISMS interface {
	SmsSend(p *domain.Sms) (Response, error)
	SmsStatus(id string) (Response, error)
	SmsCost(p *domain.Sms) (Response, error)
	MyBalance() (Response, error)
}
