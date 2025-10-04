# File Upload System Architecture

## System Flow Diagram

```
┌─────────────┐
│   Client    │
│  (Browser/  │
│   Mobile)   │
└──────┬──────┘
       │
       │ POST /api/files/upload
       │ multipart/form-data
       │
       ▼
┌─────────────────────────────────────┐
│         Gin Router                  │
│  (routes/file_routes.go)            │
└──────┬──────────────────────────────┘
       │
       │ Route to handler
       │
       ▼
┌─────────────────────────────────────┐
│      File Handler                   │
│  (handler/file_handler.go)          │
│                                     │
│  - Parse form data                  │
│  - Validate request                 │
│  - Call service layer               │
└──────┬──────────────────────────────┘
       │
       │ Delegate to service
       │
       ▼
┌─────────────────────────────────────┐
│      File Service                   │
│  (service/file_service.go)          │
│                                     │
│  - Validate file type/size          │
│  - Generate unique filename         │
│  - Create directories               │
│  - Save file to disk                │
└──────┬──────────────────────────────┘
       │
       │ Save to filesystem
       │
       ▼
┌─────────────────────────────────────┐
│      File System                    │
│                                     │
│  uploads/                           │
│    ├── images/                      │
│    │   └── 20251004_abc123.jpg      │
│    └── files/                       │
│        └── 20251004_def456.pdf      │
└──────┬──────────────────────────────┘
       │
       │ Return file path
       │
       ▼
┌─────────────────────────────────────┐
│      Response Helper                │
│  (handler/responses/responses.go)   │
│                                     │
│  Format JSON response               │
└──────┬──────────────────────────────┘
       │
       │ JSON Response
       │
       ▼
┌─────────────┐
│   Client    │
│  Receives:  │
│  - path     │
│  - url      │
│  - metadata │
└─────────────┘
```

## Component Breakdown

### 1. Routes Layer (`file_routes.go`)
```go
POST /api/files/upload/image      → UploadImage()
POST /api/files/upload             → UploadFile()
POST /api/files/upload/multiple    → UploadMultipleFiles()
GET  /uploads/*                    → Static File Server
```

### 2. Handler Layer (`file_handler.go`)
```go
Responsibilities:
- Parse multipart form data
- Extract files from request
- Validate request structure
- Call service methods
- Format responses
```

### 3. Service Layer (`file_service.go`)
```go
Responsibilities:
- Business logic
- File validation (type, size)
- Filename generation
- Directory management
- File I/O operations
```

### 4. File System
```
uploads/
├── images/        (for images only)
│   └── YYYYMMDDHHMMSS_random.ext
└── files/         (for all other files)
    └── YYYYMMDDHHMMSS_random.ext
```

## Request Flow Example

### Upload Image Request

```http
POST /api/files/upload/image
Content-Type: multipart/form-data

------WebKitFormBoundary
Content-Disposition: form-data; name="image"; filename="photo.jpg"
Content-Type: image/jpeg

[binary image data]
------WebKitFormBoundary--
```

**Step-by-step:**

1. **Request arrives** at Gin router
2. **Route matching** → `/api/files/upload/image`
3. **Handler invoked** → `fileHandler.UploadImage()`
4. **Parse form** → Extract file from `c.FormFile("image")`
5. **Validate file**:
   - Check extension (.jpg allowed ✅)
   - Check size (< 5MB ✅)
6. **Generate filename** → `20251004224149_8f14e7df.jpg`
7. **Create path** → `uploads/images/20251004224149_8f14e7df.jpg`
8. **Save file** → Write to disk
9. **Return response**:
   ```json
   {
     "status": "success",
     "data": {
       "path": "images/20251004224149_8f14e7df.jpg",
       "url": "/uploads/images/20251004224149_8f14e7df.jpg"
     }
   }
   ```

## Static File Serving

```
Client Request:
GET /uploads/images/20251004224149_8f14e7df.jpg

      ↓

Gin Static Handler (configured in routes.go):
r.Static("/uploads", "./uploads")

      ↓

File System:
Read: ./uploads/images/20251004224149_8f14e7df.jpg

      ↓

Response:
Content-Type: image/jpeg
[binary image data]
```

