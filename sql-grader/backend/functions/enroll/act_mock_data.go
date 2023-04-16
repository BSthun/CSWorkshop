package enroll

import (
	"bufio"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"backend/modules"
	"backend/modules/hub"
	"backend/types/model"
	"backend/utils/text"
)

func ActMockData(enrollment *model.Enrollment) (*string, error) {
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
	return nil
}
