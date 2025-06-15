package tests

import (
	"fmt"
	"intern/models"
	"intern/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type loginTest struct {
}

func Ping(t *testing.T, ge *gin.Engine) {
	ge.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	ge.ServeHTTP(w, req)

	// 6. Make assertions
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestLoginroute(t *testing.T) {

	ge := gin.New()
	Ping(t, ge)
	service := loginTest{}
	ge.POST("/login", routes.Login(&service))
	reader := strings.NewReader(`
	{
		"email":"aryan@gmail.com",
		"password":"aryan001"
	
	}`)
	req, err := http.NewRequest(http.MethodPost, "/login", reader)

	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	ge.ServeHTTP(w, req)

	fmt.Println(w.Body)
	// 6. Make assertions
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

}
func TestSignUproute(t *testing.T) {

	ge := gin.New()
	Ping(t, ge)
	service := loginTest{}
	ge.POST("/signup", routes.Login(&service))
	reader := strings.NewReader(`
	{
		"name":"asd",
		"role":"Doctor",
		"email":"asda@gmail.com",
		"password":"hello"
	}
	`)
	req, err := http.NewRequest(http.MethodPost, "/signup", reader)

	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	ge.ServeHTTP(w, req)

	fmt.Println(w.Body)
	// 6. Make assertions
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

}

func (lt *loginTest) FindUser(credentials models.LoginJson) (*models.User, error) {
	user := models.User{
		Name: "hello",
	}
	return &user, nil
}
func (lt *loginTest) AddUser(user *models.User) error {
	return nil
}
