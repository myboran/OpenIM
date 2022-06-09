package token_verify

import (
	"open-im/pkg/common/config"
	"open-im/pkg/common/constant"
	commonDB "open-im/pkg/common/db"
	"open-im/pkg/common/log"
	"open-im/pkg/utils"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UID      string
	Platform string //login platform
	jwt.RegisteredClaims
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.TokenPolicy.AccessSecret), nil
	}
}

func GetClaimFromToken(tokensString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokensString, &Claims{}, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, &constant.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, &constant.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, &constant.ErrTokenNotValidYet
			} else {
				return nil, &constant.ErrTokenUnknown
			}
		} else {
			return nil, &constant.ErrTokenNotValidYet
		}
	} else {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			log.NewDebug("", claims.UID, claims.Platform)
			return claims, nil
		}
		return nil, &constant.ErrTokenNotValidYet
	}
}

func GetUserIDFromToken(token string, operationID string) (bool, string) {
	claims, err := ParseToken(token)
	if err != nil {
		log.Error(operationID, "ParseToken failed, ", err.Error(), token)
		return false, ""
	}
	return true, claims.UID
}

func ParseToken(tokensString string) (claims *Claims, err error) {

	claims, err = GetClaimFromToken(tokensString)
	if err != nil {
		log.NewError("", "token validate err", err.Error())
		return nil, err
	}

	m, err := commonDB.DB.GetTokenMapByUidPid(claims.UID, claims.Platform)
	if err != nil {
		log.NewError("", "get token from redis err", err.Error())
		return nil, &constant.ErrTokenInvalid
	}
	if m == nil {
		log.NewError("", "get token from redis err", "m is nil")
		return nil, &constant.ErrTokenInvalid
	}
	if v, ok := m[tokensString]; ok {
		switch v {
		case constant.NormalToken:
			log.NewDebug("", "this is normal return", claims)
			return claims, nil
		case constant.InValidToken:
			return nil, &constant.ErrTokenInvalid
		case constant.KickedToken:
			return nil, &constant.ErrTokenKicked
		case constant.ExpiredToken:
			return nil, &constant.ErrTokenExpired
		default:
			return nil, &constant.ErrTokenUnknown
		}
	}
	return nil, &constant.ErrTokenUnknown
}

func CheckAccess(OpUserID string, OwnerUserID string) bool {
	if utils.IsContain(OpUserID, config.Config.Manager.AppManagerUid) {
		return true
	}
	if OpUserID == OwnerUserID {
		return true
	}
	return false
}
