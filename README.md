# auth-post-api

A simple Go backend API with user authentication and post management.  
Built with [Fiber](https://gofiber.io/), [GORM](https://gorm.io/), and JWT-based authentication.

---

## ✨ Features

- User registration and login
- JWT-based authentication
- Create, read, update, and delete posts
- Each post belongs to a registered user

---

## 🛠 Tech Stack

- **Go** (Golang)
- **Fiber** – Web framework
- **JWT** – Secure token-based authentication
- **SQLC + PGX** - Database

---

## 🚀 Getting Started

### 1. Build and start the containers:

```sh
docker compose up --build
```

1.1. The application should now be running and accessible at `http://localhost:3000`.

### 2. Set up environment variables

create a `.env` file in the root

```
PORT=3000
DATABASE_URL=your_database_url_here
JWT_SECRET=your_secret_key
```

### 3. Run the app

```bash
go run cmd/main.go
```