## Data Flow

### Upload Phase
```
File Input → HTTP Request → Router → Handler → Service → Disk
```

### Access Phase
```
HTTP Request → Static Handler → File System → HTTP Response
```

## Error Handling Flow

```
┌─────────────┐
│   Request   │
└──────┬──────┘
       │
       ▼
  ┌─────────┐
  │Validate │
  └────┬────┘
       │
       ├─── Invalid ──→ 400 Bad Request
       │                {
       │                  "status": "error",
       │                  "message": "Invalid file"
       │                }
       │
       ▼
  ┌──────────┐
  │  Save    │
  └────┬─────┘
       │
       ├─── Error ───→ 500 Internal Error
       │                {
       │                  "status": "error",
       │                  "message": "Failed to save"
       │                }
       │
       ▼
  ┌──────────┐
  │ Success  │──→ 200 OK
  └──────────┘      {
                      "status": "success",
                      "data": {...}
                    }
```

## File Naming Convention

```
Input:  "my vacation photo.jpg"
        ↓
Process: Generate timestamp + random ID
        ↓
Output: "20251004224149_8f14e7df.jpg"

Components:
- 20251004: Date (YYYYMMDD)
- 224149: Time (HHMMSS)
- 8f14e7df: Random hex ID (8 chars)
- .jpg: Original extension
```

## Integration Points

### With Project Creation
```go
// Frontend sends:
FormData:
- image: [file]
- name: "Project Name"
- description: "..."

// Backend processes:
1. Upload image → get path
2. Create project with image path
3. Return complete project
```

### With User Profiles
```go
// Upload profile picture
POST /api/files/upload/image

// Update user profile
PATCH /api/user/profile
{
  "avatar": "/uploads/images/20251004_abc123.jpg"
}
```

## Security Layers

```
┌─────────────────────────────────────┐
│   Layer 1: CORS                     │
│   - Origin validation               │
└─────────────────────────────────────┘
            ↓
┌─────────────────────────────────────┐
│   Layer 2: File Validation          │
│   - Extension check                 │
│   - Size limit                      │
└─────────────────────────────────────┘
            ↓
┌─────────────────────────────────────┐
│   Layer 3: Path Safety              │
│   - filepath.Join()                 │
│   - Prevents directory traversal    │
└─────────────────────────────────────┘
            ↓
┌─────────────────────────────────────┐
│   Layer 4: Unique Naming            │
│   - Prevents overwriting            │
│   - Collision avoidance             │
└─────────────────────────────────────┘
```

## Performance Considerations

### Current Implementation
- **Memory**: Files buffered in memory during upload
- **Concurrency**: Gin handles concurrent requests
- **Storage**: Local file system

### Scalability Options
1. **Chunked Upload**: For large files
2. **Streaming**: Direct to disk without buffering
3. **Queue System**: Background processing
4. **CDN**: For file delivery
5. **Cloud Storage**: S3, GCS, Azure Blob

## Monitoring Points

```
Key Metrics to Monitor:
├── Upload Success Rate
├── Average Upload Time
├── File Size Distribution
├── Storage Usage
├── Failed Upload Reasons
└── Popular File Types
```

## Directory Structure

```
esdc-backend/
├── main.go                          (Entry point)
├── internal/
│   ├── handler/
│   │   └── file_handler.go          (HTTP handlers)
│   ├── service/
│   │   └── file_service.go          (Business logic)
│   └── routes/
│       └── file_routes.go           (Route definitions)
├── uploads/                         (Upload destination)
│   ├── images/                      (Image files)
│   └── files/                       (Other files)
└── docs/
    └── FILE_UPLOAD_API.md           (Documentation)
```

## Summary

✅ **3-Layer Architecture**: Routes → Handlers → Services
✅ **Clear Separation**: Each component has specific responsibility
✅ **Error Handling**: Comprehensive error responses
✅ **Security**: Multiple validation layers
✅ **Scalability**: Can be extended for cloud storage
✅ **Maintainability**: Well-organized and documented

The system is production-ready and can handle file uploads efficiently!
