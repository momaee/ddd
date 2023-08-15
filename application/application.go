package application

import (
	"ddd/integration/salesforce"
	"net/http"
	"strconv"

	repo "ddd/domain/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Application is a struct for application
type Application struct {
	e     *echo.Echo
	repos repo.Repositories
}

// New is a constructor for Application
func New() *Application {
	// build vendors
	crm := salesforce.New()

	// build integration by vendors
	integration := repo.NewIntegration(crm, nil)

	// build repositories by integration
	repos := *repo.NewRepositories(integration)

	// build application by repositories
	return &Application{
		e:     echo.New(),
		repos: repos,
	}
}

// Start is a method for starting application
func (a *Application) Start() {
	// Middleware
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())

	// Routes
	a.e.GET("/", a.hello)
	a.e.GET("/customers/:id", a.getCustomerByID)

	// Start server
	a.e.Logger.Fatal(a.e.Start(":1323"))
}

// ------------------------------ < Handlers > ------------------------------

func (a *Application) hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (a *Application) getCustomerByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	customer, err := a.repos.Customer.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, customer)
}
