#!/bin/bash

echo "Starting ESDC Backend with Chat Support..."
echo "WebSocket endpoint: ws://localhost:9090/ws/chat"
echo "REST endpoint: http://localhost:9090/api/chat/messages"
echo ""

go run main.go
