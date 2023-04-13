package statement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveComments(t *testing.T) {
	t1 := "select * where table_schema = 'grader_dev1'"
	t2 := "/* ApplicationName=DataGrip 2023.1 */ select * where table_schema = 'grader_dev1'"
	t3 := "select * where table_schema /* ApplicationName=DataGrip 2023.1 */ = 'grader_dev1'"
	t4 := "select * where table_schema = 'grader_dev1' /* ApplicationName=DataGrip 2023.1 */ /* ApplicationName=DataGrip 2023.1 */"
	t5 := "select * where table_schema = 'grader_dev1' /* ApplicationName= /*DataGrip 2023.1 */"
	t6 := "select * where table_schema = 'grader_dev1' /*"

	t.Run("RemoveComments", func(t *testing.T) {
		assert.Equal(t, t1, *RemoveComments(t1))
	})
	t.Run("RemoveComments", func(t *testing.T) {
		assert.Equal(t, t1, *RemoveComments(t2))
	})
	t.Run("RemoveComments", func(t *testing.T) {
		assert.Equal(t, t1, *RemoveComments(t3))
	})
	t.Run("RemoveComments", func(t *testing.T) {
		assert.Equal(t, t1, *RemoveComments(t4))
	})
	t.Run("RemoveComments", func(t *testing.T) {
		assert.Equal(t, t1, *RemoveComments(t5))
	})
	t.Run("RemoveComments", func(t *testing.T) {
		assert.Equal(t, t1, *RemoveComments(t6))
	})
}
