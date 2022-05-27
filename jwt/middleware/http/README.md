# jwtmw

## Index

- [func TokenFromContext(ctx context.Context) (*jwt.Token, error)](<#func-tokenfromcontext>)
- [func TokenMiddleware(j *jwt.JWT, next http.Handler) http.Handler](<#func-tokenmiddleware>)


## func TokenFromContext

```go
func TokenFromContext(ctx context.Context) (*jwt.Token, error)
```

TokenFromContext gets the raw token from the context and parses into a \*Token

## func TokenMiddleware

```go
func TokenMiddleware(j *jwt.JWT, next http.Handler) http.Handler
```

TokenMiddleware \.\.\.
