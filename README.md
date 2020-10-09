# RTR

RTR - regex table router. Simple HTTP router for building Go web apps.

## Install

```sh
go get -u github.com/smokehill/rtr
```

## Examples

```go
func main() {
    r := rtr.NewRouter()
    r.SetRoute("GET", "/api/books", listBooks)
    r.SetRoute("GET", "/api/books/([0-9]+)", getBook)
    r.SetRoute("POST", "/api/books", createBook)
    r.SetRoute("PUT", "/api/books/([0-9]+)", updateBook)
    r.SetRoute("DELETE", "/api/books/([0-9]+)", deleteBook)
    http.ListenAndServe(":80", r)
}
```