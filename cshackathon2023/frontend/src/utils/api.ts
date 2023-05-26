import Axios from "axios";

export const httpClient = Axios.create({
  baseURL: `http://127.0.0.1:8080/v1`,
  withCredentials: true,
});

export const fetcher = (key: string) =>
  httpClient.get(key).then((res) => res.data);
