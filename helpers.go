package squirrel

import (
    "fmt"
)

func Field(tableName, fieldName string) string {
    return fmt.Sprintf("%s.%s", tableName, fieldName)
}