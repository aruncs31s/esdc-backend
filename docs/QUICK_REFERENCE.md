# File Upload Quick Reference

## API Endpoints

| Endpoint | Method | Purpose | Max Size |
|----------|--------|---------|----------|
| `/api/files/upload/image` | POST | Upload single image | 5MB |
| `/api/files/upload` | POST | Upload single file | 10MB |
| `/api/files/upload/multiple` | POST | Upload multiple files | 10MB each |
| `/uploads/{path}` | GET | Access uploaded file | - |

## cURL Examples

```bash
# Upload image
curl -X POST http://localhost:8080/api/files/upload/image -F "image=@photo.jpg"

# Upload file
curl -X POST http://localhost:8080/api/files/upload -F "file=@document.pdf"

# Upload to custom directory
curl -X POST "http://localhost:8080/api/files/upload?dir=docs" -F "file=@file.pdf"

# Multiple files
curl -X POST http://localhost:8080/api/files/upload/multiple \
  -F "files=@file1.pdf" -F "files=@file2.jpg"
```

## JavaScript/Fetch

```javascript
const formData = new FormData();
formData.append('image', fileInput.files[0]);

const response = await fetch('http://localhost:8080/api/files/upload/image', {
  method: 'POST',
  body: formData
});

const result = await response.json();
// result.data.url contains the file URL
```

## React Component

```jsx
const [file, setFile] = useState(null);

const handleUpload = async () => {
  const formData = new FormData();
  formData.append('image', file);
  
  const res = await fetch('http://localhost:8080/api/files/upload/image', {
    method: 'POST',
    body: formData
  });
  
  const data = await res.json();
  console.log('Uploaded:', data.data.url);
};

return (
  <>
    <input type="file" onChange={(e) => setFile(e.target.files[0])} />
    <button onClick={handleUpload}>Upload</button>
  </>
);
```

## Response Format

```json
{
  "status": "success",
  "data": {
    "message": "File uploaded successfully",
    "path": "images/20251004173045_a8b3c5d1.jpg",
    "url": "/uploads/images/20251004173045_a8b3c5d1.jpg",
    "filename": "original.jpg",
    "size": "1.2 MB"
  }
}
```

## Allowed Image Extensions

`.jpg`, `.jpeg`, `.png`, `.gif`, `.webp`

## File Size Limits

- Images: 5MB
- Other files: 10MB

## Configuration Locations

- Service: `internal/service/file_service.go`
- Handler: `internal/handler/file_handler.go`
- Routes: `internal/routes/file_routes.go`
- Upload dir: `./uploads`

## Common Issues

| Issue | Solution |
|-------|----------|
| "No file uploaded" | Check form field name matches endpoint |
| "File too large" | Check file size limits |
| "Invalid extension" | Check allowed extensions list |
| "Permission denied" | Check uploads directory permissions |

## Quick Test

```bash
# Create test file
echo "test" > test.txt

# Upload
curl -X POST http://localhost:8080/api/files/upload -F "file=@test.txt"

# Should return success with file URL
```

## Files Created

- ✅ `internal/service/file_service.go`
- ✅ `internal/handler/file_handler.go`
- ✅ `internal/routes/file_routes.go`
- ✅ `docs/FILE_UPLOAD_API.md`
- ✅ `IMPLEMENTATION_SUMMARY.md`
- ✅ `test_upload.sh`
- ✅ `uploads/` directory structure
