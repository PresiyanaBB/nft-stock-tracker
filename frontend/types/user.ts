import { ApiResponse } from "./api";

export enum UserRole {
	Collector = "collector",
	Admin = "admin",
}

export type AuthResponse = ApiResponse<{ token: string; user: User }>;

export type User = {
  id: number;
	email: string;
  role: UserRole;
	createdAt: string;
	updatedAt: string;
}
