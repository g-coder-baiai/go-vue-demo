import axios from 'axios';
import storageService from '../service/storageService';

const service = axios.create({
  // process是nodejs维护的一个全局变量，它描述当前nodejs的进程信息
  // env保存着当前node进程中所有的环境变量
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 1000 * 5,
  // headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` },
});

// Add a request interceptor
service.interceptors.request.use((config) => {
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => {
  return Promise.reject(error);
});

export default service;
