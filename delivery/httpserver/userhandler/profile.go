package userhandler

import (
	"fmt"
	"gameapp/entity/auth"
	entity "gameapp/entity/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) userProfileHandler(c echo.Context) error {

	//token := c.Request().Header.Get("Authorization")
	//
	//token = strings.Replace(token, "Bearer ", "", 1)

	claims := c.Get("claims").(auth.Claims)

	fmt.Println("claims", claims)
	//claims, err := h.authService.Parse(token)

	profileReq := entity.ProfileRequest{
		UserId: claims.UserID,
	}

	response, err := h.userService.Profile(profileReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, response)

}
