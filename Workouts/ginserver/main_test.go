package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/:name",greet)


	req,_ := http.NewRequest("GET","/user/Fazal", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	expected := `{"message":"hello Fazal"}`
	if w.Body.String() != expected {
		t.Errorf("Expected %s, but got %s",expected,w.Body.String())
	}
}

func TestHealthRouter(t *testing.T)  {
	
}