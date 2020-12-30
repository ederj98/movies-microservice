import axios from 'axios';

export const axiosIntance = axios.create({
  timeout: 30000,
});
