package middleware

import (
	"net/http"

	"api/internal/auth"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserID = "userID"
	ContextUser   = "user"
)

func RequireAuth(JWT *auth.JWTManager, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(AccessTokenCookie)
		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "authentication required",
				},
			)

			return
		}

		claims, err := JWT.Parse(cookie)
		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			return
		}

		user, err := authService.GetUser(
			c.Request.Context(),
			claims.UserID,
		)
		if err != nil {

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "user not found",
				},
			)

			return
		}

		c.Set(ContextUserID, user.ID)
		c.Set(ContextUser, user)

		c.Next()
	}
}

func UserID(c *gin.Context) string {
	id, _ := c.Get(ContextUserID)

	userID, _ := id.(string)

	return userID
}

func User(c *gin.Context) *auth.User {
	value, ok := c.Get(ContextUser)

	if !ok {
		return nil
	}

	user, _ := value.(*auth.User)

	return user
}
