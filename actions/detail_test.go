package actions

import (
	"net/http"
	"testing"
)

func (as *ActionSuite) Test_DetailHandler() {
	as.T().Run("returns 404 if no id is specified", func(t *testing.T) {
		res := as.JSON("/detail").Get()

		as.Equal(http.StatusNotFound, res.Code)
	})
	as.T().Run("returns 422 if plot is invalid", func(t *testing.T) {
		res := as.JSON("/detail/tt0088247?plot=test").Get()

		as.Equal(http.StatusUnprocessableEntity, res.Code)
		as.Contains(res.Body.String(), "Invalid value of plot")
	})
	as.T().Run("returns 200 if all params are valid", func(t *testing.T) {
		res := as.JSON("/detail/tt0088247?plot=full").Get()
		as.Equal(http.StatusOK, res.Code)
		as.Contains(res.Body.String(), "Terminator")
	})
}
