# go-store

**Basic E-Commerce Website built with Golang using Server-Side Template Rendering and Fake Store API**

`go-store` is a simple e-commerce web application developed in Go.  
It demonstrates how to build a server-side rendered web app using Go’s standard libraries and integrates with the **Fake Store API** to fetch product data.

---

## ✨ Features

- Server-side web application written in Go
- Product data fetched from Fake Store API
- HTML template rendering using `html/template`
- Basic authentication structure
- Static asset support (HTML / CSS)
- Clean and simple project structure

---

## 📁 Project Structure

```
├── auth/                # Authentication related logic
├── constants/           # Application constants (API URLs, etc.)
├── handler/             # HTTP handlers
├── template/            # HTML template files
├── .gitattributes
├── go.mod
├── go.sum
└── main.go              # Application entry point
```

---

## 🛠 Requirements

- Go (version 1.18 or higher recommended)
- Internet connection (for Fake Store API)

---

## 🚀 Installation

Clone the repository:

```bash
git clone https://github.com/arasdenizhan/go-store.git
cd go-store
```

Download dependencies:

```bash
go mod tidy
```

---

## ▶️ Running the Application

Start the server with:

```bash
go run main.go
```

If everything works correctly, the server will start at:

```
http://localhost:8080
```

Open this address in your browser to view the application.

---

## 🔐 Demo User Credentials

You can use the following example user to log in to the application:

```
Username: donore
Password: ewedon
```

---

## 🧠 How It Works

- **main.go** initializes the HTTP server and routes.
- **handler/** contains functions that handle incoming HTTP requests.
- **auth/** includes basic authentication-related logic.
- **template/** contains HTML templates rendered on the server side.

The application fetches product data from Fake Store API and injects it into HTML templates before sending responses to the client.

---

## 🌐 Fake Store API Usage

This project uses the public **Fake Store API** as a data source.

Common endpoints used:

- All products

  ```
  GET https://fakestoreapi.com/products
  ```

- Single product
  ```
  GET https://fakestoreapi.com/products/{id}
  ```

---

## 🖼 Template Rendering Example

Go’s built-in `html/template` package is used for rendering dynamic HTML:

```go
tmpl := template.Must(template.ParseFiles("template/index.html"))
tmpl.Execute(w, products)
```

---

## 🚧 Possible Improvements

- Shopping cart functionality
- User sessions and JWT authentication
- Product search and filtering
- Database integration
- Docker support
- Improved UI/UX

## 📄 License

This project is open source and available for educational and personal use.

---
