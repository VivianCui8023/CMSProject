package controller

import (
	"context"
	"github.com/kataras/iris/v12/sessions"
	"webProject/service"
)

type OrderController struct {
	Session *sessions.Session
	Server  service.OrderDetailServer
	Ctx     context.Context
}
