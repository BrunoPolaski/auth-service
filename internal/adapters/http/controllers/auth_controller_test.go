package controller_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/BrunoPolaski/login-service/internal/controller"
// 	"github.com/BrunoPolaski/login-service/internal/tests"
// )

// func TestSignin(t *testing.T) {
// 	mockDb := database
// 	mockController := controller.NewAuthController()
// 	t.Setenv("ENV", "dev")
// 	t.Setenv("LOG_LEVEL", "info")
// 	t.Setenv("USERNAME", "test")
// 	t.Setenv("PASSWORD", "testpassword")
// 	t.Run("should return ok when credentials are correct", func(t *testing.T) {
// 		request := httptest.NewRequest(
// 			http.MethodGet,
// 			"/auth/signin",
// 			strings.NewReader("{}"),
// 		)
// 		request.SetBasicAuth("test", "testpassword")

// 		w := httptest.NewRecorder()

// 		mockController.SignIn(w, request)

// 		tests.AssertEqual(t, http.StatusOK, w.Code)
// 	})

// 	t.Run("should return unauthorized when there are no basic auth", func(t *testing.T) {
// 		request := httptest.NewRequest(
// 			http.MethodGet,
// 			"/auth/signin",
// 			strings.NewReader("{}"),
// 		)

// 		w := httptest.NewRecorder()

// 		mockController.SignIn(w, request)

// 		tests.AssertEqual(t, http.StatusUnauthorized, w.Code)
// 	})

// 	t.Run("should return unauthorized when credentials are incorrect", func(t *testing.T) {
// 		request := httptest.NewRequest(
// 			http.MethodGet,
// 			"/auth/signin",
// 			strings.NewReader("{}"),
// 		)
// 		request.SetBasicAuth("incorrect", "incorrectpassword")

// 		w := httptest.NewRecorder()

// 		mockController.SignIn(w, request)

// 		tests.AssertEqual(t, http.StatusUnauthorized, w.Code)
// 	})
// }
