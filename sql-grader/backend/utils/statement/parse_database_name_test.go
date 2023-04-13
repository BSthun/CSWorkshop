package statement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDatabaseName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"SELECT t.* FROM mysql.general_log t WHERE argument LIKE '%grader%' LIMIT 501", "mysql"},
		{"select * from test1.orders where id = 123", "test1"},
		{"SELECT * FROM mydb.mytable", "mydb"},
		{"SELECT col1, col2 FROM mydb.mytable WHERE col1 = 'foo'", "mydb"},
		{"SELECT * FROM mydb.mytable WHERE col1 IN (SELECT col1 FROM mydb.mytable2)", "mydb"},
		{"SELECT * FROM mydb.mytable AS t1 JOIN mydb.mytable2 AS t2 ON t1.col1 = t2.col1", "mydb"},
		{"SELECT * FROM `my-db`.`my-table`", "my-db"},
		{"SELECT * FROM `my-db`.my_table WHERE my_col IN (SELECT col1 FROM `my-db2`.my_table2)", "my-db"},
		{"SELECT * FROM my_db.my_table WHERE my_col = '/*not_a_comment*/foo'", "my_db"},
		{"INSERT INTO storefront.users (name, age) VALUES ('Alice', 30)", "storefront"},
		{"update customers set name = 'Bob' where id = 456", "default"},
		{"DELETE FROM orders WHERE id = 789", "default"},
		{"SELECT database() AS db", ""},
		{"INSERT IGNORE INTO db1.table1 (id, name) VALUES (1, 'John')", "db1"},
		{"update db2.table2 set name = 'Bob' where id = 123", "db2"},
		{"DELETE FROM sth_db.test1 WHERE id = 456", "sth_db"},
		{"SELECT * FROM db4.table4 JOIN db5.table5 ON db4.table4.id = db5.table5.id", "db4"},
	}

	for _, test := range tests {
		t.Run("ParseDatabaseName", func(t *testing.T) {
			assert.Equal(t, test.expected, ParseDatabaseName(test.input))
		})
	}
}
