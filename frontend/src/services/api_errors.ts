import type { ApiErrorResponse } from "@/types/api"

export class ApiError extends Error {
    statusCode: number
    response?: ApiErrorResponse

    constructor(
        message: string,
        statusCode: number,
        response ?: ApiErrorResponse,
    ) {
        super(message)

        this.name = 'ApiError'
        this.statusCode = statusCode
        this.response = response
    }
}