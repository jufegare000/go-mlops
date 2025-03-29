# 📚 MLOps Hexagonal Demo App

Este es un proyecto de demostración de arquitectura hexagonal en Go, enfocado en explorar el ciclo de vida MLOps en un entorno desacoplado, modular y mantenible. La aplicación expone un cliente REST que se conecta a una base de datos MySQL.

---

## 🧱 Estructura del Proyecto
~~~.
├── cmd
│   └── literary_works_db.sql
├── go.mod
├── go.sum
├── go.work
├── go.work.sum
├── internal
│   ├── application
│   ├── domain
│   └── infrastructure
│       ├── src
│       │   ├── adapters
│       │   │   └── persistence
│       │   │       └── mysqlconf
│       │   │           ├── go.mod
│       │   │           ├── go.sum
│       │   │           └── msql_config.go
│       │   └── ports
│       │       └── rest
│       │           ├── go.mod
│       │           ├── go.sum
│       │           └── main
│       │               ├── config
│       │               │   └── load_env_variables.go
│       │               ├── db
│       │               │   └── database_interface.go
│       │               ├── main.go
│       │               └── rest_client_adapter
│       │                   ├── hello_world_rest_controller.go
│       │                   ├── quotes
│       │                   │   └── quotes_controller.go
│       │                   └── RestRouter.go
│       └── tests
│           ├── adapters
│           └── ports
├── pkg
│   └── database
│       ├── mysql database setup.md
│       └── README.md
└── README.md

~~~

---

## 🚀 Características

- ✅ Arquitectura hexagonal (puertos y adaptadores)
- ✅ Cliente REST en Go (`gin-gonic`)
- ✅ Conexión a base de datos MySQL
- ✅ Configuración externa vía `.env`
- ✅ Separación clara por módulos (`go.mod` + `go.work`)
- ✅ Extensible a otros adaptadores: mensajería, CLI, etc.

---

## 🔧 Requisitos

- Go 1.24+
- MySQL 8+
- [`go.work`](https://go.dev/doc/go-work) habilitado (si usás múltiples módulos)
- Archivo `.env` con los siguientes valores:

```env
HOST=localhost
MYSQL_USER=user
MYSQL_PASSWORD=secret
MYSQL_DATABASE=literary_works
PORT=3306
APPLICATION_PORT=8080
```

## 🛠 Instalación y ejecución
1. Cloná el repositorio
bash
Copy
Edit
git clone https://github.com/jufegare000/go-mlops.git
cd go-mlops
2. Iniciá el workspace
```
go work init ./internal/infrastructure/src/ports/rest ./internal/infrastructure/src/adapters/persistence/mysqlconf
```
3. Ejecutá el cliente REST
```
cd internal/infrastructure/src/ports/rest
go run main
```
## 🧪 Endpoints disponibles
Método	Ruta	Descripción
GET	/hello	Verifica que el servidor responde
GET	/quotes	Ejemplo de endpoint de dominio

## 🧩 Extensiones futuras
Adaptador de mensajería (Pub/Sub, Kafka)

Adaptador gRPC

Módulo de predicción ML

Observabilidad y métricas

Despliegue en GCP

🤝 Contribuciones
¡Contribuciones y sugerencias son bienvenidas! Este proyecto está diseñado para crecer modularmente.

📝 Licencia
MIT © [jufegare000]
