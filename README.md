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
- **MongoDB** - Database

---

## 📁 Project Structure

```
go-auth-post/
├── cmd/ # Main application entrypoint
│ └── main.go
├── internal/
│ ├── config/ # App configuration (e.g., .env loading)
│ ├── database/ # Database connection
│ ├── handlers/ # Route handlers for auth and posts
│ ├── middleware/ # JWT middleware, error handling
│ ├── models/ # GORM models: User, Post, etc.
│ └── utils/ # Helpers: password hashing, token generation
├── migrations/ # SQL migration files
├── go.mod
├── go.sum
└── README.md
```

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/go-auth-post.git
cd go-auth-post
```

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

## API Endpoints

| Method | Endpoint        | Description               |
| ------ | --------------- | ------------------------- |
| POST   | `/api/register` | Register a new user       |
| POST   | `/api/login`    | Login and get a JWT       |
| GET    | `/api/posts`    | Get posts (auth required) |
| POST   | `/api/posts`    | Create a post (auth)      |

## ✍️ Author

Made with Go ❤️ by vince-II

Let me know if you'd like to include instructions for Docker, testing, or database migrations as well.
