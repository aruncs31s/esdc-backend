#!/bin/bash

echo "Generating Swagger documentation..."
swag init

if [ $? -eq 0 ]; then
    echo "✅ Swagger documentation generated successfully!"
    echo "📖 Access the documentation at: http://localhost:9090/swagger/index.html"
else
    echo "❌ Failed to generate Swagger documentation"
    exit 1
fi