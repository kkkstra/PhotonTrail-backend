package jwt

import (
	"PhotonTrail-backend/internal/global"
	"PhotonTrail-backend/pkg/jwt"
	"time"
)

func VerifyJwtToken(claims *jwt.Claims) bool {
	if time.Now().Unix() > claims.ExpiresAt || claims.Issuer != global.Config.Jwt.Issuer {
		return false
	}
	// TODO: add more verification
	return true
}
