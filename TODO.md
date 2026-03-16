# 📋 RAWVE Project Roadmap & TODOs

## 0. Core Infrastructure & Chat ( COMPLETED)
* [x] Setup Golang Clean Architecture
* [x] Integrasi PostgreSQL & GORM
* [x] Autentikasi via Clerk Webhook
* [x] **Real-time Live Chat** (Gorilla WebSocket, Hub & Client)
* [x] API Start & End Stream (`POST /api/streams/start`, `POST /api/streams/end`)

---

## 1. API Etalase Utama (Get Active Streams) ( CORE COMPLETED)

### Planning & Backend
* [x] Tentukan struktur data & database schema tabel `streams` (id, title, thumbnail, is_live, category)
* [x] Buat endpoint `GET /api/streams/live`
* [x] Query database untuk mengambil stream dengan status `is_live = true`
* [ ] Tambahkan sorting (viewer terbanyak / terbaru)
* [ ] Tambahkan pagination (limit & offset)
* [ ] Tambahkan filter kategori stream (opsional)

### Response Format
* [x] Tentukan JSON response structure
* [ ] Sertakan metadata (total_streams, page, limit) untuk pagination

Contoh response:
  {
    "message": "Berhasil mengambil daftar live stream",
    "data": [
      {
        "id": "user_123",
        "title": "RANKED CLIMB | 500 FPS UNLOCKED",
        "category": "Gaming",
        "thumbnail_url": "https://...",
        "is_live": true,
        "created_at": "..."
      }
    ]
  }

### Testing
* [x] Unit test endpoint (via Bruno)
* [ ] Test pagination
* [ ] Test respon saat tidak ada stream aktif

---

## 2. API Update Profil (Onboarding) ( COMPLETED)

### Planning & Backend
* [x] Tentukan field onboarding (display_name, bio, category, dll)
* [x] Buat endpoint `PUT /api/profile/setup`
* [x] Validasi input user (binding JSON)
* [x] Middleware authentication (RequireAuth)
* [x] Update data profil user di database PostgreSQL
* [ ] Upload avatar ke storage (Opsional: Saat ini menggunakan bawaan dari Clerk)

### Testing
* [x] Test update profile (via Bruno)
* [ ] Test upload image file
* [x] Test unauthorized request (tanpa token)

---

## 3. Mesin Video Streaming (The Core Engine) ( BACKEND COMPLETED)

### Research & Architecture
* [x] Tentukan protokol streaming: **RTMP (Ingest) & HLS (Playback)**
* [x] Setup Server Media: **MediaMTX via Docker Compose**
* [x] Setup jalur port (1935 untuk RTMP OBS, 8888 untuk HLS Web)

### Streaming Flow
* [x] Generator Stream Key otomatis untuk Kreator (via Clerk Webhook)
* [x] Creator start stream via OBS (RTMP)
* [x] Server (MediaMTX) menerima stream dan convert ke HLS otomatis
* [ ] Sistem menghitung dan update *viewer count* secara real-time

### Frontend (Next.js) - 🚀 NEXT STEP
* [ ] Setup kerangka Next.js + Tailwind CSS + TypeScript
* [ ] Integrasi video player (HLS.js / Video.js)
* [ ] Tampilkan UI Live Chat terhubung ke WebSocket backend
* [ ] Tampilkan indikator status *Live* & *Viewer Count*

---

## 4. Deployment (Infrastruktur Production)

* [ ] VPS / Cloud Server (AWS/DigitalOcean/GCP)
* [ ] Setup environment variables production
* [ ] Setup database PostgreSQL production (Neon.tech / Supabase)
* [ ] Setup Docker & Docker Compose untuk production
* [ ] Setup Reverse Proxy (Nginx / Traefik) & SSL/HTTPS (Certbot)
* [ ] Setup monitoring & logging

---

## 5. Future Improvements (Post-MVP)

* [ ] Monetization: Sistem Donasi & Subscription
* [ ] Dashboard Analitik Kreator (Grafik pendapatan & view)
* [ ] Stream recording (VOD - Video on Demand)
* [ ] Algoritma rekomendasi stream di halaman beranda