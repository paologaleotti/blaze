package helloservice

import "net/http"

type HelloController struct {
	msg string
}

func NewHelloController(msg string) *HelloController {
	return &HelloController{
		msg: msg,
	}
}

func (hc *HelloController) GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(hc.msg))
}
