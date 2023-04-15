package tests

import (
	"testing"

	"backend/modules"
	"backend/modules/config"
	"backend/modules/db/init"
	"backend/modules/fiber"
	"backend/modules/firebase"
	"backend/modules/hub"
	"backend/tests/modules/account"
	"backend/tests/modules/admin"
	"backend/tests/modules/profile"
)

func TestMain(m *testing.M) {
	// * Initialize modules
	Init()
	m.Run()
}

func Init() {
	modules.Conf = iconfig.Init()
	modules.Hub = ihub.Init(modules.Conf)
	modules.FirebaseApp, modules.FirebaseAuth = ifirebase.Init()
	modules.SqlDB, modules.DB = idbInit.InitTest()
	modules.Fiber = ifiber.Init()
}

func TestAccount(t *testing.T) {
	account.TestCallback(t)
	account.TestInvalidCallback(t)
}

func TestAdmin(t *testing.T) {
	admin.ImportLab(t)
}

func TestProfile(t *testing.T) {
	profile.TestGetState(t)
}
