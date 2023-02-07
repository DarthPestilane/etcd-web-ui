import axios from 'axios';

let baseUrl = '//127.0.0.1:8259/api';
if (process.env.NODE_ENV === 'production') {
  baseUrl = '/api';
}

export const httpClient = axios.create({
  baseURL: baseUrl,
});
