package raapchik

import (
	"net/http"
)

type RaapchikRoute struct {
	RouteSegments []PathSegment
	Handler       http.HandlerFunc
}

type Raapchik struct {
	getRoutes     []RaapchikRoute
	postRoutes    []RaapchikRoute
	putRoutes     []RaapchikRoute
	deleteRoutes  []RaapchikRoute
	patchRoutes   []RaapchikRoute
	optionsRoutes []RaapchikRoute
	children      []*Raapchik
	middlewares   []func(http.ResponseWriter, *http.Request) bool
}

func New() *Raapchik {
	return &Raapchik{}
}

func NewSubRaapchik(
	middlewares []func(http.ResponseWriter, *http.Request) bool,
) *Raapchik {
	return &Raapchik{
		middlewares: middlewares,
	}
}

func (u *Raapchik) Get(path string, handler http.HandlerFunc) {
	route := RaapchikRoute{
		RouteSegments: u.getPathSegmentsFromPath(path),
		Handler:       handler,
	}
	u.getRoutes = append(u.getRoutes, route)
}

func (u *Raapchik) Put(path string, handler http.HandlerFunc) {
	route := RaapchikRoute{
		RouteSegments: u.getPathSegmentsFromPath(path),
		Handler:       handler,
	}
	u.putRoutes = append(u.putRoutes, route)
}

func (u *Raapchik) Post(path string, handler http.HandlerFunc) {
	route := RaapchikRoute{
		RouteSegments: u.getPathSegmentsFromPath(path),
		Handler:       handler,
	}
	u.postRoutes = append(u.postRoutes, route)
}

func (u *Raapchik) Delete(path string, handler http.HandlerFunc) {
	route := RaapchikRoute{
		RouteSegments: u.getPathSegmentsFromPath(path),
		Handler:       handler,
	}
	u.deleteRoutes = append(u.deleteRoutes, route)
}

func (u *Raapchik) Patch(path string, handler http.HandlerFunc) {
	route := RaapchikRoute{
		RouteSegments: u.getPathSegmentsFromPath(path),
		Handler:       handler,
	}
	u.patchRoutes = append(u.patchRoutes, route)
}

func (u *Raapchik) Options(path string, handler http.HandlerFunc) {
	route := RaapchikRoute{
		RouteSegments: u.getPathSegmentsFromPath(path),
		Handler:       handler,
	}
	u.optionsRoutes = append(u.optionsRoutes, route)
}

func (h *Raapchik) Group(fnc func(r *Raapchik)) {
	r := NewSubRaapchik(h.middlewares)
	h.children = append(h.children, r)
	fnc(r)
}

func (h *Raapchik) Use(
	middleware func(http.ResponseWriter, *http.Request) bool,
) {
	h.middlewares = append(h.middlewares, middleware)
}

// func (h *OurHandler) PrettyPGet() {
// 	for _, v := range h.getRoutes {
// 		if len(v) == 0 {
// 			fmt.Println("/")
// 			continue
// 		}
// 		var str string
// 		for _, seg := range v {
// 			if seg.IsVar {
// 				str += "/"
// 				str += "{{"
// 				str += seg.Var
// 				str += "}}"
// 			} else {
// 				str += "/"
// 				str += seg.Seg
// 			}
//
// 		}
// 		fmt.Println(str)
// 	}
// }
