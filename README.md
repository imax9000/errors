# Slightly better `errors` package

Just so you don't have to construct a holding variable separately and can write just this:

```go
if err, ok := errors.As[*someError](err); ok {
  // ...
}
```
