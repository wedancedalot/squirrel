package squirrel

import (
	"fmt"
)

func Field(tableName, fieldName string) string {
	return fmt.Sprintf("%s.%s", tableName, fieldName)
}

func FieldAs(tableName, fieldName string, as string) string {
	return fmt.Sprintf("%s.%s AS %s", tableName, fieldName, as)
}

func joinTable(joinType, tableName, fieldName string, sourceTable ...string) string {
	if len(sourceTable) == 0 {
		return fmt.Sprintf("%s %s USING (%s)", joinType, tableName, fieldName)
	} else {
		return fmt.Sprintf("%s %s ON %s = %s", joinType, tableName, Field(tableName, fieldName), Field(sourceTable[0], fieldName))
	}
}

func FieldAgg(f, tableName, fieldName string, as string) string {
	return fmt.Sprintf("%s(%s.%s) as %s", f, tableName, fieldName, as)
}

func GroupConcat(tableName, fieldName string, as string) string {
	return FieldAgg("GROUP_CONCAT", tableName, fieldName, as)
}
