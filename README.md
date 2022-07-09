# db-migrate-name

database migration name generator

![Build Status](https://github.com/jcyamacho/db-migrate-name/workflows/Go/badge.svg)

## install

### install script
 binary will be (go env GOPATH)/bin/db-migrate-name
```
curl -sSfL https://raw.githubusercontent.com/jcyamacho/db-migrate-name/main/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
or install it into ./bin/
```
curl -sSfL https://raw.githubusercontent.com/jcyamacho/db-migrate-name/main/install.sh | sh
```
### using go

```
go install github.com/jcyamacho/db-migrate-name/cmd/db-migrate-name@latest
```

## usage

```
db-migrate-name -sep - add person table
```

**output**: 20220327174544-add-person-table

### flags

-   `-lc`: lower case, (default `false`)
-   `-sep`: separator, (default "\_")
-   `-tp`: time prefix, (default true)
-   `-tpf`: time prefix format, (default "20060102150405")
