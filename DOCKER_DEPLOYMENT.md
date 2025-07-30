# 🐳 Docker Deployment Guide

## Overview

This guide covers deploying your file storage server using Docker and Docker Compose in both development and production environments.

## 🚀 Quick Start

### Development Environment

```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up --build

# Access the application
# Web Interface: http://localhost:8080
# Swagger Docs: http://localhost:8080/swagger/index.html
```

### Production Environment

```bash
# Start production environment
docker-compose up --build -d

# Access the application
# Web Interface: http://localhost:8080
# Swagger Docs: http://localhost:8080/swagger/index.html
```

## 📁 File Structure

```
file_storage/
├── Dockerfile                 # Multi-stage Docker build
├── docker-compose.yml         # Production configuration
├── docker-compose.dev.yml     # Development configuration
├── nginx.conf                 # Nginx reverse proxy config
├── .dockerignore             # Docker build exclusions
└── DOCKER_DEPLOYMENT.md      # This guide
```

## 🏗️ Dockerfile Features

### Multi-Stage Build
- **Builder Stage**: Compiles Go application
- **Final Stage**: Minimal Alpine Linux image
- **Security**: Runs as non-root user
- **Optimization**: Small image size (~15MB)

### Security Features
- Non-root user execution
- Minimal attack surface
- Health checks
- Proper file permissions

## 🔧 Docker Compose Configurations

### Development (`docker-compose.dev.yml`)
- Hot reload with volume mounting
- Source code changes reflect immediately
- Debug-friendly configuration
- No SSL/HTTPS (for simplicity)

### Production (`docker-compose.yml`)
- Optimized for production
- Nginx reverse proxy with SSL
- Rate limiting and security headers
- Health checks and monitoring
- Optional Redis for caching

## 🌐 Production Features

### Nginx Reverse Proxy
- **SSL/TLS Support**: HTTPS encryption
- **Rate Limiting**: API protection
- **Caching**: Static file optimization
- **Security Headers**: XSS, CSRF protection
- **Compression**: Gzip optimization

### Security Features
- Rate limiting (10 req/s for API, 2 req/s for uploads)
- Security headers (HSTS, X-Frame-Options, etc.)
- SSL/TLS encryption
- Non-root container execution

## 📊 Monitoring & Health Checks

### Health Check Endpoints
- **Application**: `http://localhost:8080/`
- **Nginx**: `http://localhost/health`
- **Docker**: Built-in health checks

### Logging
- Application logs via Docker logs
- Nginx access/error logs
- Structured logging for production

## 🔐 SSL Configuration

### Self-Signed Certificates (Development)
```bash
# Generate self-signed certificates
mkdir ssl
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout ssl/key.pem -out ssl/cert.pem
```

### Let's Encrypt (Production)
```bash
# Use certbot for automatic SSL
docker run --rm -it \
  -v /etc/letsencrypt:/etc/letsencrypt \
  -v /var/lib/letsencrypt:/var/lib/letsencrypt \
  certbot/certbot certonly --standalone \
  -d yourdomain.com
```

## 🚀 Deployment Commands

### Development
```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up --build

# View logs
docker-compose -f docker-compose.dev.yml logs -f

# Stop development environment
docker-compose -f docker-compose.dev.yml down
```

### Production
```bash
# Start production environment
docker-compose up --build -d

# View logs
docker-compose logs -f

# Stop production environment
docker-compose down

# Start with production profile (includes Nginx)
docker-compose --profile production up -d
```

### Database & Storage
```bash
# View volumes
docker volume ls

# Backup uploads
docker run --rm -v file-storage-file-storage-data:/data \
  -v $(pwd):/backup alpine tar czf /backup/uploads-backup.tar.gz -C /data .

# Restore uploads
docker run --rm -v file-storage-file-storage-data:/data \
  -v $(pwd):/backup alpine tar xzf /backup/uploads-backup.tar.gz -C /data
```

## 🔍 Troubleshooting

### Common Issues

#### Port Already in Use
```bash
# Check what's using port 8080
netstat -tulpn | grep :8080

# Kill process or change port in docker-compose.yml
```

#### Permission Issues
```bash
# Fix volume permissions
docker-compose down
sudo chown -R $USER:$USER ./uploads
docker-compose up -d
```

#### SSL Certificate Issues
```bash
# Regenerate self-signed certificates
rm -rf ssl/
mkdir ssl
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout ssl/key.pem -out ssl/cert.pem
```

#### Container Won't Start
```bash
# Check container logs
docker-compose logs file-storage

# Rebuild without cache
docker-compose build --no-cache
```

### Debug Commands
```bash
# Enter running container
docker exec -it file-storage-server sh

# Check container resources
docker stats

# Inspect container
docker inspect file-storage-server
```

## 📈 Performance Optimization

### Resource Limits
```yaml
# Add to docker-compose.yml
services:
  file-storage:
    deploy:
      resources:
        limits:
          cpus: '1.0'
          memory: 512M
        reservations:
          cpus: '0.5'
          memory: 256M
```

### Scaling
```bash
# Scale to multiple instances
docker-compose up -d --scale file-storage=3
```

## 🔄 CI/CD Integration

### GitHub Actions Example
```yaml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Deploy to server
        run: |
          docker-compose pull
          docker-compose up -d
```

## 📋 Environment Variables

### Development
```bash
# .env.dev
TZ=UTC
LOG_LEVEL=debug
```

### Production
```bash
# .env.prod
TZ=UTC
LOG_LEVEL=info
NGINX_WORKER_PROCESSES=4
```

## 🎯 Best Practices

### Security
- ✅ Use non-root containers
- ✅ Implement rate limiting
- ✅ Enable SSL/TLS
- ✅ Regular security updates
- ✅ Monitor container logs

### Performance
- ✅ Use multi-stage builds
- ✅ Implement caching
- ✅ Optimize image size
- ✅ Monitor resource usage
- ✅ Use health checks

### Maintenance
- ✅ Regular backups
- ✅ Update dependencies
- ✅ Monitor disk space
- ✅ Rotate logs
- ✅ Test deployments

## 🚀 Next Steps

1. **Set up SSL certificates** for production
2. **Configure monitoring** (Prometheus, Grafana)
3. **Implement backup strategy** for uploads
4. **Add load balancing** for high availability
5. **Set up CI/CD pipeline** for automated deployments

---

**Happy Deploying! 🐳** 