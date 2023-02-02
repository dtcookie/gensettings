package golang

import (
	"fmt"
	"strings"
)

var replacements = map[string]string{
	"ConsiderForAi": "ConsiderForAI",
}

func staticReplace(name string) string {
	if replacement, found := replacements[name]; found {
		return replacement
	}
	return name
}

func PrettyProp(name string) string {
	name = staticReplace(name)
	if strings.HasSuffix(name, "Id") {
		name = name[:len(name)-2] + "ID"
	}
	if name == "PgToServicePropagation" {
		name = "PGToServicePropagation"
	}
	if name == "PgToHostPropagation" {
		name = "PGToHostPropagation"
	}

	return name
}

func Plural(s string) string {
	lower := strings.ToLower(s)
	if strings.HasSuffix(lower, "complex") {
		return s + "es"
	}
	if strings.HasSuffix(lower, "presets") {
		return s + "List"
	}
	s = s + "s"
	if strings.HasSuffix(s, "ys") {
		s = s[:len(s)-2] + "ies"
	}
	return s
}

func Singular(s string) (string, error) {
	if strings.HasSuffix(s, "s") {
		return strings.TrimSuffix(s, "s"), nil
	}
	if strings.HasSuffix(s, "List") {
		return strings.TrimSuffix(s, "List"), nil
	}
	// if s == "metadata" {
	// 	return s, nil
	// }
	// if s == "detection" {
	// 	return s, nil
	// }
	// if s == "dimensionFilter" {
	// 	return s, nil
	// }
	// if s == "data" {
	// 	return s, nil
	// }
	// if s == "allowlist" {
	// 	return s, nil
	// }
	// if s == "excludeNic" {
	// 	return s, nil
	// }
	// if s == "excludeIp" {
	// 	return s, nil
	// }
	// if s == "detectionConditionsLinux" {
	// 	return s, nil
	// }
	return "", fmt.Errorf("don't know the singular of %s" + s)
}
