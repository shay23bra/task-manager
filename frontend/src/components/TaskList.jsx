import React from "react";
import "./TaskList.css";

const TaskList = ({ tasks, onDelete }) => (
  <ul>
    {tasks.map((task) => (
      <li key={task.id}>
        <h3>{task.title}</h3>
        <p>{task.description}</p>
        <span>Status: {task.status}</span>
        <button onClick={() => onDelete(task.id)}>Delete</button>
      </li>
    ))}
  </ul>
);

export default TaskList;
