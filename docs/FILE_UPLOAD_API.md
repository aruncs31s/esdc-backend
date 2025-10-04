# File Upload API Documentation

## Overview

This API supports file uploads with validation, multiple file uploads, and static file serving.

## Endpoints

### 1. Upload Single Image

Upload a single image file (images only).

**Endpoint:** `POST /api/files/upload/image`

**Content-Type:** `multipart/form-data`

**Parameters:**
- `image` (file, required): The image file to upload

**Allowed Extensions:** `.jpg`, `.jpeg`, `.png`, `.gif`, `.webp`

**Max Size:** 5MB

**Example Request (curl):**
```bash
curl -X POST http://localhost:8080/api/files/upload/image \
  -F "image=@/path/to/image.jpg"
```

**Success Response (200):**
```json
{
  "status": "success",
  "data": {
    "message": "Image uploaded successfully",
    "path": "images/20251004173045_a8b3c5d1.jpg",
    "url": "/uploads/images/20251004173045_a8b3c5d1.jpg"
  }
}
```

**Error Response (400):**
```json
{
  "status": "error",
  "message": "Invalid file",
  "error": "file extension .pdf is not allowed"
}
```

---

### 2. Upload Single File

Upload a single file of any type.

**Endpoint:** `POST /api/files/upload`

**Content-Type:** `multipart/form-data`

**Parameters:**
- `file` (file, required): The file to upload
- `dir` (query, optional): Directory to upload to (default: "files")

**Max Size:** 10MB

**Example Request (curl):**
```bash
# Upload to default "files" directory
curl -X POST http://localhost:8080/api/files/upload \
  -F "file=@/path/to/document.pdf"

# Upload to custom directory
curl -X POST "http://localhost:8080/api/files/upload?dir=documents" \
  -F "file=@/path/to/document.pdf"
```

**Success Response (200):**
```json
{
  "status": "success",
  "data": {
    "message": "File uploaded successfully",
    "path": "files/20251004173045_a8b3c5d1.pdf",
    "url": "/uploads/files/20251004173045_a8b3c5d1.pdf",
    "filename": "document.pdf",
    "size": "1.2 MB"
  }
}
```

---

### 3. Upload Multiple Files

Upload multiple files at once.

**Endpoint:** `POST /api/files/upload/multiple`

**Content-Type:** `multipart/form-data`

**Parameters:**
- `files` (files, required): Multiple files to upload
- `dir` (query, optional): Directory to upload to (default: "files")

**Max Size:** 10MB per file

**Example Request (curl):**
```bash
curl -X POST http://localhost:8080/api/files/upload/multiple \
  -F "files=@/path/to/file1.pdf" \
  -F "files=@/path/to/file2.jpg" \
  -F "files=@/path/to/file3.doc"
```

**Success Response (200):**
```json
{
  "status": "success",
  "data": {
    "message": "Uploaded 2 of 3 files",
    "uploaded": [
      {
        "filename": "file1.pdf",
        "path": "files/20251004173045_a8b3c5d1.pdf",
        "url": "/uploads/files/20251004173045_a8b3c5d1.pdf",
        "size": "1.2 MB"
      },
      {
        "filename": "file2.jpg",
        "path": "files/20251004173046_b9c4d2e3.jpg",
        "url": "/uploads/files/20251004173046_b9c4d2e3.jpg",
        "size": "850.5 KB"
      }
    ],
    "failed": [
      {
        "filename": "file3.doc",
        "error": "file size exceeds maximum allowed size of 10485760 bytes"
      }
    ],
    "uploaded_count": 2,
    "failed_count": 1
  }
}
```

---

### 4. Access Uploaded Files

Access uploaded files via static file serving.

**Endpoint:** `GET /uploads/{path}`

**Example:**
```
http://localhost:8080/uploads/images/20251004173045_a8b3c5d1.jpg
http://localhost:8080/uploads/files/20251004173045_a8b3c5d1.pdf
```

