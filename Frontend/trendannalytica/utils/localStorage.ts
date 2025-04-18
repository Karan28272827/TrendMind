import axios from "axios";

export const setLocalStorage = (key: string, value: unknown) => {
    localStorage.setItem(key, JSON.stringify(value));
  };  
  
  export const getLocalStorage = (key: string) => {
    const item = localStorage.getItem(key);
    return item ? JSON.parse(item) : null;
  };
  

const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:5000",
});

// ✅ Load token from localStorage if it exists
const token = localStorage.getItem("token");
if (token) {
    api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
}

export default api;