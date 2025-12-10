package tools

import (
	"github.com/google/uuid"
	"strings"
)

// UUIDHex 生成32位UUID字符串
func UUIDHex() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
