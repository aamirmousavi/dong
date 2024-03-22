package context

import "github.com/aamirmousavi/dong/internal/database/mongodb"

type Context struct {
	mongodb *mongodb.Handler
}

func NewContext(
	mongodb *mongodb.Handler,
) *Context {
	return &Context{
		mongodb,
	}
}

func (ctx *Context) Mongo() *mongodb.Handler {
	return ctx.mongodb
}
