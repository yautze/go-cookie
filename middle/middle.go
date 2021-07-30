package middle

import (
	"github.com/kataras/iris/v12"
	"github.com/tyr-tech-team/hawk/response"
)

// C -
type C struct {
	iris.Context
}

// HandleFunc -
func HandleFunc(handler func(*C)) func(iris.Context) {
	return func(c iris.Context) {
		customerContext := &C{
			c,
		}

		handler(customerContext)
	}
}

// R -
func (c *C) R(data interface{}) {
	c.StatusCode(iris.StatusOK)
	c.JSON(response.Resp(c.Request().Context(), data))
}

// E -
func (c *C) E(err error) {
	c.StatusCode(iris.StatusOK)
	c.JSON(response.Error(c.Request().Context(), err))
}
