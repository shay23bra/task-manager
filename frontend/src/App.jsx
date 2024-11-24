import React, { useState, useEffect, useRef } from "react";
import { getTasks, createTask, deleteTask, updateTaskStatus } from "./api/tasks";

const App = () => {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState({ title: "", description: "" });
  const [dropdownTaskId, setDropdownTaskId] = useState(null);
  const dropdownRef = useRef(null);

  useEffect(() => {
    getTasks().then((res) => setTasks(res.data));
  }, []);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        setDropdownTaskId(null);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const handleAddTask = (e) => {
    e.preventDefault();
    if (!newTask.title || !newTask.description) return;
    createTask({ ...newTask, status: "Pending" }).then((res) => {
      setTasks([...tasks, res.data]);
      setNewTask({ title: "", description: "" });
    });
  };

  const handleDeleteTask = (id) => {
    deleteTask(id).then(() => setTasks(tasks.filter((task) => task.id !== id)));
  };

  const handleStatusChange = (id, newStatus) => {
    updateTaskStatus(id, newStatus).then((res) => {
      setTasks(tasks.map((task) => (task.id === res.data.id ? res.data : task)));
      setDropdownTaskId(null);
    });
  };

  const toggleDropdown = (id) => {
    setDropdownTaskId(dropdownTaskId === id ? null : id);
  };

  const toggleDescription = (id) => {
    setTasks(
      tasks.map((task) =>
        task.id === id ? { ...task, expanded: !task.expanded } : task
      )
    );
  };

  return (
    <div className="container">
      <h1>Task Manager</h1>
      <form onSubmit={handleAddTask}>
        <input
          type="text"
          placeholder="Task Title"
          value={newTask.title}
          onChange={(e) => setNewTask({ ...newTask, title: e.target.value })}
          required
        />
        <textarea
          placeholder="Task Description"
          value={newTask.description}
          onChange={(e) =>
            setNewTask({ ...newTask, description: e.target.value })
          }
          required
        ></textarea>
        <button type="submit">Add Task</button>
      </form>
      <ul>
        {tasks.map((task) => (
          <li key={task.id}>
            <div className="task-details">
              <h3>{task.title}</h3>
              <div className="task-details">
                  <h3 data-title={task.title}>{task.title}</h3>
                  <div className="status-dropdown" ref={dropdownRef}>
                  <span
                    className={`status-label ${task.status.toLowerCase()}`}
                    onClick={() => toggleDropdown(task.id)}
                  >
                    {task.status}
                  </span>
                  <ul
                    className={`status-options ${
                      dropdownTaskId === task.id ? "open" : ""
                    }`}
                  >
                    <li
                      className="pending"
                      onClick={() => handleStatusChange(task.id, "Pending")}
                    >
                      Pending
                    </li>
                    <li
                      className="done"
                      onClick={() => handleStatusChange(task.id, "Done")}
                    >
                      Done
                    </li>
                    <li
                      className="hold"
                      onClick={() => handleStatusChange(task.id, "Hold")}
                    >
                      Hold
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <p
              className={`task-description ${
                task.expanded ? "expanded" : ""
              }`}
              onClick={() => toggleDescription(task.id)}
            >
              {task.description}
            </p>
            <button
              className="delete-button"
              onClick={() => handleDeleteTask(task.id)}
            >
              &times;
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default App;
