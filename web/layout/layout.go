package layout

import (
	"github.com/senforsce/tndr0cean/router"
)

func Handler(c *router.Context) error {
	return c.Render(Base(c))
}

func Robots(c *router.Context) error {
	return c.Text(200, "User-agent: * \n Disallow: /")
}
