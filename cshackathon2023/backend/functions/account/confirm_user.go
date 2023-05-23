package account

import (
	"backend/modules"
	idb "backend/modules/db"
	"backend/types/common"
	"backend/types/model"
	"github.com/golang-jwt/jwt/v4"
)

func ConfirmUser(user *model.User) (*string, error) {
	if result := modules.DB.FirstOrCreate(&user, idb.Where(model.UserFieldFirebaseUid, "= ?"), user.FirebaseUid); result.Error != nil {
		return nil, result.Error
	}

	// Create JWT claims
	claims := &common.UserClaims{
		UserId: user.Id,
		Name:   user.Name,
	}

	// Sign JWT token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, err := jwtToken.SignedString([]byte(modules.Conf.JwtSecret))
	if err != nil {
		return nil, err
	}

	return &signedJwtToken, nil
}
