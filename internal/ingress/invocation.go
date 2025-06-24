package ingress

type Invocation struct {
	InvocationId string `json:"invocationId"`
	Status       string `json:"status"`
	Error        error  `json:"-"`
}