---

## Frontend Integration Examples

### JavaScript (Fetch API)

```javascript
// Single image upload
async function uploadImage(file) {
  const formData = new FormData();
  formData.append('image', file);

  const response = await fetch('http://localhost:8080/api/files/upload/image', {
    method: 'POST',
    body: formData
  });

  return await response.json();
}

// Multiple files upload
async function uploadMultipleFiles(files) {
  const formData = new FormData();
  files.forEach(file => {
    formData.append('files', file);
  });

  const response = await fetch('http://localhost:8080/api/files/upload/multiple', {
    method: 'POST',
    body: formData
  });

  return await response.json();
}
```

### React Example

```jsx
import { useState } from 'react';

function FileUpload() {
  const [file, setFile] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [result, setResult] = useState(null);

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleUpload = async () => {
    if (!file) return;

    setUploading(true);
    const formData = new FormData();
    formData.append('image', file);

    try {
      const response = await fetch('http://localhost:8080/api/files/upload/image', {
        method: 'POST',
        body: formData
      });
      const data = await response.json();
      setResult(data);
    } catch (error) {
      console.error('Upload failed:', error);
    } finally {
      setUploading(false);
    }
  };

  return (
    <div>
      <input type="file" onChange={handleFileChange} accept="image/*" />
      <button onClick={handleUpload} disabled={uploading}>
        {uploading ? 'Uploading...' : 'Upload'}
      </button>
      {result && (
        <div>
          <p>{result.data.message}</p>
          <img src={`http://localhost:8080${result.data.url}`} alt="Uploaded" />
        </div>
      )}
    </div>
  );
}
```

### HTML Form Example

```html
<!DOCTYPE html>
<html>
<head>
    <title>File Upload</title>
</head>
<body>
    <h2>Upload Image</h2>
    <form id="uploadForm">
        <input type="file" id="fileInput" name="image" accept="image/*" required>
        <button type="submit">Upload</button>
    </form>
    <div id="result"></div>

    <script>
        document.getElementById('uploadForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            
            const formData = new FormData();
            const fileInput = document.getElementById('fileInput');
            formData.append('image', fileInput.files[0]);

            try {
                const response = await fetch('http://localhost:8080/api/files/upload/image', {
                    method: 'POST',
                    body: formData
                });
                const data = await response.json();
                
                if (data.status === 'success') {
                    document.getElementById('result').innerHTML = `
                        <p>Upload successful!</p>
                        <img src="http://localhost:8080${data.data.url}" width="300">
                    `;
                }
            } catch (error) {
                console.error('Error:', error);
            }
        });
    </script>
</body>
</html>
```

---

## Error Codes

| Status Code | Description |
|------------|-------------|
| 200 | Success |
| 400 | Bad Request (invalid file, missing parameters, etc.) |
| 500 | Internal Server Error |

---

## Security Considerations

1. **File Type Validation**: Only allowed file extensions are accepted
2. **File Size Limits**: Enforced on both image and file uploads
3. **Unique Filenames**: Files are renamed to prevent conflicts
4. **CORS**: Configured for specific origins only

---

## Configuration

You can modify the following in the code:

- **Upload Directory**: Set in `routes.go` (`./uploads`)
- **Max Image Size**: Set in `file_handler.go` (default: 5MB)
- **Max File Size**: Set in `file_handler.go` (default: 10MB)
- **Allowed Image Extensions**: Set in `file_handler.go`

---

## Testing with Postman

1. Create a new POST request
2. Set URL to: `http://localhost:8080/api/files/upload/image`
3. Select Body → form-data
4. Add key: `image` (change type to File)
5. Choose a file
6. Send request

---

## Notes

- Uploaded files are stored in the `uploads/` directory
- The `uploads/` directory is served as static files at `/uploads/` path
- Files are automatically renamed with timestamp and random ID to prevent conflicts
- The API returns both the relative path and full URL for accessing the uploaded file
