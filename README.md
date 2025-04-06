# 🚪 Gateway Service – ProCV API Gateway

This is the **Gateway microservice** for the ProCV platform, designed to act as a centralized entry point to route requests and manage configuration across internal services.

---

## 📦 Structure

This service was scaffolded with simplicity and flexibility in mind.

### Current Structure:
```bash
gateway-service/
├── config/
│   └── config.yaml         # Central config for the gateway
├── main.go                 # Application entry point
├── go.mod                  # Module definition
└── README.md
```

---

## ⚙️ Configuration

The gateway uses a single configuration file to manage service endpoints, port settings, and environment-specific flags.

### Example `config.yaml`:
```yaml
server:
  port: 8080

services:
  profile: http://localhost:5001
  cv: http://localhost:5002
  auth: http://localhost:5003
```

---

## 🚀 Usage

### 1. Clone the repo
```bash
git clone https://github.com/your-org/procv-gateway.git
cd procv-gateway
```

### 2. Update Configuration
Edit `config/config.yaml` to match your local or deployed service URLs.

### 3. Run the Gateway
```bash
go run main.go
```

---

## 📌 Notes
- This service currently supports static routing via config.
- No embedded auth or middleware yet – this can be added in later iterations.

---

## 📜 License
[MIT License](LICENSE)

---

## 👨‍💻 Author
Built by **Tangwe Caleb Ojang**
