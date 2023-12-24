package utils

import (
	"fmt"
	"strings"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func GetLogMessage(funcName string, number uint, errOrErrMsg any) string {
	return fmt.Sprintf("[Error] %s [%d]: %v", funcName, number, errOrErrMsg)
}
