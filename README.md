# database

## links
* [standard go project layout](https://github.com/golang-standards/project-layout)
* [pq - postgres driver](https://pkg.go.dev/github.com/lib/pq)
* [generic interface around sql databases](https://pkg.go.dev/database/sql)

## docker commands

```zsh
$ docker build -t database .
```
```zsh
$ docker run -it -p 5432:5432 --rm --name database database
```