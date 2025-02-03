package middleware

import (
	"net/http"
	"strings"
	"tokocikbosapi/model"
	"tokocikbosapi/utils/service"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization" binding:"required"`
}

func (amd *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var aH authHeader
		err := ctx.ShouldBindHeader(&aH)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		token := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", 1)
		tokenClaims, err := amd.jwtService.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Set("user", model.UserCredential{ID: tokenClaims.CredId, Roles: tokenClaims.Role})
		validRole := false
		for _, role := range roles {
			if role == tokenClaims.Role {
				validRole = true
				break
			}
		}
		if !validRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Access"})
			return
		}
		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
