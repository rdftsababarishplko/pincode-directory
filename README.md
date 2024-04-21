Welcome to the Pincode Directory project! This project is a simple web application built in Go to manage pincode data. It allows users to view, create, update, and delete pincode entries stored in a SQLite database.

## Features

- View a list of pincode entries
- Create new pincode entries
- Update existing pincode details
- Delete pincode entries

## Prerequisites

Before running this application, ensure you have the following installed on your system:

- Go programming language (v1.16+)
- SQLite database

## Setup

1. **Clone the Repository**
   ```bash
   git clone https://github.com/rdftsababarishplko/pincode-directory.git
   ```

2. **Navigate to the Project Directory**
   ```bash
   cd pincode-directory
   ```

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the Application**
   ```bash
   go run main.go
   ```

5. **Access the Application**
   Open your web browser and go to `http://localhost:8080` to view the application.

## Project Structure

```
pincode-directory/
├── main.go            # Main application entry point
├── go.mod             # Go module file
├── go.sum             # Go dependencies checksum
├── static/            # Static assets (CSS, JS, images)
├── templates/         # HTML templates
│   ├── index.html
│   ├── read_pincode.html
│   ├── create_pincode.html
│   ├── update_pincode.html
│   └── delete_pincode.html
└── data.db            # SQLite database file
```

## Contributing

Contributions are welcome! If you have any ideas, improvements, or bug fixes, please feel free to open an issue or submit a pull request.
