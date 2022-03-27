# db-migrate-name

database migration name generator

## install

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
