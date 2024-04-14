package characterutil

import "strings"

// StitchingBuilderStr 字符串拼接
func StitchingBuilderStr(args ...string) string {
	var build strings.Builder
	for _, v := range args {
		build.WriteString(v)
	}
	return build.String()
}
