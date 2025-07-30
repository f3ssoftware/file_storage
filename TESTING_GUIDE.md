# 🧪 Testing Guide - File Storage Server

## 🚀 Server Status
Your server is now running at: **http://localhost:8080**

## 📋 Quick Test Checklist

### ✅ 1. Basic Server Access
- [ ] **Main Interface**: http://localhost:8080
- [ ] **Swagger UI**: http://localhost:8080/swagger/index.html
- [ ] **API JSON**: http://localhost:8080/swagger/doc.json

### ✅ 2. Web Interface Test
1. Open http://localhost:8080 in your browser
2. You should see a beautiful upload interface
3. Try uploading a test file (JPG, PNG, etc.)
4. Verify you get a success message with file details

### ✅ 3. Swagger Documentation Test
1. Open http://localhost:8080/swagger/index.html
2. You should see the interactive API documentation
3. Test the `/upload` endpoint:
   - Click "Try it out"
   - Upload a file
   - Click "Execute"
   - Verify 201 response with JSON details

4. Test the `/files/{filename}` endpoint:
   - Use the filename from upload response
   - Click "Try it out"
   - Enter filename
   - Click "Execute"
   - Verify 200 response with file content

### ✅ 4. API Testing with curl

#### Test File Upload:
```bash
curl -X POST -F "file=@test.jpg" http://localhost:8080/upload
```
Expected response:
```json
{
  "filename": "test.jpg",
  "url": "/files/test.jpg",
  "size": 12345,
  "message": "File uploaded successfully"
}
```

#### Test File Download:
```bash
curl http://localhost:8080/files/test.jpg
```
Expected: File content returned

### ✅ 5. Error Handling Tests

#### Test Invalid File Type:
```bash
curl -X POST -F "file=@test.exe" http://localhost:8080/upload
```
Expected: 400 error with "file type not allowed"

#### Test Missing File:
```bash
curl -X POST http://localhost:8080/upload
```
Expected: 400 error with "No file provided"

#### Test Non-existent File Download:
```bash
curl http://localhost:8080/files/nonexistent.jpg
```
Expected: 404 error

## 🎯 Test Scenarios

### Scenario 1: Complete Upload & Download Flow
1. Upload a file via web interface
2. Copy the filename from the response
3. Test download via Swagger UI
4. Verify file content matches original

### Scenario 2: API Integration Test
1. Upload via curl
2. Download via curl
3. Compare file integrity

### Scenario 3: Error Handling Test
1. Try uploading unsupported file types
2. Try uploading files larger than 10MB
3. Try downloading non-existent files
4. Verify appropriate error messages

## 📊 Expected Results

### ✅ Success Indicators:
- [ ] Web interface loads without errors
- [ ] File uploads work (web + API)
- [ ] File downloads work (web + API)
- [ ] Swagger UI loads and functions
- [ ] Error handling works correctly
- [ ] File validation works (size + type)
- [ ] CORS headers are set correctly

### ❌ Common Issues & Solutions:

#### Issue: "Swagger UI not loading"
**Solution**: 
- Check server is running: `go run cmd/server/main.go`
- Verify port 8080 is not in use
- Check browser console for errors

#### Issue: "File upload fails"
**Solution**:
- Check file size (max 10MB)
- Check file type (JPG, PNG, GIF, PDF, TXT, DOC, DOCX, XLS, XLSX)
- Check uploads directory exists

#### Issue: "Import errors"
**Solution**:
```bash
go mod tidy
go run cmd/server/main.go
```

#### Issue: "Swagger docs not updating"
**Solution**:
```bash
swag init -g cmd/server/main.go -o docs
go run cmd/server/main.go
```

## 🔧 Manual Testing Steps

### Step 1: Web Interface
1. Open browser to http://localhost:8080
2. Upload a test image (JPG/PNG)
3. Verify success message appears
4. Click the file URL to download
5. Verify file downloads correctly

### Step 2: Swagger UI
1. Open http://localhost:8080/swagger/index.html
2. Expand the `/upload` endpoint
3. Click "Try it out"
4. Upload a test file
5. Click "Execute"
6. Verify JSON response
7. Test the `/files/{filename}` endpoint
8. Enter the filename and execute
9. Verify file content is returned

### Step 3: API Testing
1. Use curl to upload a file
2. Use curl to download the file
3. Compare original and downloaded files
4. Test error scenarios

## 📝 Test Results Log

| Test | Status | Notes |
|------|--------|-------|
| Web Interface Load | ⬜ | |
| File Upload (Web) | ⬜ | |
| File Download (Web) | ⬜ | |
| Swagger UI Load | ⬜ | |
| API Upload | ⬜ | |
| API Download | ⬜ | |
| Error Handling | ⬜ | |
| File Validation | ⬜ | |

## 🎉 Success Criteria

Your file storage server is working correctly when:
- ✅ All endpoints respond correctly
- ✅ File uploads and downloads work
- ✅ Swagger documentation is accessible
- ✅ Error handling provides clear messages
- ✅ File validation prevents invalid uploads
- ✅ Web interface is user-friendly

## 🚀 Next Steps

After successful testing:
1. **Share with others**: Send them the Swagger UI URL
2. **Add more features**: Authentication, file deletion, etc.
3. **Deploy**: Consider deploying to a cloud service
4. **Monitor**: Add logging and monitoring
5. **Scale**: Consider database integration for metadata

---

**Happy Testing! 🎯**

Your file storage server is now ready for production use with comprehensive documentation and testing capabilities. 