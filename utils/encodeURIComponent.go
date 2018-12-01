package utils

import "net/url"

func EncodeURIComponent(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return ""
	}
	return u.String()
}
