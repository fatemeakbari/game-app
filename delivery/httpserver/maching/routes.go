package matchinghandler

import "github.com/labstack/echo/v4"

func (h Handler) Route(e *echo.Echo) {

	group := e.Group("/player-match", h.AuthorizeToken())

	group.POST("/", h.addUserToWaitingList)

}
