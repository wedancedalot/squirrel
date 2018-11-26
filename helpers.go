package squirrel

import (
    "fmt"
)

func Field(tableName, fieldName string) string {
    return fmt.Sprintf("%s.%s", tableName, fieldName)
}

func joinTable(joinType, tableName, fieldName string, sourceTable...string) string {
    if len(sourceTable) == 0 {
        return fmt.Sprintf("%s %s USING (%s)", joinType, tableName, fieldName)
    } else {
        return fmt.Sprintf("%s %s ON %s = %s", joinType, tableName, Field(tableName, fieldName), Field(sourceTable[0], fieldName))
    }
}

func GroupConcat(tableName, fieldName string, as string) string {
    return fmt.Sprintf("GROUP_CONCAT(%s.%s) as %s", tableName, fieldName, as)
}