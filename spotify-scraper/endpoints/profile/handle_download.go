package profile

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/aliakseiz/go-mysqldump"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"backend/modules/config"
	"backend/modules/hub"
	"backend/types/common"
	"backend/types/response"
	"backend/utils/text"
)

func DownloadHandler(c *fiber.Ctx) error {
	// * Parse user claims
	u := c.Locals("u").(*jwt.Token).Claims.(*common.UserClaims)

	// * Get branch
	branch, ok := hub.Hub.Branches[*u.UserId]
	if !ok {
		return response.Error(c, true, "Unable to get branch")
	}

	db, err := sql.Open("mysql", branch.DBDsn)
	if err != nil {
		return response.Error(c, true, "Unable to open database", err)
	}
	defer db.Close()

	// Construct dump name
	name := fmt.Sprintf("u%03d_%s_%s", *branch.Profile.Id, time.Now().Format("20060102T150405"), *text.Random(text.StringSet.UpperAlpha, 3))

	// Register database with mysqldump
	dumper, err := mysqldump.Register(db, "./web/sql", name, "")
	if err != nil {
		return response.Error(c, true, "Unable to register database", err)
	}

	// Dump database to file
	err = dumper.Dump()
	if err != nil {
		return response.Error(c, true, "Unable to dump database", err)
	}

	return c.JSON(response.Info(c, map[string]any{
		"username": branch.Profile.Profile.DisplayName,
		"path":     config.C.BaseUrl + "/sql/" + name + ".sql",
	}))
}
