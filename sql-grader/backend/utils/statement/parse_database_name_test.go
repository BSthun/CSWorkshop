package statement

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/utils/value"
)

func TestParseDatabaseName(t *testing.T) {
	tests := []struct {
		input    string
		expected *string
	}{
		{"SELECT t.* FROM mysql.general_log t WHERE argument LIKE '%grader%' LIMIT 501", value.Ptr("mysql")},
		{"select * from test1.orders where id = 123", value.Ptr("test1")},
		{"SELECT * FROM mydb.mytable", value.Ptr("mydb")},
		{"SELECT col1, col2 FROM mydb.mytable WHERE col1 = 'foo'", value.Ptr("mydb")},
		{"SELECT * FROM mydb.mytable WHERE col1 IN (SELECT col1 FROM mydb.mytable2)", value.Ptr("mydb")},
		{"SELECT * FROM mydb.mytable AS t1 JOIN mydb.mytable2 AS t2 ON t1.col1 = t2.col1", value.Ptr("mydb")},
		{"SELECT * FROM `my-db`.`my-table`", value.Ptr("my-db")},
		{"SELECT * FROM `my-db`.my_table WHERE my_col IN (SELECT col1 FROM `my-db2`.my_table2)", value.Ptr("my-db")},
		{"SELECT * FROM my_db.my_table WHERE my_col = '/*not_a_comment*/foo'", value.Ptr("my_db")},
		{"INSERT INTO storefront.users (name, age) VALUES ('Alice', 30)", value.Ptr("storefront")},
		{"update customers set name = 'Bob' where id = 456", nil},
		{"DELETE FROM orders WHERE id = 789", nil},
		{"SELECT database() AS db", nil},
		{"INSERT IGNORE INTO db1.table1 (id, name) VALUES (1, 'John')", value.Ptr("db1")},
		{"update db2.table2 set name = 'Bob' where id = 123", value.Ptr("db2")},
		{"DELETE FROM sth_db.test1 WHERE id = 456", value.Ptr("sth_db")},
		{"SELECT * FROM db4.table4 JOIN db5.table5 ON db4.table4.id = db5.table5.id", value.Ptr("db4")},
		{"SELECT * FROM mysql.general_log WHERE argument LIKE '%grader%'", value.Ptr("mysql")},
		{"SELECT * FROM db_name.table_name WHERE column_name = 'value'", value.Ptr("db_name")},
		{"UPDATE table_name SET column1 = 'value1' WHERE id = 1", nil},
		{"DELETE FROM table_name WHERE id = 2", nil},
		{"INSERT INTO table_name (column1, column2) VALUES ('value1', 'value2')", nil},
		{"CREATE DATABASE db_name", nil},
		{"CREATE TABLE db_name.table_name (id int, name varchar(255))", value.Ptr("db_name")},
		{"DROP TABLE db_name.table_name", value.Ptr("db_name")},
		{"DROP DATABASE db_name", nil},
		{"ALTER TABLE db_name22.table_name ADD COLUMN new_column varchar(255)", value.Ptr("db_name22")},
		{"GRANT ALL PRIVILEGES ON upper.* TO 'user'@'localhost'", value.Ptr("upper")},
		{"REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'user'@'localhost'", nil},
	}

	for _, test := range tests {
		t.Run("ParseDatabaseName", func(t *testing.T) {
			val := ParseDatabaseName(test.input)
			if test.expected == nil {
				assert.Nil(t, val)
				return
			} else if val == nil {
				t.Error("expected non-nil value")
				return
			}
			assert.Equal(t, *test.expected, *val)
		})
	}
}
