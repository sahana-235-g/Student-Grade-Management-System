import axios from "axios";

// In dev: "" = same origin, Vite proxy forwards /auth & /api to backend
// In prod: set VITE_API_URL or defaults to backend URL
const baseURL = import.meta.env.DEV ? "" : (import.meta.env.VITE_API_URL || "http://localhost:8080");
const api = axios.create({ baseURL });

// Automatically attach JWT token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;