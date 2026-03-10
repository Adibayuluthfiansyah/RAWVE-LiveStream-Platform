# Go-LiveChat

Go-LiveChat adalah backend layanan pesan waktu nyata (real-time) yang dibangun menggunakan Golang dan WebSockets. Proyek ini memfasilitasi komunikasi dua arah yang cepat dan efisien antara klien dan server.

## Tech Stack
* **Language:** Golang (Go)
* **WebSocket:** [Gorilla WebSocket](https://github.com/gorilla/websocket)
* **Architecture:** Standard Go Project Layout

## Project Structure
```text
.
├── cmd/
│   └── api/
│       └── main.go           # Entry point aplikasi
├── internal/
│   ├── config/               # Konfigurasi aplikasi (ENV, dll)
│   ├── handlers/             # HTTP Handlers
│   ├── models/               # Struktur data & entitas
│   ├── services/             # Logika bisnis
│   └── websockets/           # Hub, Client, dan routing pesan WebSocket
├── pkg/
│   └── utils/                # Fungsi helper/utilitas
├── go.mod                    # Go module file
├── go.sum                    # Go dependencies checksum
└── README.md