package options

import (
	"time"
)

// BEGIN NEW OPTIONS

type IngressOptions struct {
	BaseUrl string
}

type IngressOption interface {
	BeforeIngress(*IngressOptions)
}

// END NEW OPTIONS

type RequestOptions struct {
	IdempotencyKey string
	Headers        map[string]string
}

type RequestOption interface {
	BeforeRequest(*RequestOptions)
}

type SendOptions struct {
	IdempotencyKey string
	Headers        map[string]string
	Delay          time.Duration
}

type SendOption interface {
	BeforeSend(*SendOptions)
}
