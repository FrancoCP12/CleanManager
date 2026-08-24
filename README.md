#  Sistema de Inventarios (Go + MySQL + Docker)

Proyecto colaborativo para aprender desarrollo backend con Go, manejo de bases de datos relacionales con MySQL y control de contenedores con Docker.

##  Requisitos previos
Asegúrate de tener instalado en tu computadora:
- [Go](https://golang.org/) (versión 1.18 o superior)
- [Docker y Docker Desktop](https://www.docker.com/)

---

##  Instrucciones de Configuración y Ejecución

### 1. Levantar la base de datos con Docker
Para no requerir una instalación nativa de MySQL, usamos Docker. Ejecuta el siguiente comando en tu terminal para levantar el contenedor:

```bash
docker run --name inventario-db -e MYSQL_ROOT_PASSWORD=secreto -e MYSQL_DATABASE=inventario_db -p 3306:3306 -d mysql:latest
```
---
### 2. Configura la tabla en MySQL
Conéctar al contenedor de Docker para crear la estructura de la base de datos: 
```bash
docker exec -it inventario-db mysql -u root -p

(Contraseña: **secreto**)
```
Una vez dentro de MySQL, ejecuta: 
```bash
USE inventario_db;

CREATE TABLE IF NOT EXISTS productos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    sku VARCHAR(50) NOT NULL UNIQUE,
    precio DECIMAL(10, 2) NOT NULL,
    stock_actual INT NOT NULL DEFAULT 0
);
```
### 3. Ejecuta la aplicacion en Go

Clona el repositorio, entra en la carpeta del proyecto e inicia el servidor web:
```bash

    En terminal: go run main.go
```
Abre tu navegador y entra a: http//localhost:8080

