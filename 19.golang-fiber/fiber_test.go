package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiber(t *testing.T) {
	newApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	err := newApp.Listen("localhost:4000")
	if err != nil {
		panic(err)
	}
}

var engine = mustache.New("./template", ".mustache")

var app = fiber.New(fiber.Config{
	ErrorHandler: ErrorHandler,
	Views:        engine,
})

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	ctx.Status(fiber.StatusInternalServerError)
	return ctx.SendString("Error : " + err.Error())
}

func TestRoutingHelloWorld(t *testing.T) {

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	request := httptest.NewRequest(http.MethodGet, "/hello", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
}

func TestCtx(t *testing.T) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString(name)
	})

	req := httptest.NewRequest(http.MethodGet, "/hello?name=isro", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, "isro", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		first := c.Get("firstname")   //header
		last := c.Cookies("lastname") //cookie
		return c.SendString("Hello " + first + " " + last)
	})

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	req.Header.Set("firstname", "isro")
	req.AddCookie(&http.Cookie{
		Name:  "lastname",
		Value: "sabanur",
	})

	res, err := app.Test(req)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Hello isro sabanur", string(bytes))
}

func TestRouteParams(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")

		return c.SendString("Get Order " + orderId + " From User " + userId)
	})

	req := httptest.NewRequest(http.MethodGet, "/users/1/orders/2", nil)

	res, err := app.Test(req)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Get Order 2 From User 1", string(bytes))
}

func TestRequestForm(t *testing.T) {
	app.Post("/", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Muhamad Isro Sabanur")
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := app.Test(req)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Hello Muhamad Isro Sabanur", string(bytes))
}

//go:embed source/contoh.txt
var contohFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			panic(err)
		}

		return c.SendString("File " + file.Filename)
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, _ := writer.CreateFormFile("file", "contoh.txt")
	file.Write(contohFile)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := app.Test(req)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "File contoh.txt", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()

		request := new(LoginRequest)
		err := json.Unmarshal(body, request)

		if err != nil {
			panic(err)
		}

		return c.SendString("Login Success " + request.Username)
	})

	body := strings.NewReader(`{
		"username" : "isro167",
		"password" : "isro"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/login", body)

	res, err := app.Test(req)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)

	assert.Equal(t, "Login Success isro167", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(ctx *fiber.Ctx) error {
		request := new(RegisterRequest)
		err := ctx.BodyParser(request)
		if err != nil {
			panic(err)
		}

		return ctx.SendString("Berhasil registrasi dengan nama " + request.Name)
	})
}

func TestBodyParserJSON(t *testing.T) {

	TestBodyParser(t)

	body := strings.NewReader(`{
		"username" : "isroo167",
		"name" : "Muhamad Isro Sabanur",
		"password" : "isroo167"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/register", body)

	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Berhasil registrasi dengan nama Muhamad Isro Sabanur", string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`username=isroo167&password=isro&name=Muhamad+Isro+Sabanur`)

	req := httptest.NewRequest(http.MethodPost, "/register", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := app.Test(req)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Berhasil registrasi dengan nama Muhamad Isro Sabanur", string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`
		<RegisterRequest>
			<username>isroo167</username>
			<password>isroo167</password>
			<name>Muhamad Isro Sabanur</name>
		</RegisterRequest>
	`)

	req := httptest.NewRequest(http.MethodPost, "/register", body)

	req.Header.Set("Content-Type", "application/xml")

	res, err := app.Test(req)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Berhasil registrasi dengan nama Muhamad Isro Sabanur", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"username": "isroo167",
			"name":     "Muhamad Isro Sabanur",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set("accept", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, `{"name":"Muhamad Isro Sabanur","username":"isroo167"}`, string(bytes))
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(ctx *fiber.Ctx) error {
		return ctx.Download("./source/contoh.txt", "contoh.txt")
	})

	req := httptest.NewRequest(http.MethodGet, "/download", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "attachment; filename=\"contoh.txt\"", res.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, "HELLO", string(bytes))
}

func TestRouteGroup(t *testing.T) {
	helloWorld := func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	}

	api := app.Group("/api")

	api.Get("/hello", helloWorld)
	api.Get("/world", helloWorld)

	web := app.Group("/web")

	web.Get("/hello", helloWorld)
	web.Get("/world", helloWorld)

	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, "Hello World", string(bytes))
}

func TestStatic(t *testing.T) {
	app.Static("/public", "./source")

	req := httptest.NewRequest(http.MethodGet, "/public/contoh.txt", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)

	assert.Equal(t, "HELLO", string(bytes))
}

func TestErrorHandling(t *testing.T) {
	app.Get("/error", func(*fiber.Ctx) error {
		return errors.New("ups")
	})

	req := httptest.NewRequest(http.MethodGet, "/error", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)

	assert.Equal(t, "Error : ups", string(bytes))
}

func TestTemplateEngine(t *testing.T) {
	app.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title":   "Hello Title",
			"header":  "Hello Header",
			"content": "Hello Content",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/view", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)

	assert.Contains(t, string(bytes), "<title>Hello Title</title>")
	assert.Contains(t, string(bytes), "<h1>Hello Header</h1>")
	assert.Contains(t, string(bytes), "<p>Hello Content</p>")
}

func TestMiddleware(t *testing.T) {
	app.Use("/api", func(ctx *fiber.Ctx) error {
		fmt.Println("I'm middleware before before request")
		err := ctx.Next()
		fmt.Println("I'm middleware before after request")

		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Api")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	bytes, err := io.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.Equal(t, "Hello Api", string(bytes))
}

func TestHTTPClient(t *testing.T) {
	
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	agent := client.Get("https://example.com/")
	status, response, err := agent.String()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Contains(t, response, "Example Domain")
}