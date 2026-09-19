import { defineStore } from "pinia";
import { getCurrentUser } from "./api";
import { ApiError } from "@/services/api_errors";
import type { UserData } from "./types";

interface AuthState {
    user: UserData | null,
    initialized: boolean
}

export const useAuthStore = defineStore("auth", {
    state: (): AuthState => ({
        user: null,
        initialized: false
    }),
    getters: {
        isAuthenticated: (state) => state.user != null,
    },
    actions: {
        async initialize() {
            try {
                const response = await getCurrentUser();
                this.user = response.data;
                this.initialized = true;
            } catch (e) {
                if (e instanceof ApiError) {
                    if (e.statusCode === 401) {
                        this.initialized = true;
                        this.user = null;
                        return
                    }

                    this.initialized = false;
                    console.log(e.message);
                    return
                }

                this.initialized = false;
                console.log(e);
            }
        },
    }
});