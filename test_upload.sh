#!/bin/bash

# File Upload API Test Script
# This script demonstrates how to test the file upload endpoints

BASE_URL="http://localhost:8080"

echo "================================"
echo "File Upload API Test"
echo "================================"
echo ""

# Test 1: Upload a single image
echo "Test 1: Uploading a single image..."
echo "Command: curl -X POST $BASE_URL/api/files/upload/image -F \"image=@test_image.jpg\""
echo ""
echo "Expected Response:"
echo '{
  "status": "success",
  "data": {
    "message": "Image uploaded successfully",
    "path": "images/20251004173045_a8b3c5d1.jpg",
    "url": "/uploads/images/20251004173045_a8b3c5d1.jpg"
  }
}'
echo ""
echo "--------------------------------"
echo ""

# Test 2: Upload a single file
echo "Test 2: Uploading a single file..."
echo "Command: curl -X POST $BASE_URL/api/files/upload -F \"file=@test_document.pdf\""
echo ""
echo "Expected Response:"
echo '{
  "status": "success",
  "data": {
    "message": "File uploaded successfully",
    "path": "files/20251004173045_a8b3c5d1.pdf",
    "url": "/uploads/files/20251004173045_a8b3c5d1.pdf",
    "filename": "test_document.pdf",
    "size": "1.2 MB"
  }
}'
echo ""
echo "--------------------------------"
echo ""

# Test 3: Upload multiple files
echo "Test 3: Uploading multiple files..."
echo "Command: curl -X POST $BASE_URL/api/files/upload/multiple -F \"files=@file1.pdf\" -F \"files=@file2.jpg\""
echo ""
echo "Expected Response:"
echo '{
  "status": "success",
  "data": {
    "message": "Uploaded 2 of 2 files",
    "uploaded": [...],
    "failed": [],
    "uploaded_count": 2,
    "failed_count": 0
  }
}'
echo ""
echo "--------------------------------"
echo ""

# Test 4: Upload with custom directory
echo "Test 4: Uploading to custom directory..."
echo "Command: curl -X POST \"$BASE_URL/api/files/upload?dir=documents\" -F \"file=@test.pdf\""
echo ""
echo "--------------------------------"
echo ""

# Test 5: Access uploaded file
echo "Test 5: Accessing uploaded file..."
echo "Command: curl $BASE_URL/uploads/images/20251004173045_a8b3c5d1.jpg -o downloaded_image.jpg"
echo ""
echo "--------------------------------"
echo ""

echo "To run actual tests, replace the placeholder filenames with real files."
echo ""
echo "Quick Test Example:"
echo "  # Create a test image"
echo "  echo 'test' > test.txt"
echo "  curl -X POST $BASE_URL/api/files/upload -F \"file=@test.txt\""
echo ""
