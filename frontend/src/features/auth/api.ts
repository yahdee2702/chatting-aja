import { http } from "@/services/http";

export async function getCurrentUser() {
    const response = await http.get("/auth/me");

    return response.data;
}