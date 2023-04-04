package login

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/functions"
	"backend/modules/db"
	"backend/modules/db/model"
	"backend/modules/hub"
	"backend/types/common"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/crypto"
	"backend/utils/value"
)

func HandleCallback(c *fiber.Ctx) error {
	// * Parse body
	var body *payload.LoginSpotifyCallback
	if err := c.BodyParser(&body); err != nil {
		return response.Error(c, true, "Unable to parse body", err)
	}

	// * Check state
	var state *model.ClientState
	if result := db.DB.Preload("Client").First(&state, "state = ?", body.State); result.Error != nil {
		return response.Error(c, true, "Unable to find OAuth state", result.Error)
	}

	// * Check code
	credentials, err := functions.SpotifyGetRefreshToken(c, state.Client, "authorization_code", body.Code)
	if err != nil {
		return err
	}

	// * Check user exists
	profile, err := functions.SpotifyGetProfile(c, *credentials.AccessToken)
	if err != nil {
		return err
	}

	var user *model.User
	if result := db.DB.Preload("Client").First(&user, "spotify_id = ?", *profile.Id); result.Error == gorm.ErrRecordNotFound {
		// * Create user
		user = &model.User{
			Id:           nil,
			ClientId:     state.Client.Id,
			Client:       state.Client,
			SpotifyId:    profile.Id,
			Profile:      profile,
			RefreshToken: credentials.RefreshToken,
			CreatedAt:    nil,
			UpdatedAt:    nil,
		}
		if result := db.DB.Create(user); result.Error != nil {
			return response.Error(c, true, "Unable to create user", result.Error)
		}
	} else if result.Error != nil {
		return response.Error(c, true, "Unable to query user", result.Error)
	} else {
		user.Profile = profile
		user.RefreshToken = credentials.RefreshToken
		if result := db.DB.Save(user); result.Error != nil {
			return response.Error(c, true, "Unable to update user", result.Error)
		}
	}

	// * Register branch
	branchDsn, branchDb := hub.Create(user)
	if branchDb == nil {
		return response.Error(c, true, "Unable to create database connection")
	}
	hub.Hub.Branches[*user.Id] = &hub.Branch{
		DBDsn:             branchDsn,
		DB:                branchDb,
		Profile:           user,
		AccessToken:       "",
		AccessTokenExpire: time.Time{},
	}

	// * Sign JWT
	token, err := crypto.SignJwt(c, &common.UserClaims{
		UserId: user.Id,
		Exp:    value.Ptr(time.Now().Add(24 * time.Hour).Unix()),
	})

	// * Set cookie
	c.Cookie(&fiber.Cookie{
		Name:        "user",
		Value:       token,
		Path:        "/",
		Domain:      "",
		MaxAge:      0,
		Expires:     time.Now().Add(24 * time.Hour),
		Secure:      false,
		HTTPOnly:    false,
		SameSite:    "",
		SessionOnly: false,
	})

	return c.JSON(response.Info(c, map[string]interface{}{
		"accessToken": *credentials.AccessToken,
		"token":       token,
	}))
}
