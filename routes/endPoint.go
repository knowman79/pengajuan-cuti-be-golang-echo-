package routes

import (
	"example/service"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/v4/middleware"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

//RoleEndPoint function
func Endpoint() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Allowing request from certain origin
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://localhost:8080", "http://51.79.185.235:9094", "http://localhost:8080", "https://matrix-3.cloudias79.com:9094/", "https://localhost:8081", "http://localhost:8081", "https://matrix-3.cloudias79.com", "https://localhost"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//login endpoint
	e.GET("/login", service.Login)

	//roles endpoint
	e.GET("/role/readAll", service.ReadAllRole)

	//user endpoint
	e.POST("/user/save", service.CreateUser)
	e.GET("/user/readAll", service.ReadAllUser)
	e.PUT("/user/update", service.UpdateUser)
	e.DELETE("/user/delete", service.DeleteUser)
	e.GET("/user/list", service.ReadAllOlUser)

	//leave_allowance endpoint
	e.POST("/allowance/save", service.CreateAllowance)
	e.GET("/allowance/readAll", service.ReadAllAllowance)
	e.PUT("/allowance/update", service.UpdateAllowance)
	e.DELETE("/allowance/delete", service.DeleteAllowance)

	//leave endpoint
	e.GET("/leave/readAll", service.ReadAllLeave)
	e.POST("/leave/save", service.CreateLeave)
	e.DELETE("/leave/delete", service.DeleteLeave)
	e.PUT("/leave/update", service.UpdateLeave)
	e.GET("/leave/readById", service.ReadIdLeave)
	e.DELETE("/leave/deleteDraft", service.DeleteLeaveDraft)
	e.PUT("/leave/updateApproved", service.UpdateLeaveApproved)
	e.PUT("/leave/updateToInprogress", service.UpdateLeaveOpenToInprogress)
	e.GET("/leave/readLeaveByName", service.ReadLeaveByName)
	e.PUT("/leave/updateToOpen", service.UpdateLeaveDraftToOpen)
	e.PUT("/leave/updateToCanceled", service.UpdateLeaveCanceled)
	e.PUT("/leave/UpdateRejectBySPV", service.UpdateRejectBySPV)
	e.PUT("/leave/updateToRejectByHRD", service.UpdateLeaveRejectByHRD)
	e.PUT("/leave/updateIfStatusDraft", service.UpdateStatusDraft)

	e.Logger.Fatal(e.Start(":1323"))
}
