# mongo-exporter

Este proyecto consiste en un sistema de recolección de métricas de una base de datos para monitorizar los datos. Desarrollado como parte de mis prácticas de primer año en el Grado Superior de Desarrollo de Aplicaciones Web en la empresa [iAR Soft](https://www.iar-soft.com/).

# API Endpoints

|Método|Endpoint|Descripción|
|-|-|-|
|`GET`|`/v1/healthcheck`|Obtener el estado de la API|
|`GET`|`/v1/hello`|Obtener el mensaje "hello world"|
|`GET`|`/v1/metrics`|Obtener una lista de métricas|
|`GET`|`/v1/metrics/:name`|Obtener una métrica por nombre|
|`GET`|`/v1/students`|Obtener una lista de todos estudiantes|
|`GET`|`/v1/students/:id`|Obtener un estudiante por ID|

# Uso

Clonar el repositorio
```
git clone https://github.com/SrVariable/mongo-exporter
```

Navegar al directorio del proyecto
```
cd mongo-exporter
```

Construye los contenedores
```
make
```

> [!NOTE]
>
> Si no tienes `make`, puedes ejecutar:
> ```
> docker compose down
> docker compose up --build -d
> ```

Una vez esté construido, puedes interactuar con la API utilizando tu navegador, curl o cualquier otro método que prefieras.

- Para obtener una lista de métricas:
```
curl localhost:8080/v1/metrics
```

- Para obtener el número de inserts a la base de datos
```
curl localhost:8080/v1/metrics/insert
```

Revisa [API Endpoints](#api-endpoints) para ver los endpoints disponibles.

# Referencias

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
