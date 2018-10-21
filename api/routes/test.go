package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Test model
type Test struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type routeServer struct {
}

// Routes - exported test routes
func (r routeServer) TestRoutes(e Env) http.Handler {
	router := chi.NewRouter()

	router.Get("/{testID}", e.GetTest)
	router.Delete("/{testID}", e.DeleteTest)

	return router
}

// GetTest returns a test given an id
func (e Env) GetTest(w http.ResponseWriter, r *http.Request) {
	testID := chi.URLParam(r, "testID")
	test := Test{
		Slug:  testID,
		Title: "Teststs",
		Body:  "Test body",
	}
	render.JSON(w, r, test)
}

// DeleteTest deletes a test given an id
func (e Env) DeleteTest(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Delete test"
	render.JSON(w, r, response)
}
