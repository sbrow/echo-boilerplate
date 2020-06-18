package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHomeController_index(t *testing.T) {
	want := struct {
		err    bool
		status int
		body   string
	}{
		false,
		http.StatusOK,
		"pong!",
	}

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	h := HomeController{}

	// Make sure error is nil
	if err := h.index(c); (err != nil) != want.err {
		t.Errorf("HomeController.index() error = %v, wantErr %v", err, want.err)
	}

	res := rec.Result()
	// Make sure code is correct
	if got := res.StatusCode; got != want.status {
		t.Errorf("HomeController.index() got = %v, want = %v", got, want.status)
	}
	// Make sure body is correct.
	if got := rec.Body.String(); got != want.body {
		t.Errorf(`HomeController.index() got = "%v", want = "%v"`, got, want.body)
	}
}
