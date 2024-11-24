import axios from "axios";

const API_BASE = process.env.REACT_APP_API_BASE;

export const getTasks = () => axios.get(`${API_BASE}/tasks`);
export const createTask = (task) => axios.post(`${API_BASE}/tasks`, task);
export const updateTask = (id, task) => axios.put(`${API_BASE}/tasks/${id}`, task);
export const deleteTask = (id) => axios.delete(`${API_BASE}/tasks/${id}`);
export const updateTaskStatus = (id, status) => {
  return axios.put(`${API_BASE}/tasks/${id}/status`, { status });
};

