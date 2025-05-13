// services/user.ts
import { AuthResponse } from "@/types/user";
import { Api } from "./api";

async function login(email: string, password: string): Promise<AuthResponse> {
  const response = await Api.post("/auth/login", { email, password });
  return response.data;
}

async function register(email: string, password: string): Promise<AuthResponse> {
  const response = await Api.post("/auth/register", { email, password });
  return response.data;
}

const userService = {
  login,
  register,
};

export { userService };
