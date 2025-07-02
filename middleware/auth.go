package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"starzeng.com/gin-demo/common"
	"starzeng.com/gin-demo/model"
	"starzeng.com/gin-demo/utils"
	"strings"
	"time"
)

var JwtKey = []byte("my_super_secret_key")

const TokenExpireDuration = time.Minute * 5
const RefreshThreshold = time.Minute * 2

type MyClaims struct {
	Username   string   `json:"username"`
	Role       string   `json:"role"`
	Permission []string `json:"permission"`
	jwt.RegisteredClaims
}

func GenerateToken(u model.User) (string, string, error) {
	jti := uuid.NewString()
	now := time.Now()
	claims := MyClaims{
		u.Username,
		u.Role,
		u.Permission,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ID:        jti,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(JwtKey)
	return signed, jti, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		claims := &MyClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil || !token.Valid {
			common.Error(c, common.CodeUnauthorized, "token 无效或过期")
			c.Abort()
			return
		}
		if utils.BlacklistCheck(claims.ID) {
			common.Error(c, common.CodeUnauthorized, "token 已注销")
			c.Abort()
			return
		}
		if time.Until(claims.ExpiresAt.Time) < RefreshThreshold {
			u := model.Users[claims.Username]
			newToken, _, _ := GenerateToken(u)
			c.Header("X-Refresh-Token", newToken)
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("permission", claims.Permission)
		c.Set("jti", claims.ID)
		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != role {
			common.Error(c, common.CodeForbidden, "角色权限不足")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequirePermission(p string) gin.HandlerFunc {
	return func(c *gin.Context) {
		perms, _ := c.Get("permission")
		for _, perm := range perms.([]string) {
			if perm == p {
				c.Next()
				return
			}
		}
		common.Error(c, common.CodeForbidden, "操作权限不足")
		c.Abort()
	}
}
