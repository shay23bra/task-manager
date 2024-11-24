
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

## **Components**

### **Frontend Components**
1. **`App.jsx`**:
   - Main React component managing task list, dropdowns, and API interactions.
2. **`api/tasks.js`**:
   - Contains API methods (`getTasks`, `createTask`, `updateTaskStatus`, `deleteTask`) to communicate with the backend.
3. **`styles/index.css`**:
   - Handles application-wide styling.

### **Backend Components**
1. **`main.go`**:
   - Configures routes and database and starts the server.
2. **Endpoints**:
   - `GET /tasks`
   - `POST /tasks`
   - `PUT /tasks/:id/status`
   - `DELETE /tasks/:id`

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

---
