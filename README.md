# planets-fact-site-graphql

comments api for information of planets.

## Requirements

Install [docker desktop](https://www.docker.com/products/docker-desktop/)

We will use [docker compose](https://docs.docker.com/compose/) for this api

### Start

```shell
docker compose --env-file=./.env up --build
```

### Stop

Just press `<CTRL>+C`

### Delete containers

```shell
docker compose --env-file=./.env down
```

## Metadata, Migrations & Seeds

The next commands will create your api & database structure, and add initial data.

Select **All (all available databases)** option during _migrate apply_ and _seed apply_

```shell
docker compose exec hasura-cli hasura metadata apply
docker compose exec hasura-cli hasura migrate apply
docker compose exec hasura-cli hasura metadata reload
docker compose exec hasura-cli hasura seed apply
```

## Useful links

* [Hasura console](http://localhost:8088) Use admin secret (HASURA_GRAPHQL_ADMIN_SECRET) from .env file

## License

MIT
