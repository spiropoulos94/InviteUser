import axios from "axios";

export const setupApiInstance = (params = {}) => {
  const instance = axios.create(params);

  return instance;
};
