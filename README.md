# my-go-app

## Run

```bash
docker-compose up --build
```

## Endpoints

```bash
GET /
GET /books
GET /api/books
GET /requests
GET /healthz

--

$ curl http://127.0.0.1:8080/requests
{
    "count": 6
}
```
