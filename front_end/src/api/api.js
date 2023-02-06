import axios from 'axios';

export const httpClient = axios.create({
  baseURL: '//127.0.0.1:8259',
});
