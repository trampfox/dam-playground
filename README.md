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

## Examples

```sql
SELECT * FROM damdata WHERE info->>'temperature' = '15.5';
```
