package pages

import "github.com/senforsce/tndr0cean/router"

func Index(c *router.Context) error {
	return c.Render(IndexPage(c))
}
