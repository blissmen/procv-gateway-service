# 🚪 Gateway Service – API Gateway for ProCV

This is the Gateway microservice for ProCV – responsible for routing, authentication, and request validation across the microservices powering the ProCV platform.

---

## 🚀 Responsibilities

- 🌐 Route requests to appropriate microservices (CV, Auth, Profile, etc.)
- 🔐 Handle authentication and token validation
- 🛡 Rate limiting and basic throttling
- 🧰 Request/response logging for observability
- ⚠ Centralized error handling and response shaping

---

## 🛠 Tech Stack

- **Golang** – High-performance core
- **Gin** – Lightweight HTTP web framework
- **JWT** – Authentication middleware
- **Envconfig** – Environment variable management
- **Swagger (optional)** – API documentation

---

## 📁 Folder Structure

```bash
gateway-service/
├── main.go
├── router/
│   └── routes.go
├── middleware/
│   ├── auth.go
│   └── logger.go
├── config/
│   └── config.go
├── handlers/
│   └── healthcheck.go
└── go.mod
```

---

## ⚙️ Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/your-username/procv-gateway.git
cd procv-gateway
```

### 2. Set environment variables
Create a `.env` or export the following:
```env
PORT=8080
JWT_SECRET=your_super_secret_key
CV_SERVICE_URL=http://localhost:8001
AUTH_SERVICE_URL=http://localhost:8002
```

### 3. Run the server
```bash
go run main.go
```

---

## 🧪 Example Endpoints

- `GET /health` → Health check
- `POST /auth/login` → Routed to auth service
- `POST /cv/generate` → Routed to CV service

---

## 📜 License

[MIT License](LICENSE)

---

## 👨‍💻 Maintainer

Built and maintained by **Tangwe Caleb Ojang** with 💙
