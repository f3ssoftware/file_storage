# File Storage Server

A simple file storage server built with Go that allows you to upload and serve files.

## Features

- 📤 **File Upload**: Upload files via HTTP POST request
- 📂 **File Serving**: Download files via HTTP GET request
- 🌐 **Web Interface**: Beautiful HTML interface for easy testing
- 📚 **API Documentation**: Interactive Swagger/OpenAPI documentation
- 🔒 **File Validation**: Size and type restrictions for security
- 📊 **Progress Tracking**: Visual upload progress
- 🎨 **Modern UI**: Clean and responsive design

## Supported File Types

- Images: JPG, JPEG, PNG, GIF
- Documents: PDF, TXT, DOC, DOCX, XLS, XLSX
- Maximum file size: 10MB

## Quick Start

### 1. Run the Server

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

### 2. Use the Web Interface

Open your browser and go to `http://localhost:8080` to use the web interface for uploading files.

### 3. API Documentation

Access the interactive API documentation at `http://localhost:8080/swagger/index.html` to test endpoints and view detailed API specifications.

### 4. API Endpoints

#### Upload a File
```bash
curl -X POST -F "file=@your-file.jpg" http://localhost:8080/upload
```

Response:
```json
{
  "filename": "your-file.jpg",
  "url": "/files/your-file.jpg",
  "size": 12345,
  "message": "File uploaded successfully"
}
```

#### Download a File
```bash
curl http://localhost:8080/files/your-file.jpg
```

## Project Structure

```
file_storage/
├── cmd/
│   └── server/
│       └── main.go          # Server entry point
├── internal/
│   ├── handler/
│   │   └── upload.go        # HTTP handlers
│   └── storage/
│       └── local.go         # File storage implementation
├── docs/                    # Generated Swagger documentation
│   ├── docs.go              # Swagger docs Go file
│   ├── swagger.json         # OpenAPI JSON specification
│   └── swagger.yaml         # OpenAPI YAML specification
├── uploads/                 # Uploaded files directory (created automatically)
├── go.mod                   # Go module file
├── API_DOCUMENTATION.md     # Detailed API documentation guide
└── README.md               # This file
```

## API Reference

### POST /upload

Upload a file to the server.

**Request:**
- Method: `POST`
- Content-Type: `multipart/form-data`
- Body: Form with field `file` containing the file

**Response:**
```json
{
  "filename": "uploaded-file.jpg",
  "url": "/files/uploaded-file.jpg",
  "size": 12345,
  "message": "File uploaded successfully"
}
```

### GET /files/{filename}

Download a file from the server.

**Request:**
- Method: `GET`
- URL: `/files/{filename}` where `{filename}` is the name of the file

**Response:**
- File content with appropriate Content-Type header

## Error Handling

The server returns appropriate HTTP status codes:

- `200 OK`: File served successfully
- `201 Created`: File uploaded successfully
- `400 Bad Request`: Invalid request (no file, file too large, unsupported type)
- `404 Not Found`: File not found
- `500 Internal Server Error`: Server error

## Development

### Prerequisites

- Go 1.24 or later

### Dependencies

The project uses the following main dependencies:
- `github.com/go-chi/chi`: HTTP router and middleware

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o file-storage-server cmd/server/main.go
```

## Security Considerations

- File size is limited to 10MB
- Only specific file types are allowed
- Filenames are sanitized to prevent path traversal
- Files are stored in a dedicated uploads directory

## Future Enhancements

- [ ] Add authentication
- [ ] Support for file deletion
- [ ] File metadata storage
- [ ] Image thumbnail generation
- [ ] Cloud storage backend (AWS S3, Google Cloud Storage)
- [ ] Database integration for file metadata
- [ ] Rate limiting
- [ ] File compression

## License

This project is open source and available under the MIT License. 