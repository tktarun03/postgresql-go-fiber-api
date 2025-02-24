# postgresql-go-fiber-api

A simple REST API built with Go Fiber and PostgreSQL for efficient data management.

## 🚀 Features
- Uses **Go Fiber** for high-performance web API development.
- Connects to **PostgreSQL** for persistent data storage.
- Implements **GET & POST** endpoints for task management.

## 📂 Project Structure
```
postgresql-go-fiber-api/
│── src/
│   ├── main.go  # Fiber API server
│── go.mod  # Go module file
│── README.md
```

## 🛠 Installation & Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/tktarun03/postgresql-go-fiber-api.git
   cd postgresql-go-fiber-api/src
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Start PostgreSQL database (Docker recommended):
   ```bash
   docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=taskdb -p 5432:5432 -d postgres
   ```

4. Start the Go Fiber API server:
   ```bash
   go run main.go
   ```

5. Test the API endpoints:
   - **GET** `/tasks` → Fetch all tasks  
   - **POST** `/tasks` → Add a new task  

## 👨‍💻 About the Author

🚀 Created by [Arunkumar Thamilarasu](https://github.com/tktarun03) | UI Technical Architect | Go & Database Expert

## ⭐ Contribute & Support
- Fork & Star this repository! ⭐
- Submit Issues and PRs to improve the API.

---
🎯 **Follow me on GitHub**: [@tktarun03](https://github.com/tktarun03)
