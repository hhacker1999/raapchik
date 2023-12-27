# Raapchik: A Raapchik HTTP Router for Go

**Raapchik** is a sleek and easy-to-use HTTP router package for Go, designed to simplify route handling and support nested routes with middleware. With an API inspired by `chi` router, Raapchik provides a familiar and expressive syntax while offering additional features for handling complex routing scenarios.

## Features
- **Tiny Footprint:** Raapchik boasts a minimalistic design, ensuring a tiny footprint for your applications. No external dependencies and relies only on stdlib. Making it easy to maintain and scale.
- **Chi-Inspired API:** Raapchik's API is similar to `chi` router, making it easy for users familiar with chi to transition seamlessly.
- **Nested Routes:** Create nested routes effortlessly to organize your API endpoints hierarchically.
- **Middleware Support:** Raapchik allows you to attach middleware to individual routes or apply them globally. Middleware is also shared with nested routes, providing a clean and consistent way to handle common functionality across routes.
- **Easy Custom Middleware:** Raapchik makes it super easy to create and use your own middleware functions, giving you complete flexibility in handling HTTP requests.
- **Simplified URI Parameters:** Define multiple URL parameters in a single path, making it easy to extract and use them in your handlers.

## Installation

`
go get -u github.com/hhacker1999/raapchik
`

## Example usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/hhacker1999/raapchik"
)

func main() {
	port := 9000
	r := raapchik.New()
	r.Get("/", basicHandler)
	r.Get("/url/{{id}}/param", urlParamHandler)
	r.Get("/foo/{{id}}/bar/{{cid}}/baz", fooHandler)
	r.Group(func(r *raapchik.Raapchik) {
		r.Use(AuthMiddleware)
		r.Get("/user/info", userHandler)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request) bool {
	authenticated := false
	if !authenticated {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Not allowed"))
		return true
	}
	return false
}

func urlParamHandler(w http.ResponseWriter, r *http.Request) {
	id := raapchik.GetPathParam(r, "id")
	w.Write([]byte("Id received " + id))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	id := raapchik.GetPathParam(r, "id")
	cid := raapchik.GetPathParam(r, "cid")
	w.Write([]byte("Ids received " + id + " " + cid))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from user"))
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Raapchik!"))
}
```
