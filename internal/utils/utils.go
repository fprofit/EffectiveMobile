package utils

import "encoding/json"

func StructToString(b any) string {
	res, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(res)
}
