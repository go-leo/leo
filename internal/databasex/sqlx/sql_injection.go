package sqlx

import (
	"regexp"
	"strings"

	"github.com/go-leo/gox/stringx"
)

/**
 * SQL语法检查正则：符合两个关键字（有先后顺序）才算匹配
 * <p>
 * 参考: mybatis-plus-core/src/main/java/com/baomidou/mybatisplus/core/toolkit/sql/SqlInjectionUtils.java
 */
/**
 * 使用'、;或注释截断SQL检查正则
 * <p>
 * 参考: mybatis-plus-core/src/main/java/com/baomidou/mybatisplus/core/toolkit/sql/SqlInjectionUtils.java
 */

var sqlSyntaxPattern = regexp.MustCompile("(?i)(insert|delete|update|select|create|drop|truncate|grant|alter|deny|revoke|call|execute|exec|declare|show|rename|set).+(into|from|set|where|table|database|view|index|on|cursor|procedure|trigger|for|password|union|and|or)")
var sqlCommentPattern = regexp.MustCompile("(?i)'.*(or|union|--|#|/*|;)")

// CheckSqlInjection 检查参数是否存在 SQL 注入
// value 检查参数
// return true 非法 false 合法
func CheckSqlInjection(value string, funcAllowed bool) bool {
	if stringx.IsBlank(value) {
		return false
	}
	if funcAllowed {
		return sqlCommentPattern.MatchString(value) || sqlSyntaxPattern.MatchString(value)
	}
	// 不允许使用任何函数（不能出现括号），否则无法检测后面这个注入 order by id,if(1=2,1,(sleep(100)));
	return strings.Contains(value, "(") || sqlCommentPattern.MatchString(value) || sqlSyntaxPattern.MatchString(value)
}
