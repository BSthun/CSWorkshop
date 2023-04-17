package enroll

import (
	"bufio"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"backend/modules"
	"backend/modules/fiber/websocket"
	"backend/modules/hub"
	"backend/types/model"
	"backend/utils/text"
)

func ActMockData(enrollment *model.Enrollment) (*string, error) {
	// * Check database status
	if *enrollment.DbValid {
		return nil, fmt.Errorf("database is already valid")
	}

	// * Check generating
	if _, ok := modules.Hub.Mocks[*enrollment.Id]; ok {
		return nil, fmt.Errorf("already generating mock data")
	}

	// * Generate mock data
	token := text.Random(text.RandomSet.MixedAlphaNum, 16)
	mock := &ihub.Mock{
		Lines:     make([]string, 0),
		Token:     token,
		Conn:      nil,
		ConnMutex: new(sync.Mutex),
	}
	modules.Hub.Mocks[*enrollment.Id] = mock

	go func() {
		_ = HelpGenerateMockData(enrollment, mock)
	}()

	return token, nil
}

func HelpGenerateMockData(enrollment *model.Enrollment, mock *ihub.Mock) error {
	cmd := exec.Command(text.RelativePath(*enrollment.Lab.Generator))
	env := "DATABASE_DSN=" + strings.Replace(modules.Conf.MysqlDsn, "{{DB_NAME}}", *enrollment.DbName, 1)
	cmd.Env = append(cmd.Env, env)
	cmd.Dir = filepath.Dir(text.RelativePath(*enrollment.Lab.Generator))

	defer func() {
		mock.Append("SUCCESS")
		mock.ConnMutex.Lock()
		if mock.Conn != nil {
			websocket.HandleMockConnectionSwitch(mock)
		}
		mock.ConnMutex.Unlock()
		delete(modules.Hub.Mocks, *enrollment.Id)
	}()

	// Create a pipe to read the output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		mock.Append(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}

	// * Update database status
	if result := modules.DB.Model(enrollment).Update("db_valid", true); result.Error != nil {
		return result.Error
	}

	return nil
}
