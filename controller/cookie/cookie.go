package cookie

import (
	"go-cookie/middle"
	"go-cookie/module/cookie"

	"github.com/davecgh/go-spew/spew"
)

var testKey = "test"

// SetCookie -
func SetCookie(c *middle.C) {
	m := map[string]string{
		"text": "test",
		"name": "YauTz",
	}

	cke, err := cookie.SetCookie(testKey, m)
	if err != nil {
		c.E(err)
		return
	}

	c.SetCookie(cke)
	c.R(nil)
}

// ReadCookie -
func ReadCookie(c *middle.C) {
	cke, err := c.Request().Cookie(testKey)
	if err != nil {
		c.E(err)
		return
	}

	m := make(map[string]string)
	if err := cookie.ReadCookie(cke, testKey, m); err != nil {
		spew.Dump(err)
		c.E(err)
		return
	}

	spew.Dump(m)
	c.R(nil)
}
