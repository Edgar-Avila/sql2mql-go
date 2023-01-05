# MySQL2Mongo
This project translates a very small subset of SQL to a mongo syntax and wraps it in a cli interface, with options to translate the query or execute it given a uri and db name.

## Usage 
### Translate to mongo
```bash
sqlmql translate -f <file>
```

### Execute mongo
```bash
sqlmql execute -f <file> -u <uri> -d <db_name> 
```

### Get more help
```bash
sqlmql -h
```

## Running with go
You can get the same result running the project with ```go run .```
```bash
go run . translate -f examples/query.sql
```
