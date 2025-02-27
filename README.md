# greenlight
A JSON API for retrieving and managing information about movies


# Run application and verify
```
    go run ./cmd/api
    go run ./cmd/api/main.go -port=3030 -env=production
```

```
    curl -i localhost:4000/v1/healthcheck
    curl -i -X POST localhost:4000/v1/healthcheck
```

```
    curl -i localhost:4000/v1/movies/123
    curl -i -X POST localhost:4000/v1/movies
```

