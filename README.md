# go_projects

This repository contains a collection of Go projects for learning, experimentation, and demonstration purposes. Each subfolder in this repository is a standalone Go project with its own source code and resources.

## Projects

### 1. web_server

A simple HTTP web server written in Go. This project demonstrates how to serve static files and handle basic form submissions.

#### Features

- Serves static files from the `static/` directory (e.g., HTML pages).
- Handles form submissions at `/form` and displays submitted data.
- Provides a `/hello` endpoint that returns a simple greeting.

#### Getting Started

1. **Navigate to the project directory:**
    ```sh
    cd web_server
    ```

2. **Run the server:**
    ```sh
    go run main.go
    ```

3. **Access the web server:**
    - Open [http://localhost:8080](http://localhost:8080) in your browser to view the static site.
    - Visit [http://localhost:8080/form.html](http://localhost:8080/form.html) to try the sample form.
    - Access [http://localhost:8080/hello](http://localhost:8080/hello) for a simple greeting.

#### File Structure

```
web_server/
    main.go          # Source code for the web server
    static/
        index.html   # Homepage
        form.html    # Sample form page
```

---

## Adding More Projects

To add a new Go project, create a new subfolder in the root directory and add your Go source files there. Update this README with a brief description of your new project.

---

## License

This repository is for educational purposes. See individual project folders for additional details.
