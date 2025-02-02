# Dosetti

Dosetti is a progressive web application for tracking medicine intake.

## Installing self-hosted version

Self-hosted version can be run using Docker image or preferably using Docker Compose:

```yaml
services:
  dosetti:
    image: ghcr.io/jariinc/dosetti:release
    restart: unless-stopped
    volumes:
      - dosetti_data:/data
    environment:
      DATABASE_URL: file:./data/dosetti.db

volumes:
  dosetti_data:
```

For full example see [docker-compose.yaml](./docker-compose.yaml)

`release` tag will always point to latest stable release. Alternatively each release will have version number tags as follows:

| Release | tags                           |
| ------- | ------------------------------ |
| v1.1.1  | `release` `v1.1.1` `v1.1` `v1` |
| v1.1.0  | `v1.1.0`                       |
| v1.0.0  | `v1.0.0` `v1.0`                |
| v0.9.0  | `v0.9.0` `v0.0` `v0`           |

## Development

```sh
cp .env.local .env
just install
just migrate
just seed
just run
# or
just watch
```
