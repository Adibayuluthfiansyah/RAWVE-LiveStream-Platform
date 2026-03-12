<div align="center">

<img src="https://readme-typing-svg.demolab.com?font=Bebas+Neue&size=80&pause=1000&color=FF2D4E&center=true&vCenter=true&width=600&height=120&lines=RAWVE" alt="RAWVE" />

### **Your Stream. Your Rules. Your Revenue.**
*Freedom without limits — but never without responsibility.*

<br/>

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-FF2D4E?style=for-the-badge)](LICENSE)
[![Architecture](https://img.shields.io/badge/Architecture-Clean-FF6225?style=for-the-badge&logo=buffer&logoColor=white)](#architecture)
[![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-00C853?style=for-the-badge&logo=github&logoColor=white)](CONTRIBUTING.md)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://postgresql.org)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://docker.com)

<br/>

> RAWVE is an **open-source live streaming backend** engineered for extreme performance,
> massive scalability, and creator-first fairness. We believe creators deserve
> full control over their community and revenue — backed by a moderation system
> that keeps the ecosystem healthy and safe for everyone.

<br/>

[**Get Started**](#-getting-started) · [**Architecture**](#-architecture) · [**Roadmap**](#-roadmap) · [**Contributing**](#-contributing)

</div>

---

## ✨ Why RAWVE?

Most streaming platforms take your audience, clip your revenue, and decide your rules.

**RAWVE flips that model.**

| Platform | Revenue Cut | Creator Control | Open Source |
|---|---|---|---|
| Twitch | ~50% | Limited | ❌ |
| YouTube Live | ~45% | Limited | ❌ |
| TikTok Live | ~70%+ | Minimal | ❌ |
| **RAWVE** | **0% platform cut** | **Full** | **✅** |

---

## 🚀 Roadmap

| Status | Feature | Description |
|---|---|---|
| ✅ | **Secure Authentication** | Webhook integration & JWT verification via Clerk |
| ✅ | **Clean Architecture** | Strict Domain / Usecase / Repository / Delivery separation |
| 🔧 | **Real-time Live Chat** | Lightning-fast messaging with Gorilla WebSocket |
| 🔧 | **Stream Management** | Creator live stage (Stream) creation and control |
| 📋 | **Fair Monetization** | Payment gateway with zero platform cuts |
| 📋 | **Creator Analytics** | Real-time viewers, revenue, and engagement dashboard |
| 📋 | **Content Moderation** | Community-driven + automated moderation system |

> `✅ Done` · `🔧 In Progress` · `📋 Planned`

---

## 🛠️ Tech Stack

```
Language        →  Golang (Go 1.25+)
HTTP Framework  →  Gin Web Framework
WebSocket       →  Gorilla WebSocket
ORM & Database  →  GORM + PostgreSQL
Authentication  →  Clerk SDK (JWT Verification)
Infrastructure  →  Docker & Docker Compose
```

---

## 📁 Architecture

RAWVE is built with **Clean Architecture** — strict separation of concerns, zero dependency leakage between layers.

```
rawve/
├── cmd/
│   └── api/
│       └── main.go              # Entry point & Dependency Injection
│
├── internal/
│   ├── domain/                  # ♥  Heart of the app
│   │   ├── user.go              #    Entities & Interface contracts
│   │   └── stream.go
│   │
│   ├── usecase/                 # 🧠 Brain of the app
│   │   ├── user_usecase.go      #    Core business logic
│   │   └── stream_usecase.go
│   │
│   ├── repository/              # 💪 Muscle of the app
│   │   ├── user_repository.go   #    Database & external service calls
│   │   └── stream_repository.go
│   │
│   └── delivery/                # 🚪 Receptionist of the app
│       ├── http/                #    HTTP handlers & routing
│       ├── websocket/           #    WebSocket handlers
│       └── middleware/          #    Auth, logging, rate limiting
│
├── docker-compose.yml
├── .env.example
└── README.md
```

### Data Flow

```
HTTP Request
     │
     ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Delivery  │────▶│   Usecase   │────▶│ Repository  │────▶│  Database   │
│  (Gin/WS)   │     │ (Business)  │     │  (Data)     │     │ (Postgres)  │
└─────────────┘     └─────────────┘     └─────────────┘     └─────────────┘
     ▲                    │                    │
     │                    ▼                    ▼
     │             ┌─────────────┐     ┌─────────────┐
     └─────────────│   Domain    │     │    Clerk    │
      Response     │  (Entities) │     │   (Auth)    │
                   └─────────────┘     └─────────────┘
```

---

## ⚡ Getting Started

### Prerequisites

Make sure you have these installed:

- [Go 1.25+](https://golang.org/doc/install)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- A [Clerk](https://clerk.com) account (free tier works)

### 1. Clone the Repository

```bash
git clone [https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform.git](https://github.com/Adibayuluthfiansyah/RAWVE-LiveStream-Platform.git)
cd RAWVE-LiveStream-Platform
```

### 2. Configure Environment

```bash
cp .env.example .env
```

Open `.env` and fill in your values:

```env
# Server
PORT=8080

# Database
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=rawve
DB_PORT=5435

# Clerk Authentication
CLERK_SECRET_KEY=sk_test_your_clerk_secret_key_here
```

> 🔑 Get your `CLERK_SECRET_KEY` from [Clerk Dashboard](https://dashboard.clerk.com) → API Keys

### 3. Start the Database

```bash
docker-compose up -d
```

### 4. Run the Application

```bash
go run cmd/api/main.go
```

Server will be running at `http://localhost:8080` 🚀

---

## 🔐 Authentication Flow

RAWVE uses **Clerk** for authentication with a webhook-based user sync pattern:

```
User signs in via Clerk
        │
        ▼
Clerk issues JWT token
        │
        ▼
Frontend sends request with Bearer token
        │
        ▼
RAWVE middleware verifies JWT (Clerk public key)
        │
        ▼
Extract claims → SyncUserFromAuth()
        │
        ▼
CreateOrUpdate user in PostgreSQL
        │
        ▼
Request continues with authenticated context
```

---

## 🤝 Contributing

We welcome contributions from everyone — bug fixes, new features, or documentation improvements.

```bash
# 1. Fork this repository

# 2. Create your feature branch
git checkout -b feature/your-awesome-feature

# 3. Commit your changes
git commit -m "feat: add your awesome feature"

# 4. Push to the branch
git push origin feature/your-awesome-feature

# 5. Open a Pull Request
```

### Commit Convention

We follow [Conventional Commits](https://www.conventionalcommits.org/):

| Prefix | Use for |
|---|---|
| `feat:` | New feature |
| `fix:` | Bug fix |
| `docs:` | Documentation only |
| `refactor:` | Code refactoring |
| `test:` | Adding tests |
| `chore:` | Build process or tooling |

---

## 👨‍💻 Creator

<div align="center">

<br/>

<img src="https://github.com/adibayuluthfiansyah.png" width="100" height="100" style="border-radius:50%" alt="Adibayu Luthfiansyah"/>

### **Adibayu Luthfiansyah**

*Founder & Lead Developer of RAWVE*

[![Website](https://img.shields.io/badge/Website-adibayuluthfiansyah.dev-FF2D4E?style=for-the-badge&logo=safari&logoColor=white)](https://adibayuluthfiansyah.dev)
[![GitHub](https://img.shields.io/badge/GitHub-adibayuluthfiansyah-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/adibayuluthfiansyah)

<br/>

> *"I built RAWVE because I believe the internet needs a streaming platform*
> *that genuinely puts creators first — not as a product, but as owners."*

<br/>

</div>

---

## 📄 License

RAWVE is open-source software licensed under the [MIT License](LICENSE).

---

<div align="center">

Built with ❤️ for creators who deserve better.

**[rawve.live](https://rawve.live)** · [@rawve](https://github.com/rawve)

<br/>

*"Your Stream. Your Rules. Your Revenue."*

</div>