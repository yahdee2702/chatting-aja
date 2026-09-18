import type { ApiResponse } from "@/types/api"

interface LoginData {
    token: string
}

interface RegisterData {
    token: string
}

export interface UserData {
    id: string
    name: string
    email: string
    createdAt: string
}

export interface LoginRequest {
    email: string
    password: string
}

export interface RegisterRequest {
    name: string
    email: string
    password: string
    password_confirm: string
}

export type LoginResponse = ApiResponse<LoginData>
export type RegisterResponse = ApiResponse<RegisterData>
export type MeResponse = ApiResponse<UserData>