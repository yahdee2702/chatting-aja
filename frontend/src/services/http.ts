import axios from "axios"
import { ApiError } from "./api_errors";
import type { ApiErrorResponse } from "@/types/api";

export const http = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
    withCredentials: true,
});

http.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response) {
            const statusCode = error.response.status
            const data = error.response.data as ApiErrorResponse

            throw new ApiError(
                data.message ?? "Something went wrong",
                statusCode,
                data
            )
        }

        throw new ApiError(
            "Unable to connect to server",
            0,
        )
    }
);