package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	fmt.Println("init")
	res := m.Run()
	fmt.Println("done");
	os.Exit(res)
}

func WithGin(t *testing.T, fn func(e *httpexpect.Expect, r *httpexpect.AssertReporter)) {
	gin.SetMode(gin.TestMode)
	// Create new gin instance
	engine := gin.New()
	// Add /example route via handler function to the gin instance
	handler := GinHandler(engine)
	r := httpexpect.NewAssertReporter(t)
	// Create httpexpect instance
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: r,
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	fn(e, r)
}

func TestUserList(t *testing.T) {
	WithGin(t, func(e *httpexpect.Expect, r *httpexpect.AssertReporter) {
		res := e.GET("/api/v1/users").Expect().Status(http.StatusOK).JSON().Array()
		res.Length().Equal(3)
		res.ContainsOnly(
			map[string]interface{} {"id":"1", "name":"user_1"},
			map[string]interface{} {"id":"2", "name":"user_2"},
			map[string]interface{} {"id":"3", "name":"user_3"},
		)
	})
}