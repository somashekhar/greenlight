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

# Error Handling 
```
    curl -d '{"title": "Moana","rating":"PG"}' localhost:4000/v1/movies
    curl -d '{"title": "Moana"}{"title": "Top Gun"}' localhost:4000/v1/movies
    curl -d '{"title": "Moana"} :~()' localhost:4000/v1/movies

    wget -O /tmp/largefile.json https://www.alexedwards.net/static/largefile.json
    curl -d @/tmp/largefile.json localhost:4000/v1/movies
```

```
    curl -d '{"title": "Moana","runtime": "107 mins"}' localhost:4000/v1/movies
    
    curl -d '{"title": "Moana","runtime": 107}' localhost:4000/v1/movies
    curl -d '{"title": "Moana","runtime": "107 minutes"}' localhost:4000/v1/movies
```

