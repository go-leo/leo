package sqlx

import "testing"

func TestCheckSqlSafe(t *testing.T) {
	assertSql(t, false, "insert abc")
	assertSql(t, true, "insert user (id,name) value (1, 'qm')")
	assertSql(t, true, "SELECT * FROM user")
	assertSql(t, true, "delete from user")
	assertSql(t, true, "drop TABLE user")
	assertSql(t, true, ";TRUNCATE from user")
	assertSql(t, false, "update")
	assertSql(t, false, "trigger")
	assertSql(t, true, "and name like '%s123%s'")
	assertSql(t, true, "convert(name using GBK)")
	assertSql(t, true, "order by id,if(1=2,1,(sleep(100)));")
}

func assertSql(t *testing.T, injection bool, sql string) {
	if injection != CheckSqlInjection(sql, false) {
		t.Fatal(injection, CheckSqlInjection(sql, false), sql, "not equal")
	}
}
