package matchinghandler

import (
	"gameapp/entity/auth"
	matchingentity "gameapp/entity/matching"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) addUserToWaitingList(c echo.Context) error {

	var req matchingentity.AddUserToWaitingListRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request format is wrong")
	}

	claims := c.Get("claims").(auth.Claims)
	req.UserId = claims.UserID

	res, err := h.service.AddUserToWaitingList(req)

	//TODO call later
	//h.service.MatchWaitingPlayer(context.Background(), req.Category)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}
