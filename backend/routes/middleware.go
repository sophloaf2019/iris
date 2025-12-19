package routes

import (
	"errors"
	"iris/application/auth"
	authTypes "iris/domain/types/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authContextKey = "authContext"

func RequireAuth(authSvc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		parts := strings.SplitN(h, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			return
		}

		token := parts[1]
		ctx, err := authSvc.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// stash auth context for later handlers
		c.Set(authContextKey, ctx)

		c.Next()
	}
}

func AuthContext(c *gin.Context) (authTypes.Context, error) {
	v, ok := c.Get(authContextKey)
	if !ok {
		return authTypes.Context{}, errors.New("auth context missing")
	}

	ctx, ok := v.(authTypes.Context)
	if !ok {
		return authTypes.Context{}, errors.New("auth context has wrong type")
	}

	return ctx, nil
}
