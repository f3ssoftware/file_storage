package handler

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/f3ssoftware/file_storage/internal/storage"
	"github.com/go-chi/chi"
)

// UploadResponse represents the response structure for file upload operations
// @Description Response structure for file upload operations
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

// ErrorResponse represents the error response structure
// @Description Error response structure
type ErrorResponse struct {
	// @Description Error message
	Message string `json:"message" example:"File too large (max 10MB)"`
}

// UploadHandler handles file upload requests
// @Summary Upload a file
// @Description Upload a file to the storage server
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload" default("")
// @Success 201 {object} UploadResponse "File uploaded successfully"
// @Failure 400 {object} ErrorResponse "Bad request - invalid file or no file provided"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /upload [post]
func UploadHandler(store *storage.LocalStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for web interface
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Parse multipart form with 32MB max memory
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			// Try "image" field for backward compatibility
			file, header, err = r.FormFile("image")
			if err != nil {
				sendErrorResponse(w, "No file provided", http.StatusBadRequest)
				return
			}
		}
		defer file.Close()

		// Validate file
		if err := validateFile(header); err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Generate safe filename
		filename := generateSafeFilename(header.Filename)

		// Save file
		if err := store.Save(filename, file); err != nil {
			sendErrorResponse(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		// Send success response
		response := UploadResponse{
			Filename: filename,
			URL:      fmt.Sprintf("/files/%s", filename),
			Size:     header.Size,
			Message:  "File uploaded successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

// ServeHandler handles file serving requests
// @Summary Download a file
// @Description Download a file from the storage server
// @Tags files
// @Accept json
// @Produce octet-stream
// @Param filename path string true "Name of the file to download" example("example.jpg")
// @Success 200 {file} file "File content"
// @Failure 400 {object} ErrorResponse "Bad request - filename required"
// @Failure 404 {object} ErrorResponse "File not found"
// @Router /files/{filename} [get]
func ServeHandler(store *storage.LocalStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := chi.URLParam(r, "filename")
		if filename == "" {
			http.Error(w, "Filename required", http.StatusBadRequest)
			return
		}

		filePath, err := store.Load(filename)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// Set appropriate headers for file serving
		ext := strings.ToLower(filepath.Ext(filename))
		switch ext {
		case ".jpg", ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".gif":
			w.Header().Set("Content-Type", "image/gif")
		case ".pdf":
			w.Header().Set("Content-Type", "application/pdf")
		case ".txt":
			w.Header().Set("Content-Type", "text/plain")
		default:
			w.Header().Set("Content-Type", "application/octet-stream")
		}

		http.ServeFile(w, r, filePath)
	}
}

// validateFile validates the uploaded file
// @Description Validates file size and type
func validateFile(header *multipart.FileHeader) error {
	// Check file size (10MB limit)
	if header.Size > 10<<20 {
		return fmt.Errorf("file too large (max 10MB)")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif", ".pdf", ".txt", ".doc", ".docx", ".xls", ".xlsx"}

	allowed := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			allowed = true
			break
		}
	}

	if !allowed {
		return fmt.Errorf("file type not allowed")
	}

	return nil
}

// generateSafeFilename generates a safe filename
// @Description Generates a safe filename by removing path separators
func generateSafeFilename(originalName string) string {
	// Remove any path separators and replace with underscore
	safeName := strings.ReplaceAll(originalName, "/", "_")
	safeName = strings.ReplaceAll(safeName, "\\", "_")

	// Add timestamp to prevent conflicts
	// For now, just return the safe name
	// In a real app, you might want to add timestamp or UUID
	return safeName
}

// sendErrorResponse sends an error response
// @Description Sends a JSON error response
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := ErrorResponse{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
