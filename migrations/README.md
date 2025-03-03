# SQL Migration files for our database

# Commands
```
    cd  migrations
    goose create create_movie_table sql

    goose postgres postgres://priyashekhar:password@localhost:5432/greenlight up    
```

```
    psql postgres 
    postgres=# select current_database();
    postgres=# \c greenlight
    greenlight=# \dt
```
