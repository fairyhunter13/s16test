package actions

import (
	"net/http"
	"testing"
)

func (as *ActionSuite) Test_SearchHandler() {
	as.T().Run("returns 422 if no title is specified", func(t *testing.T) {
		res := as.JSON("/search").Get()

		as.Equal(http.StatusUnprocessableEntity, res.Code)
		as.Contains(res.Body.String(), "Please specify the title")
	})
	as.T().Run("returns 422 if page is invalid", func(t *testing.T) {
		res := as.JSON("/search?s=hello&page=test").Get()

		as.Equal(http.StatusUnprocessableEntity, res.Code)
		as.Contains(res.Body.String(), "Invalid page number")
	})
	as.T().Run("returns 200 if all params are valid", func(t *testing.T) {
		res := as.JSON("/search?s=terminator&page=1").Get()

		as.Equal(http.StatusOK, res.Code)
		as.Contains(res.Body.String(), "Terminator")
	})
}
