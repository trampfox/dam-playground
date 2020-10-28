# dam-playground

Playing with Go and PostgreSQL JSONB data type.

## Configuration

### DAM data table

```sql
CREATE TABLE damdata (
    id serial NOT NULL PRIMARY KEY,
    info json NOT NULL
);
```

### Database URL

```bash
export DATABASE_URL="postgresql://postgres:test@localhost:5432/data"
```

## Build


## Run

## Load test

```bash
ab -p ab/payload.txt -T application/json -H 'Accept: application/json' -c 30 -n 3000 -l -k -v 2 http://localhost:8081/data > post_results.txt
```

## Examples

```sql
SELECT * FROM damdata WHERE info->>'temperature' = '15.5';
```
