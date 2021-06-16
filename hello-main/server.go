package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	_ "os"
	"strconv"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes test
	e.GET("/", hello)
	e.POST("/login", login)
	e.GET("/user", User)
	//e.GET("/list", list)

	//Routes task
	e.POST("/register", createItem)
	e.GET("/item/:id", findByItem)
	e.GET("/list", listItems)
	e.PUT("item/:id", editItem)
	e.DELETE("/delete/:id", deleteItem)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type item struct {
	ID int "json:'id'"
	Name string "json:'name'"
	Status string "json:'status'"
}
var (
	items = map[int]*item{}
	count =1
)

// Add item
func createItem(c echo.Context)error  {
	u := &item{
		ID: count,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	items[u.ID] = u
	count++
	return c.JSON(http.StatusCreated, u)
}
// findByItem
func findByItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, items[id])
}

//List item
func listItems(c echo.Context) error {
	return c.JSON(http.StatusOK, items)
}
//delete item
func deleteItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(items, id)
	return c.NoContent(http.StatusNoContent)
}
//Edit item
func editItem(c echo.Context) error {
	u := new(item)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	items[id].Status = u.Status
	items[id].Name = u.Name
	return c.JSON(http.StatusOK, items[id])
}

func login(c echo.Context)error  {
	req := new(LoginRequest)
	c.Bind(req)
	log.Printf("req data %+v", req)
	if req.Username != "admin" || req.Password != "111"{
		return c.JSON(http.StatusUnauthorized, "Username/ Password invalid")
	}
	return c.JSON(http.StatusOK, &LoginResponse{
		Token: "done",
	})
}
type LoginResponse struct {
	Token string "json:'token'"
}
type LoginRequest struct {
	Username string "json:'username' form:'username' xml:'username' query:'username'"
	Password string "json:'password' form:'password' xml:'password' query:'password'"
}
func hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h4>Hello, World!</h4>")
}
func list(c echo.Context)error  {
	return c.HTML(http.StatusOK, "<h4>Home</h4>" +
		"<a href=\"user\">View</a>")
}

type Users struct {
	Name string "json:'Name' xml:'Name'"
	Email string "json:'Email' xml:'Email'"
}

func User(c echo.Context)error {
	u := &Users{
		Name:  "John",
		Email: "John@gmail",
	}
	return c.JSON(http.StatusOK, u)
}