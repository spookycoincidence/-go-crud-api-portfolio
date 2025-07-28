# 🛠️ Project API - CRUD en Go con Gin

API RESTful desarrollada en **Golang** utilizando el framework **Gin**. Permite gestionar proyectos mediante un CRUD completo, con endpoints para crear, listar, obtener, actualizar y eliminar proyectos.

La arquitectura del proyecto sigue buenas prácticas de separación en capas:
- **Modelo**
- **Repositorio**
- **Handler**
- **Rutas**

Se utiliza **GORM** para la conexión y manipulación de la base de datos.

Además, esta API puede ser **deployada fácilmente en Railway**.

---

## 📌 Endpoints Disponibles

| Método | Endpoint           |Descripción                    |
|--------|------------------- |-----------------------------  |
| POST   | `/projects`        | Crear un nuevo proyecto       |
| GET    | `/projects`        | Obtener todos los proyectos   |
| GET    | `/projects/:id`    | Obtener un proyecto por ID    |
| PUT    | `/projects/:id`    | Actualizar un proyecto por ID |
| DELETE | `/projects/:id`    | Eliminar un proyecto por ID   |

---

## 🗂️ Estructura

```bash
/project
handler.go # Handlers (controladores HTTP)
project.go # Definición del modelo Project
repository.go # Acceso a la base de datos
```

Las rutas están definidas en el archivo `router/router.go`, y la migración del modelo se aplica en `main.go` mediante `db.AutoMigrate`.

---

## 🚀 Cómo Ejecutarlo Localmente

1. Clonar el repositorio:
```bash
git clone <URL_DE_TU_REPO>
cd <nombre>
```
2. Instalar dependencias::
```bash
go mod tidy
```
3. Configurar la base de datos en main.go o utilizar SQLite para facilitar la prueba.

4. Ejecutar el servidor:
```bash
go run main.go
```
5. La API quedará disponible en:
```bash
http://localhost:8080
```
##✅ Payload de Ejemplo:

```bash
POST /projects
{
    "name": "Nombre",
    "description": "Descripcion"
}
```
Respuesta:

```bash
{
    "ID": 1,
    "CreatedAt": "2025-07-22T12:00:00Z",
    "UpdatedAt": "2025-07-22T12:00:00Z",
    "DeletedAt": null,
    "name": "Nombre",
    "description": "Descripcion"
}
```

## Tecnologías
* ⚛️ Golang
* 🧠 Gin (framework)
* 💜 SQLite / MySQL / PostgreSQL (según la configuración)

## 🚀 Deploy en Railway
Requisitos
Tener una cuenta gratuita en Railway.
Tener Railway CLI instalado (opcional para deploy manual).

Pasos para desplegar:
Ir a https://railway.app/new.

Seleccionar Deploy from GitHub Repo y conectar el repositorio de este proyecto.
Railway detectará que es un proyecto Go y lo buildará automáticamente.
Configurar la base de datos desde Railway si es necesario (por ejemplo, añadiendo un plugin de PostgreSQL o MySQL).
Agregar las variables de entorno necesarias si el proyecto las requiere (ej: conexión a la DB).
Deploy automático o manual al hacer push a la branch principal.
Railway generará una URL pública para consumir la API en la nube.

## 📝 Inspiración
Este proyecto se inspiró en la estructura y el enfoque del repositorio gothinkster/golang-gin-realworld-example-app, el cual implementa una API completa basada en el estándar RealWorld API Spec.
Licencia original: MIT License

El código fue modificado con fines educativos y de demostración técnica, sin fines comerciales.

## Desarrollado con ❤️ por spookycoincidence


