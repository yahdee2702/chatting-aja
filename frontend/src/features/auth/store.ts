import { defineStore } from "pinia";

export const useAuthStore = defineStore("auth", {
    state: () => ({
        user: null,
        initialized: false
    }),
    getters: {
        isAuthenticated: (state) => state.user !== null,
    },
    actions: {
        async initialize() {
            
        },
    }
});