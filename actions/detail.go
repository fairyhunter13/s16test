package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
)

// DetailHandler is a default handler to serve up
// a movie detail.
func DetailHandler(c buffalo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"error": "Please specify the id of a movie!"}))
	}

	var plot = c.Param("plot")
	if plot == "" {
		plot = "short"
	}

	if plot != "short" && plot != "full" {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"error": "Invalid value of plot, valid value is either short or full!"}))
	}

	resp, err := client.R().
		SetQueryString("apikey="+apiKey).
		SetQueryParam("i", id).
		SetQueryParam("plot", plot).
		Get(urlMovies)
	if err != nil {
		return renderError(c, err)
	}

	bytes, err := resp.ToBytes()
	if err != nil {
		return renderError(c, err)
	}

	return c.Render(resp.StatusCode, JSONBytes(bytes))
}
