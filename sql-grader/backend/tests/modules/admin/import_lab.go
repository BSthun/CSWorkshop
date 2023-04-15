package admin

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/tests/endpoints/admin"
	"backend/types/payload"
	"backend/utils/text"
)

func ImportLab(t *testing.T) {
	t.Run("import lab", func(t *testing.T) {
		// * Read json from file /test.json
		var imp *payload.AdminLabImport

		bytes, err := os.ReadFile(text.RelativePath("tests/resources/lab_import_spotify.json"))
		if err != nil {
			t.Error(err)
		}

		if err := json.Unmarshal(bytes, &imp); err != nil {
			t.Error(err)
		}

		// * Import lab
		res, _, errr := admin.RequestImportLabPost(imp)
		if errr != nil {
			t.Error(errr)
		}

		assert.Equal(t, 200, res.StatusCode)
	})
}
