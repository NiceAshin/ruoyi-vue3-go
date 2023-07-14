package utils

import (
	"github.com/google/uuid"
	"strings"
)

// 生成新的GUID
func RandomUUID() string {
	id := uuid.New()
	return id.String()
}

// 生成新的GUID
func SimpleUUID() string {
	id := uuid.New()
	guid := strings.ReplaceAll(id.String(), "-", "")
	return guid
}
