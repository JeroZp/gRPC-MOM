# 🧠 Proyecto: gRPC-MOM – Microservicios con gRPC y Middleware Orientado a Mensajes en Go

## 📌 Descripción General

**gRPC-MOM** es una arquitectura de microservicios desarrollada en Go que utiliza **gRPC** para la comunicación síncrona entre servicios y un **Middleware Orientado a Mensajes (MOM)** personalizado para manejar operaciones asincrónicas y garantizar la resiliencia del sistema.  
Todos los componentes son aplicaciones de consola, sin interfaces web, y están diseñados para ejecutarse en entornos de desarrollo o producción utilizando Docker.

---

## 🧱 Arquitectura de Microservicios
![gRPCs-MOM](https://github.com/user-attachments/assets/3418daca-525e-4071-b8f0-c16b084ece6e)


### 🧑‍💼 user-service
- **Responsabilidad**: Gestión de usuarios (registro, consulta, etc.).
- **Base de datos**: MySQL.
- **Comunicación**:
  - gRPC (directa).
  - MOM (solo si falla la directa).

### 💸 transaction-service
- **Responsabilidad**: Manejo de transacciones entre usuarios.
- **Base de datos**: MySQL.
- **Comunicación**:
  - gRPC directa.
  - Publica en MOM si `notification-service` falla.

### 📢 notification-service
- **Responsabilidad**: Registro y entrega de notificaciones.
- **Base de datos**: MySQL.
- **Comunicación**:
  - gRPC directa.
  - Escucha eventos del MOM si falló la entrega directa.

### 📬 mom-service
- **Responsabilidad**: Middleware personalizado que actúa como cola de mensajes.
- **Tecnología**: Solo Go, sin RabbitMQ, Kafka ni Redis.
- **Funcionamiento**:
  - Recibe mensajes de respaldo si falla una entrega directa.
  - El servicio destino puede consumir mensajes pendientes cuando se recupera.

### 🌐 API-Gateway
- **Responsabilidad**: Punto de entrada de las solicitudes REST.
- **Comunicación**:
  - REST con el cliente.
  - gRPC con los servicios internos.

### 🖥️ cli-client
- **Responsabilidad**: Cliente por consola para probar el sistema.
- **Comunicación**: Realiza peticiones HTTP al `API-Gateway`.

---

## 🔄 Flujo de Comunicación

```plaintext
cli-client ── REST ──> API-Gateway ── gRPC ──> Microservicio

                             |
                             |──(si falla)
                             ▼
                        mom-service
                             |
                             ▼
               Microservicio receptor recuperado
