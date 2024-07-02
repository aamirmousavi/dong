package lib

import "encoding/json"

func Ptr[T any](v T) *T {
	return &v
}
func ToJsonIndent(t interface{}) string {
	b, _ := json.MarshalIndent(t, "", "  ")
	return string(b)
}
