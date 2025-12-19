package auth

import (
	"iris/application/auth"
	authTypes "iris/domain/types/auth"
	"iris/routes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Configure(g *gin.RouterGroup, authSvc *auth.Service) {
	g.POST("auth/login", func(c *gin.Context) {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var loginRequest LoginRequest
		if err := c.ShouldBind(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		token, err := authSvc.Login(loginRequest.Username, loginRequest.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, routes.SuccessResponse(token))
	})

	g.POST("auth/logout", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		err = authSvc.Logout(ctx.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, routes.SuccessResponse(nil))
	})

	g.POST("auth/issue_user", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		type IssueUserRequest struct {
			Username  string              `json:"username"`
			Password  string              `json:"password"`
			Clearance authTypes.Clearance `json:"clearance"`
		}
		var issueUserRequest IssueUserRequest
		if err := c.ShouldBind(&issueUserRequest); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		user, err := authSvc.IssueNewUser(
			ctx,
			issueUserRequest.Username,
			issueUserRequest.Password,
			issueUserRequest.Clearance,
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, routes.SuccessResponse(user))
	})

	g.POST("auth/reset_password", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}

		type ResetPasswordRequest struct {
			UserID      int    `json:"user_id"`
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}
		var resetPasswordRequest ResetPasswordRequest
		if err := c.ShouldBind(&resetPasswordRequest); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		err = authSvc.ResetPassword(ctx, resetPasswordRequest.UserID, resetPasswordRequest.OldPassword, resetPasswordRequest.NewPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, routes.SuccessResponse(nil))
	})

	g.GET("auth/hi", routes.RequireAuth(authSvc), func(c *gin.Context) {
		_, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, routes.SuccessResponse("hey"))
	})

	g.GET("user", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		if idStr, ok := c.GetQuery("id"); ok {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
				return
			}

			user, err := authSvc.GetUserByID(ctx, id)
			if err != nil {
				c.JSON(http.StatusNotFound, routes.ErrorResponse(err))
				return
			}

			c.JSON(http.StatusOK, routes.SuccessResponse(user))
			return
		}

		if username, ok := c.GetQuery("username"); ok {
			user, err := authSvc.GetUserByUsername(ctx, username)
			if err != nil {
				c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
				return
			}

			c.JSON(http.StatusOK, routes.SuccessResponse(user))
			return
		}
		users, err := authSvc.GetUsers(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, routes.SuccessResponse(users))
	})

	g.POST("user/can", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, routes.ErrorResponse(err))
			return
		}
		type PermissionRequest struct {
			UserID  int               `json:"userID"`
			Action  authTypes.Action  `json:"action"`
			Content authTypes.Content `json:"content"`
			IsOwner bool              `json:"isOwner"`
		}
		var permissionRequest PermissionRequest
		if err := c.ShouldBind(&permissionRequest); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		c.JSON(
			routes.SmartResponse(
				authSvc.UserCan(
					ctx,
					permissionRequest.UserID,
					permissionRequest.Action,
					permissionRequest.Content,
					permissionRequest.IsOwner,
				),
			),
		)
	})
}
