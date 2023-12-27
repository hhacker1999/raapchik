package raapchik

import (
	"strings"
)

type PathSegment struct {
	Seg   string
	Var   string
	IsVar bool
}

func (h *Raapchik) getPathSegmentsFromPath(path string) []PathSegment {
	if len(path) == 1 && path == "/" {
		return []PathSegment{}
	}

	strippedPath := path
	if strippedPath[0] == '/' {
		strippedPath = strippedPath[1:]
	}
	if strippedPath[len(strippedPath)-1] == '/' {
		strippedPath = strippedPath[0 : len(strippedPath)-1]
	}

	rawSegments := strings.Split(strippedPath, "/")
	var segs []PathSegment
	for _, v := range rawSegments {
		temp := PathSegment{}
		if len(v) < 5 {
			temp.Seg = v
			temp.IsVar = false
			segs = append(segs, temp)
			continue
		}

		if v[0] == '{' && v[1] == '{' && v[len(v)-1] == '}' &&
			v[len(v)-2] == '}' {
			temp.IsVar = true
			temp.Var = v[2 : len(v)-2]
		} else {
			temp.Seg = v
			temp.IsVar = false
		}
		segs = append(segs, temp)

	}

	return segs
}
