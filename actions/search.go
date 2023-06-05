package actions

import (
	"github.com/gobuffalo/envy"
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
)

var (
	apiKey    = envy.Get("API_KEY", "faf7e5bb&s")
	urlMovies = envy.Get("URL", "http://www.omdbapi.com/")
)

// SearchHandler is a default handler to serve up
// a search API.
func SearchHandler(c buffalo.Context) error {
	var title = c.Param("s")
	if title == "" {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"error": "Please specify the title of the movie!"}))
	}

	var (
		page         = c.Param("page")
		pageInt, err = strconv.Atoi(page)
	)
	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"error": "Invalid page number!"}))
	}

	if pageInt < 1 || pageInt > 100 {
		page = "1"
	}

	resp, err := client.R().
		SetQueryString("apikey="+apiKey).
		SetQueryParam("page", page).
		SetQueryParam("s", title).
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

func renderError(c buffalo.Context, err error) error {
	return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
}
