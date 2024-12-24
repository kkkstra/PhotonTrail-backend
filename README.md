# **PhotonTrail-backend**

[简体中文](./README_zhCN.md)

## **Project Overview**

**PhotonTrail-backend** is a community focused on sharing and interacting with photography works.

This backend system is built with **Go** and the **Gin** web framework, using **MySQL** as the database, and **JWT** for authentication. The system integrates with OSS STS Token service.

## **Key Features**

- **User Authentication**: JWT-based login and registration.
- **Post Management**: Create, update, delete, and search for posts.
- **OSS Integration**: STS Token service for OSS.
- **RESTful API**: For smooth front-end and back-end interaction.

## **Project Structure**

```
├── Dockerfile            # Dockerfile
├── LICENSE               # License file
├── README.md             # Project overview and instructions
├── cmd                   # App entry points
├── configs               # Config files
├── data                  # Data storage
├── docker-compose.yml    # Docker Compose config for multi-service setup
├── internal              # Internal code, project-only
├── pkg                   # Reusable packages
└── scripts               # Scripts
```

## **Tech Stack**

- **Language**: Go
- **Web Framework**: Gin
- **Database**: MySQL
- **Authentication**: JWT

## **Requirements**

- **Go** 1.16+
- **MySQL** 5.7+
- **Python** 3.8+ (for OSS STS Token service)
- **Docker** (optional, for deployment)

## **Installation Guide**

### **1. Clone the repository**

```bash
git clone https://github.com/your-repo/PhotonTrail-backend.git
cd PhotonTrail-backend
```

### **2. Install Go dependencies**

Ensure you have Go installed, then run:

```bash
go mod tidy
```

### **3. Set up the MySQL Database**

Ensure MySQL is running and create the database:

```sql
CREATE DATABASE photon_trail;
```

### **4. Run Database Migrations**

Run any required migrations to set up the database schema.

### **5. Start the Application**

Launch the backend locally:

```bash
go run main.go
```

The server will run at `http://localhost:8001`.

### **7. Start the OSS STS TOKEN Service**

Navigate to the `oss-uploader` directory and start the service:

```bash
cd oss-uploader
python main.py
```

The service will run on `http://localhost:8000`.

## **Deployment**

### **Docker Deployment**

To deploy the backend using Docker:

1. **Build the Docker image:**

   ```bash
   docker build -t PhotonTrail-backend .
   ```

2. **Run the Docker container:**

   ```bash
   docker run -p 8001:8000 PhotonTrail-backend
   ```

### **Docker Compose Deployment**

```bash
docker compose up -d --build
```

## **Contributing**

We welcome contributions! If you'd like to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-xyz`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push your changes to your fork (`git push origin feature-xyz`).
5. Open a pull request to the main repository.

## **License**

This project is licensed under the **MIT License**. For more details, see the [LICENSE](LICENSE) file.