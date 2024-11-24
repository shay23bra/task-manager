
# **Task Manager Application**

A full-stack **Task Manager Application** that allows users to manage tasks with features to create, update, delete, and organize tasks by status. The application is divided into a **Golang backend** and a **React frontend**.

---

## **Project Overview**

### **Frontend**
The frontend is built using **React** and provides a user-friendly interface for managing tasks. It includes:
- **Task List**: Displays all tasks with title, description, and status.
- **Task Actions**: Add new tasks, update their status, and delete tasks.
- **Status Dropdown**: A smooth dropdown interface for changing the status of a task.
- **Responsive Design**: Ensures a clean and readable interface.

### **Backend**
The backend is built using **Golang** with **Gin Framework** and provides a RESTful API for managing tasks. It includes:
- **Endpoints**:
  - `GET /tasks`: Retrieve all tasks.
  - `POST /tasks`: Create a new task.
  - `PUT /tasks/:id/status`: Update the status of a task.
  - `DELETE /tasks/:id`: Delete a task.
- **Database**:
  - Uses an **in-memory SQLite database** for lightweight storage.

---

## **Project Structure**

### **Backend**
```
backend/
├── main.go            # Entry point of the backend application
├── go.mod             # Module dependencies
├── db/
│   └── tasks.db       # SQLite database (created automatically)
├── models/
│   └── task.go        # Task model definition
├── routes/
│   └── tasks.go       # API routes for task management
├── Dockerfile         # Dockerfile for backend service
```

### **Frontend**
```
frontend/
├── public/
│   └── index.html     # Main HTML template
├── src/
│   ├── App.jsx        # Main React component
│   ├── index.jsx      # Entry point for React
│   ├── api/
│   │   └── tasks.js   # API service for interacting with the backend
│   ├── styles/
│   │   └── index.css  # Main CSS file for the application
│   ├── components/    # Reusable React components (optional future use)
│   │   ├── Task.js     # Task component for displaying individual tasks
│   │   └── TaskList.js # TaskList component for displaying all tasks
├── package.json       # NPM configuration and dependencies
├── Dockerfile         # Dockerfile for frontend service
```

---

## **Setup Instructions**

### **Backend**

1. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Run the Backend**:
   ```bash
   go run main.go
   ```
   The backend will start on `http://localhost:8080`.

3. **Test Backend API**:
   Use a tool like **Postman** or **cURL** to test API endpoints. Examples:
   ```bash
   # Create a new task
   curl -X POST http://localhost:8080/tasks    -H "Content-Type: application/json"    -d '{"title":"Sample Task", "description":"Description of the task"}'

   # Get all tasks
   curl http://localhost:8080/tasks

   # Update task status
   curl -X PUT http://localhost:8080/tasks/1/status    -H "Content-Type: application/json"    -d '{"status":"Done"}'

   # Delete a task
   curl -X DELETE http://localhost:8080/tasks/1
   ```

---

### **Frontend**

1. **Install Dependencies**:
   Navigate to the `frontend/` directory and run:
   ```bash
   npm install
   ```

2. **Run the Frontend**:
   ```bash
   npm start
   ```
   The frontend will be available at `http://localhost:3000`.

3. **Configure API Endpoint**:
   Ensure the backend API URL is set in the `frontend/.env` file:
   ```
   REACT_APP_API_BASE=http://localhost:8080
   ```

4. **Test the Frontend**:
   Open a browser and navigate to `http://localhost:3000` to interact with the application.

---

## **Docker Deployment**

This project is fully containerized using Docker.

### **Steps to Run Using Docker Compose**

1. **Build and Start the Services**:
   In the root directory of the project (where `docker-compose.yml` is located), run:
   ```bash
   docker-compose up --build
   ```

2. **Access the Services**:
   - Frontend: [http://localhost:3000](http://localhost:3000)
   - Backend: [http://localhost:8080](http://localhost:8080)

3. **Stop the Services**:
   ```bash
   docker-compose down
   ```

---

## **CI/CD Pipeline**

A CI/CD pipeline is implemented using GitHub Actions.

### **Pipeline Workflow**

1. **Trigger**:
   - The pipeline runs on every pull request to the `main` branch and on direct pushes to the `main` branch.

2. **Stages**:
   - **Testing**:
     - Runs backend tests using `go test`.
     - Runs frontend tests using `npm test`.
   - **Linting**:
     - Lints the backend code with `go vet`.
     - Lints the frontend code with `npm run lint`.
   - **Build**:
     - Builds Docker images for the backend and frontend.

### **File: .github/workflows/ci.yml**
The CI pipeline is defined in `.github/workflows/ci.yml` and includes running tests, linting, and building Docker images.

---

## **Testing**

### **Backend Testing**
Run unit tests for the backend using the Go `testing` package. Ensure that you’ve written tests for each API endpoint.

Example:
```bash
go test ./... -v
```

### **Frontend Testing**
Run frontend tests using **Jest** and **React Testing Library**.

1. **Run Frontend Tests**:
   ```bash
   npm test
   ```

2. **Example Test Cases**:
   - Test if tasks are correctly displayed.
   - Test if new tasks are added successfully.
   - Test if dropdown animations and status changes work.

---

## **Improvements and Future Enhancements**
1. **Authentication**:
   - Add user authentication to secure task management.
2. **Persistent Database**:
   - Replace in-memory SQLite with a persistent database like PostgreSQL or MySQL.
3. **Search and Filters**:
   - Add search functionality and status-based filters in the UI.
4. **Mobile Responsiveness**:
   - Improve styling for smaller screens.
5. **Packages and Dependencies**:
   - Use newer versions of packages and dependencies, remove deprecated ones.

---
