import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  headers: {
    Authorization: 'Bearer your_static_token_here', // match your backend token
  },
});

export default api;
