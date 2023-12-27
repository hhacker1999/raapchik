# Raapchik: A Raapchik HTTP Router for Go

**Raapchik** is a sleek and easy-to-use HTTP router package for Go, designed to simplify route handling and support nested routes with middleware. With an API inspired by `chi` router, Raapchik provides a familiar and expressive syntax while offering additional features for handling complex routing scenarios.

## Features
- **Chi-Inspired API:** Raapchik's API is similar to `chi` router, making it easy for users familiar with chi to transition seamlessly.
- **Nested Routes:** Create nested routes effortlessly to organize your API endpoints hierarchically.
- **Middleware Support:** Raapchik allows you to attach middleware to individual routes or apply them globally. Middleware is also shared with nested routes, providing a clean and consistent way to handle common functionality across routes.
- **Easy Custom Middleware:** Raapchik makes it super easy to create and use your own middleware functions, giving you complete flexibility in handling HTTP requests.

## Installation

`
go get -u github.com/hhacker1999/raapchik
`
