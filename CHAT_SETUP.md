# Chat Feature Setup

## Backend (Go)

The chat feature uses WebSocket for real-time communication.

### Files Created:
- `internal/model/message.go` - Message model
- `internal/handler/chat_handler.go` - WebSocket handler with mock data
- `internal/routes/chat_routes.go` - Chat routes

### Endpoints:
- `ws://localhost:9090/ws/chat` - WebSocket endpoint
- `GET /api/chat/messages` - Get message history (REST)

### Mock Data:
The backend includes 2 mock messages on startup:
- Alice: "Hey everyone! Welcome to the chat!"
- Bob: "Thanks! Excited to be here."

### Running:
```bash
cd esdc-backend
go run main.go
```

## Frontend (React)

### Files Created:
- `src/components/Chatroom.tsx` - Chat component
- `src/styles/chatroom.css` - Chat styles

### Features:
- Real-time messaging via WebSocket
- Fallback to mock mode if WebSocket fails
- Auto-scroll to latest messages
- Responsive design

### Environment:
Add to `.env`:
```
VITE_WS_URL=ws://localhost:9090/ws/chat
```

### Usage:
Click the "Chat" button in the navbar to open the chatroom.
