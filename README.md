# go-microservices-test

## Run
```
make up
```

## Methods
user-service on http://localhost:8080
order-service on http://localhost:8081

#### user-service

| description  |  method |  url |   
|---|---|---|
|  ping | GET  |  http://localhost:8080/ping |
|  create user |  POST | http://localhost:8080/user  |
|  get by id |   GET| http://localhost:8080/user/{id}   | 

#### order-service

| description  |  method |  url |   
|---|---|---|
|  ping | GET  |  http://localhost:8081/ping |
|  create order |  POST | http://localhost:8081/user  |
|  get by id |   GET| http://localhost:8081/order/{id}   | 