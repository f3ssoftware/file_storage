# API Documentation Guide

## Overview

Your file storage server now includes comprehensive Swagger/OpenAPI documentation that can be accessed through a web interface. This documentation is automatically generated from the code comments and provides an interactive way to test and understand your API.

## Accessing the Documentation

### Web Interface
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **API JSON**: http://localhost:8080/swagger/doc.json
- **API YAML**: http://localhost:8080/swagger/doc.yaml

## Swagger Annotations Used

### Main API Information
```go
// @title File Storage Server API
// @version 1.0
// @description A simple file storage server built with Go that allows you to upload and serve files.
// @host localhost:8080
// @BasePath /
```

### Endpoint Documentation
```go
// @Summary Upload a file
// @Description Upload a file to the storage server
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 201 {object} UploadResponse "File uploaded successfully"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Router /upload [post]
```

### Response Structure Documentation
```go
type UploadResponse struct {
    // @Description The name of the uploaded file
    Filename string `json:"filename" example:"example.jpg"`
    // @Description The URL path to access the uploaded file
    URL string `json:"url" example:"/files/example.jpg"`
    // @Description The size of the uploaded file in bytes
    Size int64 `json:"size" example:"12345"`
    // @Description A message describing the operation result
    Message string `json:"message" example:"File uploaded successfully"`
}
```

## Available Endpoints

### 1. Upload File
- **URL**: `POST /upload`
- **Content-Type**: `multipart/form-data`
- **Parameters**: 
  - `file` (formData, required): The file to upload
- **Response**: JSON with file details
- **Status Codes**:
  - `201`: File uploaded successfully
  - `400`: Bad request (invalid file, no file, file too large)
  - `500`: Internal server error

### 2. Download File
- **URL**: `GET /files/{filename}`
- **Parameters**:
  - `filename` (path, required): Name of the file to download
- **Response**: File content with appropriate Content-Type
- **Status Codes**:
  - `200`: File served successfully
  - `400`: Bad request (filename required)
  - `404`: File not found

## File Validation Rules

### Supported File Types
- Images: `.jpg`, `.jpeg`, `.png`, `.gif`
- Documents: `.pdf`, `.txt`, `.doc`, `.docx`, `.xls`, `.xlsx`

### File Size Limits
- Maximum file size: 10MB

## Testing with Swagger UI

1. **Start the server**:
   ```bash
   go run cmd/server/main.go
   ```

2. **Open Swagger UI**: Navigate to http://localhost:8080/swagger/index.html

3. **Test endpoints**:
   - Click on any endpoint to expand it
   - Click "Try it out" to test the endpoint
   - Fill in the required parameters
   - Click "Execute" to make the request
   - View the response and status code

## Regenerating Documentation

When you make changes to your API endpoints or add new ones, regenerate the Swagger documentation:

```bash
swag init -g cmd/server/main.go -o docs
```

## Swagger Annotations Reference

### API Information
- `@title`: API title
- `@version`: API version
- `@description`: API description
- `@host`: API host
- `@BasePath`: API base path
- `@contact.name`: Contact name
- `@contact.url`: Contact URL
- `@contact.email`: Contact email
- `@license.name`: License name
- `@license.url`: License URL

### Endpoint Documentation
- `@Summary`: Brief endpoint description
- `@Description`: Detailed endpoint description
- `@Tags`: Group endpoints by tags
- `@Accept`: Accepted content types
- `@Produce`: Produced content types
- `@Param`: Parameter documentation
- `@Success`: Success response documentation
- `@Failure`: Error response documentation
- `@Router`: Route path and method

### Parameter Types
- `path`: URL path parameter
- `query`: Query string parameter
- `formData`: Form data parameter
- `body`: Request body parameter
- `header`: HTTP header parameter

### Response Types
- `{object} TypeName`: JSON object response
- `{file}`: File download response
- `{string}`: String response
- `{array} TypeName`: Array response

## Example Usage

### Using curl
```bash
# Upload a file
curl -X POST -F "file=@example.jpg" http://localhost:8080/upload

# Download a file
curl http://localhost:8080/files/example.jpg
```

### Using Swagger UI
1. Open http://localhost:8080/swagger/index.html
2. Find the `/upload` endpoint
3. Click "Try it out"
4. Upload a file using the file picker
5. Click "Execute"
6. View the response with file details

## Benefits of Swagger Documentation

1. **Interactive Testing**: Test APIs directly from the browser
2. **Auto-generated**: Documentation stays in sync with code
3. **Client Generation**: Can generate client libraries
4. **Team Collaboration**: Easy to share API specifications
5. **API Discovery**: New team members can quickly understand the API

## Troubleshooting

### Common Issues

1. **Swagger UI not loading**: Make sure the server is running and accessible
2. **Documentation not updating**: Regenerate docs with `swag init`
3. **Import errors**: Run `go mod tidy` to fix dependencies
4. **Missing endpoints**: Check that all endpoints have proper Swagger annotations

### Regenerating Documentation
```bash
# Remove old docs
rm -rf docs/

# Generate new docs
swag init -g cmd/server/main.go -o docs

# Restart server
go run cmd/server/main.go
```

## Next Steps

1. **Add Authentication**: Include `@security` annotations for protected endpoints
2. **Add More Endpoints**: Document additional API endpoints as you add them
3. **Customize UI**: Modify the Swagger UI theme and branding
4. **Export Documentation**: Export to PDF or other formats for sharing
5. **Version Control**: Track documentation changes in your repository 