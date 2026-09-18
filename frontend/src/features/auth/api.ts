import { http } from "@/services/http";
import type { LoginRequest, LoginResponse, MeResponse, RegisterRequest } from "./types";

export async function getCurrentUser() {
    const response = await http.get<MeResponse>("/auth/me");

    return response.data;
}

export async function login(req: LoginRequest) {
    const response = await http.post<LoginResponse>("/auth/login", req);

    return response.data;
}

export async function register(req: RegisterRequest) {
    const response = await http.post<RegisterRequest>("/auth/register", req);

    return response.data;
}