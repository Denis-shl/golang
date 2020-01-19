package handlers

type (
	response = string
	request  = string
)

// Handler ...
type Handler interface {
	Response(request) response
}
