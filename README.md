# GO Fiber + GROM + gRPC (for microservice)

**baca terlebih dahulu sampai bawah/selesai sebelum clone repo ini...**

---

## Structure Directory of System

| Name         | Path              | Description                                                                                                     | In Generate |
|--------------|-------------------|-----------------------------------------------------------------------------------------------------------------|:-----------:|
| docs         | /docs             | hasil generate dari **swagger** (swag init)                                                                     |      ✅      |
| proto        | /proto            | hasil generate dari protoc, menyimpan file **protobuf** dan juga generate nya                                   |      ✅      |
| apps         | /src/apps         | semua **Ready App** dari **Server, Database, gRPC server, Redis, RabbitMQ, WebSocket**                          |      ❌      |
| configs      | /src/configs      | cara/method pengambilan value dari **.env**                                                                     |      ❌      |
| controllers  | /src/controllers  | handling request (api) value nya apa saja/type                                                                  |      ❌      |
| dto          | /src/dto          | **(data transfer object)** management format request & response                                                 |      ❌      |
| handlers     | /src/handlers     | membuat method handler seperti **try catch** dan lain2...                                                       |      ❌      |
| helpers      | /src/helpers      | menampung semua logic kebutuhan yang bisa di **reusable**                                                       |      ❌      |
| messages     | /src/messages     | **format penulisan** apapun (semua) sehingga jika ada perubahan (text) tinggal rubah 1x saja                    |      ❌      |
| middlewares  | /src/middlewares  | **jembatan logic** sebelum masuk ke **logic utama**, biasanya seperti **token validation** atau yang lainnya... |      ❌      |
| entities     | /src/entities     | schema tables digunakan untuk **validasi table database** dan **migration table**                               |      ❌      |
| repositories | /src/repositories | logic khusus untuk pengolahan database (insert, read, update, delete)                                           |      ❌      |
| remotes      | /src/remotes      | function pengambilan data dari gRPC service lain (communications)                                               |      ❌      |
| routers      | /src/routers      | tempat routing endpoint url (api) dan juga pemasangan middleware                                                |      ❌      |
| services     | /src/services     | tempat **logic utama** dari service ini (logic utama hanya boleh disini)                                        |      ❌      |
| utils        | /src/utils        | method penting yang sangat digunakan dan bisa saja reusable                                                     |      ❌      |

---

## Files Information

| Name                         | Extension    | Path   | Description                                                        |
|------------------------------|--------------|--------|--------------------------------------------------------------------|
| swagger                      | .json, .yaml | /docs  | untuk dokumentasi API                                              |
| proto                        | .proto       | /proto | schema protobuf untuk komunikasi gRPC                              |
| environment (select)         | .env         | /      | untuk memilih sekarang ini environment apa                         |
| environment (value)          | .env.*       | /      | isi variable yang ingin di expose                                  |
| package.json script (nodejs) | .json        | /      | untuk menyimpan segala format execute biar mempermudah development |
| database (sqlite)            | .db          | /      | file database lokal (file)                                         |

---

## Structure Directory of Microservice

```
go-fiber-gorm-grpc-for-microservice
│
└─── api-gateway
│       file project...
│
└─── main-service (this repo)
│       file project...
│
└─── grpc-basic-service
        file project...
```

---

## How to Use

1. install Docker Desktop or Docker (whatever)
2. create container :
```bash
docker-compose up -d
```
3. create all container microservices testing (link bellow)

---

## Clone All Microservices Testing

- [Api Gateway](https://github.com/jefripunza/nginx-load-balance-microservice.git)
- [gRPC Basic Service](https://github.com/jefripunza/example-grpc-basic-service.git)

![Schemas Routing (diagram)](git_assets/schema.jpg)

---

## URL Testing (result)

- Add [http://localhost/api/main/v1/add/3/6](http://localhost/api/main/v1/add/3/6)
  ![Operation Add](git_assets/operation-add.png)
- Multiply [http://localhost/api/main/v1/multiply/3/6](http://localhost/api/main/v1/multiply/3/6)
  ![Operation Multiply](git_assets/operation-multiply.png)
- Result (gRPC Basic Service) in Docker
  ![Docker Result](git_assets/docker-result.png)

Note :
- clone semua ini didalam 1 folder (seperti di **Folder Structure**)
- perhatikan port default pada docker network (disini di set 172.17.0.1) (lihat di Dockerfile)
