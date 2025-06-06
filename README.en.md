# mongo-exporter

This project consists in a Metric Collection System for a database to monitor
the data. It's developed as part of my first year internship in
[iAR Soft](https://www.iar-soft.com/).

# API Endpoints

| Method | Endpoint                  | Description                                                                          |
| ------ | ------------------------- | ------------------------------------------------------------------------------------ |
| `GET`  | `/v1/healthcheck`         | Get the health status of the API                                                     |
| `GET`  | `/v1/hello`               | Get "hello world" message                                                            |
| `GET`  | `/v1/metrics/collection`  | Get metrics related to collection. Requires query parameters `dbName` and `collName` |
| `GET`  | `/v1/metrics/connections` | Get metrics related to server connections                                            |
| `GET`  | `/v1/metrics/cpu`         | Get metrics related to server CPU                                                    |
| `GET`  | `/v1/metrics/opcounters`  | Get metrics related to server operations                                             |
| `GET`  | `/v1/metrics/ram`         | Get metrics related to server RAM                                                    |

# Usage

Clone the repository

```
git clone https://github.com/SrVariable/mongo-exporter
```

Navigate to the project folder

```
cd mongo-exporter
```

Create `.env` file following the `.env.example` file to configure the environment
variables. For default configuration, just copy `.env.example` to `.env`:

```
cp .env.example .env
```

`.env` file should look like this:

```
APP_PORT=8080

DB_NAME=MyDatabaseName
DB_HOST=mongo
DB_PORT=27017

GRAFANA_USER=admin
GRAFANA_PASS=grafana
```

Build the containers

```
make
```

> [!NOTE]
>
> If you don't have `make`, you can run:
>
> ```
> docker compose down
> docker compose up --build -d
> ```

Once it's built, you can interact with the API using your browser, curl, or
any method you prefer.

- To get metrics related to CPU:

```
curl http://localhost:8080/v1/metrics/cpu
```

- To get metrics related to Collection `bar` from Database `foo`:

```
curl http://localhost:8080/v1/metrics/collection?dbName=foo&collName=bar
```

Check [API Endpoints](#api-endpoints) to see available endpoints.

# References

- https://go.dev/doc/tutorial/web-service-gin
- https://youtu.be/67yGbvyM1is
- https://gin-gonic.com/docs
- https://github.com/gin-gonic/examples/tree/master/group-routes
- https://stackoverflow.com/questions/33322103/multiple-froms-what-it-means
- https://stackoverflow.com/questions/75973805/creating-dockerfile-for-golang-web-application
- https://www.docker.com/blog/developing-go-apps-docker/
- https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo
- https://www.youtube.com/watch?v=bDWApqAUjEI
- https://www.youtube.com/watch?v=g7cNQB2kCgE
- https://www.mongodb.com/docs/manual/reference/command/serverStatus/
- https://github.com/docker/awesome-compose/tree/master/prometheus-grafana
