package raapchik

import (
	"context"
	"net/http"
)

func (h *Raapchik) findHandler(
	segs []PathSegment, routes []RaapchikRoute,
) *RaapchikRoute {
	for _, v := range routes {
		isMatch := false
		if len(v.RouteSegments) == len(segs) {
			isMatch = true
		inner:
			for j, segment := range v.RouteSegments {
				outSegment := segs[j]
				if !segment.IsVar {
					if segment.Seg != outSegment.Seg {
						isMatch = false
						break inner
					}
				}
			}
		}
		if isMatch {
			return &v
		}
	}

	return nil
}

func (h *Raapchik) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL
	segs := h.getPathSegmentsFromPath(uri.Path)
	handled := h.internalServe(w, r, segs)
	if !handled {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Invalid path\n"))
	}
}

func (h *Raapchik) internalServe(
	w http.ResponseWriter,
	r *http.Request,
	segs []PathSegment,
) bool {

	routes := h.getRoutes

	if r.Method == "PUT" {
		routes = h.putRoutes
	}
	if r.Method == "PATCH" {
		routes = h.patchRoutes
	}
	if r.Method == "POST" {
		routes = h.postRoutes
	}
	if r.Method == "DELETE" {
		routes = h.deleteRoutes
	}

	route := h.findHandler(segs, routes)
	if route == nil {
		for _, v := range h.children {
			handled := v.internalServe(w, r, segs)
			if handled {
				return true
			}
		}
		return false
	}

	params := make(map[string]string)
	for i, v := range route.RouteSegments {
		if v.IsVar {
			params[v.Var] = segs[i].Seg
		}
	}
	req := r
	if len(params) != 0 {
		ctx := context.WithValue(r.Context(), "params", params)
		req = r.WithContext(ctx)
	}

	for _, middleware := range h.middlewares {
		handled := middleware(w, req)
		if handled {
			return true
		}
	}

	route.Handler(w, req)

	return true
}
