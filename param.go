package raapchik

import "net/http"

func GetPathParam(r *http.Request, param string) string {
	paramMap, ok := r.Context().Value("params").(map[string]string)
	if ok {
		value, ok := paramMap[param]
		if ok {
			return value
		}
	}
	return ""
}
