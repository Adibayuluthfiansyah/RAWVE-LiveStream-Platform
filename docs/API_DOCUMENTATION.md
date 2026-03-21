# RAWVE Livestream API Documentation

> **API Version:** 1.0  
> **Base URL:** `http://localhost:8080/api`  
> **Swagger UI:** `http://localhost:8080/swagger/index.html`

## 📚 Table of Contents

1. [Authentication](#authentication)
2. [System Endpoints](#system-endpoints)
3. [Stream Management](#stream-management)
4. [User Management](#user-management)
5. [Webhooks](#webhooks)
6. [WebSocket](#websocket)
7. [Error Responses](#error-responses)

---

## 🔐 Authentication

API ini menggunakan **Bearer Token Authentication** dengan JWT dari Clerk.

### Header Format

```http
Authorization: Bearer <your_jwt_token>
```

### Protected Endpoints

Endpoint yang memerlukan authentication ditandai dengan 🔒.

---

## 🔧 System Endpoints

### Health Check

Check if the server is running.

**Endpoint:** `GET /api/ping`  
**Auth:** Public ✅

#### Request

```bash
curl -X GET http://localhost:8080/api/ping
```

#### Response (200 OK)

```json
{
  "message": "Server RAWVE is running!"
}
```

---

## 🎥 Stream Management

### 1. Get All Active Livestreams

Retrieve a list of all currently active livestreams.

**Endpoint:** `GET /api/streams/live`  
**Auth:** Public ✅

#### Request

```bash
curl -X GET http://localhost:8080/api/streams/live
```

#### Response (200 OK)

```json
{
  "message": "Success get list livestream",
  "data": [
    {
      "id": "user_12345",
      "title": "Gaming Session",
      "category": "Gaming",
      "thumbnail_url": "https://example.com/thumbnail.jpg",
      "is_live": true,
      "enable_donation": true,
      "followers_only_chat": false,
      "created_at": "2026-03-21T10:00:00Z",
      "updated_at": "2026-03-21T10:00:00Z"
    }
  ]
}
```

#### Response (500 Internal Server Error)

```json
{
  "error": "Failed to fetch stream list"
}
```

---

### 2. Start a Livestream 🔒

Create and start a new livestream session.

**Endpoint:** `POST /api/streams/start`  
**Auth:** Bearer Token Required 🔒

#### Request Headers

```http
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

#### Request Body

```json
{
  "title": "My First Stream",
  "category": "Music",
  "thumbnail_url": "https://example.com/thumb.jpg"
}
```

**Field Descriptions:**
- `title` (string, **required**): Stream title
- `category` (string, optional): Stream category
- `thumbnail_url` (string, optional): URL to stream thumbnail

#### Request Example

```bash
curl -X POST http://localhost:8080/api/streams/start \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Stream",
    "category": "Music"
  }'
```

#### Response (200 OK)

```json
{
  "message": "Stream succes, your stream now !",
  "data": {
    "id": "user_12345",
    "title": "My First Stream",
    "category": "Music",
    "thumbnail_url": "https://example.com/thumb.jpg",
    "is_live": true,
    "enable_donation": true,
    "followers_only_chat": false,
    "created_at": "2026-03-21T10:00:00Z",
    "updated_at": "2026-03-21T10:00:00Z"
  }
}
```

#### Response (400 Bad Request)

```json
{
  "error": "Wrong format or title empty"
}
```

#### Response (401 Unauthorized)

```json
{
  "error": "user unathorized"
}
```

#### Response (500 Internal Server Error)

```json
{
  "error": "Failed to start stream"
}
```

---

### 3. End a Livestream 🔒

Stop the current livestream session.

**Endpoint:** `POST /api/streams/end`  
**Auth:** Bearer Token Required 🔒

#### Request Headers

```http
Authorization: Bearer <your_jwt_token>
```

#### Request Example

```bash
curl -X POST http://localhost:8080/api/streams/end \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### Response (200 OK)

```json
{
  "message": "Stream ended"
}
```

#### Response (401 Unauthorized)

```json
{
  "error": "Unauthorized"
}
```

#### Response (500 Internal Server Error)

```json
{
  "error": "Failed to end stream"
}
```

---

## 👤 User Management

### 1. Setup User Profile 🔒

Update user profile with display name, bio, and category.

**Endpoint:** `PUT /api/profile/setup`  
**Auth:** Bearer Token Required 🔒

#### Request Headers

```http
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

#### Request Body

```json
{
  "display_name": "John Streamer",
  "bio": "Professional gamer and content creator",
  "category": "Gaming"
}
```

**Field Descriptions:**
- `display_name` (string, **required**): User display name
- `bio` (string, optional): User biography
- `category` (string, optional): Primary content category

#### Request Example

```bash
curl -X PUT http://localhost:8080/api/profile/setup \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "John Streamer",
    "bio": "Professional gamer",
    "category": "Gaming"
  }'
```

#### Response (200 OK)

```json
{
  "message": "Profile updated successfully, WELCOME TO RAWVE"
}
```

#### Response (400 Bad Request)

```json
{
  "error": "Wrong format or Display Name Empty"
}
```

#### Response (401 Unauthorized)

```json
{
  "error": "Unathorized Access"
}
```

#### Response (500 Internal Server Error)

```json
{
  "error": "Failed to update profile"
}
```

---

### 2. Creator Dashboard 🔒

Access to creator studio dashboard.

**Endpoint:** `GET /api/dashboard`  
**Auth:** Bearer Token Required 🔒

#### Request Headers

```http
Authorization: Bearer <your_jwt_token>
```

#### Request Example

```bash
curl -X GET http://localhost:8080/api/dashboard \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### Response (200 OK)

```json
{
  "message": "Welcome to Creator Studio RAWVE!",
  "user_id": "user_12345",
  "status": "success"
}
```

#### Response (401 Unauthorized)

```json
{
  "error": "Unauthorized"
}
```

---

## 🔔 Webhooks

### Clerk Webhook Handler

Receives user lifecycle events from Clerk authentication service.

**Endpoint:** `POST /api/webhooks/clerk`  
**Auth:** Clerk Signature (not Bearer Token)

> **Note:** This endpoint is designed to be called by Clerk servers, not directly by clients.

#### Request Body Example (User Created)

```json
{
  "type": "user.created",
  "data": {
    "id": "user_12345",
    "username": "johndoe",
    "email_address": "john@example.com",
    "image_url": "https://example.com/avatar.jpg"
  }
}
```

#### Request Body Example (User Updated)

```json
{
  "type": "user.updated",
  "data": {
    "id": "user_12345",
    "username": "johndoe",
    "email_address": "john@example.com",
    "image_url": "https://example.com/avatar.jpg"
  }
}
```

#### Response (200 OK)

```json
{
  "message": "webhook successfully processed"
}
```

#### Response (400 Bad Request)

```json
{
  "Error": "Invalid payload format"
}
```

#### Response (500 Internal Server Error)

```json
{
  "error": "Failed to syncron user from auth"
}
```

---

## 💬 WebSocket

### Live Chat Connection

Connect to live chat for a specific stream.

**Endpoint:** `ws://localhost:8080/api/ws/chat/:stream_id`  
**Protocol:** WebSocket

#### Connection Example (JavaScript)

```javascript
const streamId = "user_12345";
const ws = new WebSocket(`ws://localhost:8080/api/ws/chat/${streamId}`);

ws.onopen = () => {
  console.log("Connected to chat");
};

ws.onmessage = (event) => {
  const message = JSON.parse(event.data);
  console.log("New message:", message);
};

ws.onerror = (error) => {
  console.error("WebSocket error:", error);
};

ws.onclose = () => {
  console.log("Disconnected from chat");
};

// Send a message
ws.send(JSON.stringify({
  content: "Hello everyone!"
}));
```

#### Message Format (Received)

```json
{
  "id": 123,
  "stream_id": "user_12345",
  "user_id": "user_67890",
  "content": "Hello everyone!",
  "created_at": "2026-03-21T10:00:00Z"
}
```

---

## ❌ Error Responses

### Standard Error Format

All errors follow this format:

```json
{
  "error": "Error message description"
}
```

### Common HTTP Status Codes

| Code | Meaning | Description |
|------|---------|-------------|
| 200 | OK | Request successful |
| 400 | Bad Request | Invalid request format or missing required fields |
| 401 | Unauthorized | Missing or invalid authentication token |
| 403 | Forbidden | Valid token but insufficient permissions |
| 404 | Not Found | Resource not found |
| 500 | Internal Server Error | Server-side error |

---

## 🚀 Quick Start for Frontend Development

### 1. Environment Setup

Set your backend URL in your frontend `.env`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
NEXT_PUBLIC_WS_URL=ws://localhost:8080/api
```

### 2. API Client Example (Next.js)

```typescript
// lib/api-client.ts
import { useAuth } from "@clerk/nextjs";

const API_URL = process.env.NEXT_PUBLIC_API_URL;

export function useApiClient() {
  const { getToken } = useAuth();

  async function apiCall(endpoint: string, options: RequestInit = {}) {
    const token = await getToken();
    
    const response = await fetch(`${API_URL}${endpoint}`, {
      ...options,
      headers: {
        "Content-Type": "application/json",
        ...(token && { Authorization: `Bearer ${token}` }),
        ...options.headers,
      },
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || "API Error");
    }

    return response.json();
  }

  return { apiCall };
}
```

### 3. Usage Example

```typescript
// components/StartStreamButton.tsx
"use client";

import { useApiClient } from "@/lib/api-client";

export function StartStreamButton() {
  const { apiCall } = useApiClient();

  async function startStream() {
    try {
      const data = await apiCall("/streams/start", {
        method: "POST",
        body: JSON.stringify({
          title: "My Stream",
          category: "Gaming",
        }),
      });
      console.log("Stream started:", data);
    } catch (error) {
      console.error("Failed to start stream:", error);
    }
  }

  return <button onClick={startStream}>Start Stream</button>;
}
```

### 4. WebSocket Hook Example

```typescript
// hooks/use-chat.ts
import { useEffect, useState, useRef } from "react";

const WS_URL = process.env.NEXT_PUBLIC_WS_URL;

export function useChat(streamId: string) {
  const [messages, setMessages] = useState<any[]>([]);
  const wsRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const ws = new WebSocket(`${WS_URL}/ws/chat/${streamId}`);
    wsRef.current = ws;

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      setMessages((prev) => [...prev, message]);
    };

    return () => ws.close();
  }, [streamId]);

  function sendMessage(content: string) {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      wsRef.current.send(JSON.stringify({ content }));
    }
  }

  return { messages, sendMessage };
}
```

---

## 📊 Rate Limiting

The API has rate limiting enabled on all `/api/*` endpoints.

**Current Limits:**
- Default: Check with backend team for specific limits
- Rate limit headers are included in responses:
  - `X-RateLimit-Limit`
  - `X-RateLimit-Remaining`
  - `X-RateLimit-Reset`

---

## 🔒 Security Notes

### Development Mode

⚠️ **IMPORTANT:** The API is currently running in development mode with:

1. **CORS:** `AllowAllOrigins = true`
   - In production, this will be restricted to specific domains
   
2. **Auth Middleware:** Currently using hardcoded `user_id = "Adibayu"`
   - JWT verification is commented out for development
   - Will be enabled in production

### Production Checklist

Before deploying to production:

- [ ] Enable Clerk JWT verification in `auth_middleware.go`
- [ ] Configure CORS with specific allowed origins
- [ ] Set `GIN_MODE=release`
- [ ] Disable Swagger UI in production (or add auth)
- [ ] Configure proper rate limiting
- [ ] Enable HTTPS only
- [ ] Add request logging
- [ ] Setup monitoring and alerts

---

## 🛠️ Testing with Swagger UI

Access the interactive API documentation at:

```
http://localhost:8080/swagger/index.html
```

### Steps to test authenticated endpoints:

1. Click the **"Authorize"** button (top right)
2. Enter your token in format: `Bearer YOUR_JWT_TOKEN`
3. Click **"Authorize"** and close the dialog
4. Try any protected endpoint

---

## 📝 Notes

- All timestamps are in ISO 8601 format (UTC)
- Request/Response bodies use `application/json` content type
- Maximum request body size: 16MB (configurable)
- WebSocket connections have automatic ping/pong for keepalive

---

## 🆘 Support

For issues or questions:
- Email: support@rawve.com
- Check server logs for detailed error information
- Review Swagger UI for live API specs

---

**Last Updated:** March 21, 2026  
**API Version:** 1.0
