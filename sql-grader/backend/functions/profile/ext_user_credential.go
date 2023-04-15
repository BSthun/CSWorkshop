package profile

import (
	"fmt"
	"strconv"
	"strings"

	"backend/modules"
	idb "backend/modules/db"
	"backend/types/embed"
	"backend/types/model"
	"backend/utils/text"
)

func ExtUserCredential(user *model.User) (*embed.Credential, error) {
	// * Check if the user has a credential
	if user.Credential == nil {
		// * Randomize password
		salt := strings.ToLower(*text.Random(text.RandomSet.UpperAlpha, 4))
		username := fmt.Sprintf("lab_%s%s", strconv.FormatUint(*user.Id, 32), salt)
		password := *text.Random(text.RandomSet.MixedAlphaNum, 16)
		exec := fmt.Sprintf("CREATE USER '%s'@'%%' IDENTIFIED BY '%s'", username, password)
		if result := modules.DB.Exec(exec); result.Error != nil {
			return nil, result.Error
		}
		user.Credential = &embed.Credential{
			Username: &username,
			Password: &password,
		}
		if result := modules.DB.
			Model(new(model.User)).
			Where(idb.Where(model.UserFieldId, "= ?"), user.Id).
			Update(model.UserFieldCredential, user.Credential); result.Error != nil {
			return nil, result.Error
		}
	}

	return user.Credential, nil
}
