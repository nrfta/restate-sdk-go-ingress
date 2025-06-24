package restateIngress

import (
	"github.com/nrfta/restate-sdk-go-ingress/internal/options"
	"time"
)

// BEGIN NEW OPTIONS

func WithBaseUrl(baseUrl string) withBaseUrl {
	return withBaseUrl{baseUrl}
}

type withBaseUrl struct {
	baseUrl string
}

func (w withBaseUrl) BeforeIngress(opts *options.IngressOptions) {
	opts.BaseUrl = w.baseUrl
}

// END NEW OPTIONS

type withHeaders struct {
	headers map[string]string
}

var _ options.RequestOption = withHeaders{}
var _ options.SendOption = withHeaders{}

func (w withHeaders) BeforeRequest(opts *options.RequestOptions) {
	opts.Headers = w.headers
}

func (w withHeaders) BeforeSend(opts *options.SendOptions) {
	opts.Headers = w.headers
}

// WithHeaders is an option to specify outgoing headers when making a call
func WithHeaders(headers map[string]string) withHeaders {
	return withHeaders{headers}
}

type withIdempotencyKey struct {
	idempotencyKey string
}

var _ options.RequestOption = withIdempotencyKey{}
var _ options.SendOption = withIdempotencyKey{}

func (w withIdempotencyKey) BeforeRequest(opts *options.RequestOptions) {
	opts.IdempotencyKey = w.idempotencyKey
}

func (w withIdempotencyKey) BeforeSend(opts *options.SendOptions) {
	opts.IdempotencyKey = w.idempotencyKey
}

// WithIdempotencyKey is an option to specify the idempotency key to set when making a call
func WithIdempotencyKey(idempotencyKey string) withIdempotencyKey {
	return withIdempotencyKey{idempotencyKey}
}

type withDelay struct {
	delay time.Duration
}

var _ options.SendOption = withDelay{}

func (w withDelay) BeforeSend(opts *options.SendOptions) {
	opts.Delay = w.delay
}

// WithDelay is an [SendOption] to specify the duration to delay the request
func WithDelay(delay time.Duration) withDelay {
	return withDelay{delay}
}
