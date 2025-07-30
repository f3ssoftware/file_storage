package main

import (
	"log"
	"net/http"

	"github.com/f3ssoftware/file_storage/docs"
	"github.com/f3ssoftware/file_storage/internal/handler"
	"github.com/f3ssoftware/file_storage/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title File Storage Server API
// @version 1.0
// @description A simple file storage server built with Go that allows you to upload and serve files.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Initialize Swagger docs
	docs.SwaggerInfo.Title = "File Storage Server API"
	docs.SwaggerInfo.Description = "A simple file storage server built with Go that allows you to upload and serve files."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	store := storage.NewLocalStorage("./uploads")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Swagger documentation endpoint
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// File upload endpoint
	r.Post("/upload", handler.UploadHandler(store))

	// File serving endpoint
	r.Get("/files/{filename}", handler.ServeHandler(store))

	// Simple HTML interface for testing
	r.Get("/", serveHTML)

	log.Println("🚀 File Storage Server running at http://localhost:8080")
	log.Println("📁 Upload files at: http://localhost:8080")
	log.Println("📂 Access files at: http://localhost:8080/files/{filename}")
	log.Println("📚 API Documentation at: http://localhost:8080/swagger/index.html")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>File Storage Server</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
            text-align: center;
            margin-bottom: 30px;
        }
        .upload-section {
            border: 2px dashed #ddd;
            padding: 40px;
            text-align: center;
            border-radius: 10px;
            margin-bottom: 20px;
            transition: border-color 0.3s;
        }
        .upload-section:hover {
            border-color: #007bff;
        }
        .file-input {
            margin: 20px 0;
        }
        .upload-btn {
            background: #007bff;
            color: white;
            border: none;
            padding: 12px 30px;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            transition: background 0.3s;
        }
        .upload-btn:hover {
            background: #0056b3;
        }
        .upload-btn:disabled {
            background: #ccc;
            cursor: not-allowed;
        }
        .result {
            margin-top: 20px;
            padding: 15px;
            border-radius: 5px;
            display: none;
        }
        .success {
            background: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
        }
        .error {
            background: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
        }
        .info {
            background: #e2e3e5;
            color: #383d41;
            border: 1px solid #d6d8db;
            margin-bottom: 20px;
        }
        .progress {
            width: 100%;
            height: 20px;
            background: #f0f0f0;
            border-radius: 10px;
            overflow: hidden;
            margin: 10px 0;
            display: none;
        }
        .progress-bar {
            height: 100%;
            background: #007bff;
            width: 0%;
            transition: width 0.3s;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>📁 File Storage Server</h1>
        
        <div class="info">
            <strong>Supported file types:</strong> JPG, PNG, GIF, PDF, TXT, DOC, DOCX, XLS, XLSX<br>
            <strong>Maximum file size:</strong> 10MB
        </div>

        <div class="upload-section">
            <h3>📤 Upload File</h3>
            <form id="uploadForm" enctype="multipart/form-data">
                <div class="file-input">
                    <input type="file" id="fileInput" name="file" accept=".jpg,.jpeg,.png,.gif,.pdf,.txt,.doc,.docx,.xls,.xlsx" required>
                </div>
                <button type="submit" class="upload-btn" id="uploadBtn">Upload File</button>
            </form>
            
            <div class="progress" id="progress">
                <div class="progress-bar" id="progressBar"></div>
            </div>
            
            <div id="result" class="result"></div>
        </div>

        <div class="info">
            <strong>API Endpoints:</strong><br>
            • POST /upload - Upload a file<br>
            • GET /files/{filename} - Download a file
        </div>
    </div>

    <script>
        document.getElementById('uploadForm').addEventListener('submit', async function(e) {
            e.preventDefault();
            
            const fileInput = document.getElementById('fileInput');
            const uploadBtn = document.getElementById('uploadBtn');
            const result = document.getElementById('result');
            const progress = document.getElementById('progress');
            const progressBar = document.getElementById('progressBar');
            
            if (!fileInput.files[0]) {
                showResult('Please select a file', 'error');
                return;
            }
            
            const formData = new FormData();
            formData.append('file', fileInput.files[0]);
            
            uploadBtn.disabled = true;
            uploadBtn.textContent = 'Uploading...';
            progress.style.display = 'block';
            result.style.display = 'none';
            
            try {
                const response = await fetch('/upload', {
                    method: 'POST',
                    body: formData
                });
                
                const data = await response.json();
                
                if (response.ok) {
                    const fileUrl = window.location.origin + data.url;
                    showResult(
                        '<strong>✅ Upload successful!</strong><br>' +
                        '<strong>Filename:</strong> ' + data.filename + '<br>' +
                        '<strong>Size:</strong> ' + (data.size / 1024).toFixed(2) + ' KB<br>' +
                        '<strong>URL:</strong> <a href="' + fileUrl + '" target="_blank">' + fileUrl + '</a>',
                        'success'
                    );
                } else {
                    showResult('❌ Error: ' + data.message, 'error');
                }
            } catch (error) {
                showResult('❌ Network error: ' + error.message, 'error');
            } finally {
                uploadBtn.disabled = false;
                uploadBtn.textContent = 'Upload File';
                progress.style.display = 'none';
                progressBar.style.width = '0%';
            }
        });
        
        function showResult(message, type) {
            const result = document.getElementById('result');
            result.innerHTML = message;
            result.className = 'result ' + type;
            result.style.display = 'block';
        }
        
        // Simulate progress bar
        document.getElementById('fileInput').addEventListener('change', function() {
            const progressBar = document.getElementById('progressBar');
            let width = 0;
            const interval = setInterval(() => {
                if (width >= 90) {
                    clearInterval(interval);
                } else {
                    width++;
                    progressBar.style.width = width + '%';
                }
            }, 50);
        });
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
