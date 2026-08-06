import api from "../lib/api";

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  full_name: string;
  username: string;
  email: string;
  password: string;
}

export const login = (data: LoginRequest) =>
  api.post("/auth/login", data);

export const register = (data: RegisterRequest) =>
  api.post("/auth/register", data);

export const me = () =>
  api.get("/me");
