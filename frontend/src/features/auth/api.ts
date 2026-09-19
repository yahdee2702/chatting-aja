import { http } from "@/services/http";
import type { LoginRequest, LoginResponse, MeResponse, RegisterRequest, RegisterResponse } from "./types";

export async function getCurrentUser() {
    const response = await http.get<MeResponse>("/api/auth/me");

    return response.data;
}

export async function login(req: LoginRequest) {
    const response = await http.post<LoginResponse>("/api/auth/login", req);

    return response.data;
}

export async function register(req: RegisterRequest) {
    const response = await http.post<RegisterResponse>("/api/auth/register", req);

    return response.data;
}