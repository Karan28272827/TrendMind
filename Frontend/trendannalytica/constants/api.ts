// api.ts
import axios from 'axios';

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080/auth', // fallback
  withCredentials: true // optional if you're using cookies/sessions
});

export default api;
