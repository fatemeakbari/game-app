package middleware

import (
	"gameapp/entity/auth"
	model "gameapp/model/accesscontrol"
	"gameapp/model/usermodel"
	"gameapp/pkg/errorhandler/errormessage"
	service "gameapp/service/accesscontrol"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ACLMiddleware struct {
	service service.Service
}

func NewACLMiddleware(srv service.Service) ACLMiddleware {

	return ACLMiddleware{
		service: srv,
	}
}

func (acl ACLMiddleware) IsUserHasAccess(permissions ...model.PermissionTitle) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			claims := c.Get("claims").(auth.Claims)

			res, err := acl.service.IsUserHasAccess(claims.UserID, uint(usermodel.MapStrToRole(claims.Role)), permissions)

			if err != nil || !res {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": errormessage.ForbiddenMessage,
				})
			}

			return next(c)
		}
	}
}
