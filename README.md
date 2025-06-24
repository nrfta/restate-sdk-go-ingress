restate-sdk-go-ingress
======================

Provides a Go client for the restate ingress API.

# Installation

Add to your Go project with:

```shell
go get github.com/nrfta/restate-sdk-go-ingress
```

# Usage

## IngressService

```go
var input MyInputStruct
output, err := restateIngress.IngressService[*MyInputStruct, *MyOutputStruct](
	serviceName, handlerName, restateIngress.WithBaseUrl("http://localhost:8080")).
	Request(ctx, &input, 
		restateIngress.WithIdempotencyKey("idem-key-1"),
		restateIngress.WithHeaders(map[string]string{"header-name","header-value"}))
```

## IngressServiceSend

```go
var input MyInputStruct
invocation := restateIngress.IngressServiceSend[*MyInputStruct](
	serviceName, handlerName, restateIngress.WithBaseUrl("http://localhost:8080")).
	Send(ctx, &input, 
		restateIngress.WithIdempotencyKey("idem-key-1"),
		restateIngress.WithDelay(time.Minute),
		restateIngress.WithHeaders(map[string]string{"header-name","header-value"}))
if invocation.Error != nil {
	println(invocation.Error.Error())
}
println(invocation.Id)
```

## IngressObject

```go
var input MyInputStruct
output, err := restateIngress.IngressObject[*MyInputStruct, *MyOutputStruct](
	serviceName, objectKey, handlerName, restateIngress.WithBaseUrl("http://localhost:8080")).
	Request(ctx, &input, 
		restateIngress.WithIdempotencyKey("idem-key-1"),
		restateIngress.WithHeaders(map[string]string{"header-name","header-value"}))
```

## IngressObjectSend

```go
var input MyInputStruct
invocation := restateIngress.IngressObjectSend[*MyInputStruct](
	serviceName, objectKey, handlerName, restateIngress.WithBaseUrl("http://localhost:8080")).
	Send(ctx, &input, 
		restateIngress.WithIdempotencyKey("idem-key-1"),
		restateIngress.WithDelay(time.Minute),
		restateIngress.WithHeaders(map[string]string{"header-name","header-value"}))
if invocation.Error != nil {
	println(invocation.Error.Error())
}
println(invocation.Id)
```

## IngressWorkflow

```go
var input MyInputStruct
output, err := restateIngress.IngressWorkflow[*MyInputStruct, *MyOutputStruct](
	serviceName, workflowId, "send", restateIngress.WithBaseUrl("http://localhost:8080")).
	Request(ctx, &input, 
		restateIngress.WithIdempotencyKey("idem-key-1"),
		restateIngress.WithHeaders(map[string]string{"header-name","header-value"}))
```

## IngressWorkflowSend

```go
var input MyInputStruct
invocation := restateIngress.IngressWorkflowSend[*MyInputStruct](
	serviceName, workflowId, handlerName, restateIngress.WithBaseUrl("http://localhost:8080")).
	Send(ctx, &input, 
		restateIngress.WithIdempotencyKey("idem-key-1"),
		restateIngress.WithDelay(time.Minute),
		restateIngress.WithHeaders(map[string]string{"header-name","header-value"}))
if invocation.Error != nil {
	println(invocation.Error.Error())
}
println(invocation.Id)
```

## IngressAttachInvocation

**Blocking until invocation returns output**
```go
output, err := restateIngress.IngressAttachInvocation[*MyOutputStruct](
	invocationId, restateIngress.WithBaseUrl("http://localhost:8080")).
	Attach(ctx)
```

**Non-blocking, returns error if invocation is not found or not done**
```go
output, err := restateIngress.IngressAttachInvocation[*MyOutputStruct](
	invocationId, restateIngress.WithBaseUrl("http://localhost:8080")).
	Output(ctx)
```

**Cancel and invocation. NOTE: use the admin base URL here.**
```go
err := restateIngress.IngressAttachInvocation[any](
	invocationId, restateIngress.WithBaseUrl("http://localhost:9080")).
	Cancel(ctx)
```

TODO: Add options for Cancel, Kill, or Purge

## IngressAttachService

**Blocking until service invocation returns output**
```go
output, err := restateIngress.IngressAttachService[*MyOutputStruct](
	serviceName, handlerName, idempotencyKey, restateIngress.WithBaseUrl("http://localhost:8080")).
	Attach(ctx)
```

**Non-blocking, returns error if service invocation is not found or not done**
```go
output, err := restateIngress.IngressAttachService[*MyOutputStruct](
	serviceName, handlerName, idempotencyKey, restateIngress.WithBaseUrl("http://localhost:8080")).
	Output(ctx)
```

## IngressAttachObject

**Blocking until object invocation returns output**
```go
output, err := restateIngress.IngressAttachObject[*MyOutputStruct](
	serviceName, objectKey, handlerName, idempotencyKey, restateIngress.WithBaseUrl("http://localhost:8080")).
	Attach(ctx)
```

**Non-blocking, returns error if object invocation is not found or not done**
```go
output, err := restateIngress.IngressAttachObject[*MyOutputStruct](
	serviceName, objectKey, handlerName, idempotencyKey, restateIngress.WithBaseUrl("http://localhost:8080")).
	Output(ctx)
```

## IngressAttachWorkflow

**Blocking until workflow returns output**
```go
output, err := restateIngress.IngressAttachWorkflow[*MyOutputStruct](
	serviceName, workflowId, restateIngress.WithBaseUrl("http://localhost:8080")).
	Attach(ctx)
```

**Non-blocking, returns error if workflow is not found or not done**
```go
output, err := restateIngress.IngressAttachWorkflow[*MyOutputStruct](
	serviceName, workflowId, restateIngress.WithBaseUrl("http://localhost:8080")).
	Output(ctx)
```
