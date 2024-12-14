# Dosetti

Medicine tracking app.

## Quick start guide



## Goose

```shell
setopt allexport ; . ./.env ; unsetopt allexport
goose -dir ./internal/database/migrations up
```

## Seed

```shell
turso db shell dosetti-dev < seed.sql
```
