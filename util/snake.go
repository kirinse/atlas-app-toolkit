package util

import (
	"github.com/kirinse/atlas-app-toolkit/util/cases"
)

// CamelToSnake is deprecated, use github.com/kirinse/util/cases.CamelToSnake.
func CamelToSnake(str string) string {
	return cases.CamelToSnake(str)
}
