package hander

type Handler interface {
	Authorize()
	Token()
	Registration()
}
