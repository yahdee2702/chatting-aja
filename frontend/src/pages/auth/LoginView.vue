<script setup lang="ts">
import { useRouter } from "vue-router";
import { IconMessage } from "@tabler/icons-vue";

import FormInput from "@/components/ui/FormInput.vue"
import FormButton from "@/components/ui/FormButton.vue"

import type { LoginRequest } from "@/features/auth/types";
import { login } from "@/features/auth/api";
import { ApiError } from "@/services/api_errors";
import { useAuthStore } from "@/features/auth/store";

const router = useRouter();

async function onSubmit(e: SubmitEvent) {
    const form = e.currentTarget as HTMLFormElement
    const formData = new FormData(form);
    
    const data: LoginRequest = {
        email: formData.get('email') as string,
        password: formData.get('password') as string,
    }

    try {
        const res = await login(data);
        console.log(res.message);

        const auth = useAuthStore();
        await auth.initialize();

        router.push({name: "home"});
    } catch (e) {
        if (e instanceof ApiError) {
            console.log(e.response?.message);
            console.log(e.response?.errors);
            return
        }

        console.log(e);
    }

}
</script>

<template>
    <form class="w-2xl 2xl:w-3xl px-6 2xl:px-9 pb-24 pt-8 2xl:pt-32" v-on:submit.prevent="onSubmit">
        <div class="w-16 h-16 2xl:w-20 2xl:h-20 rounded-2xl p-2 bg-primary-50 text-primary-500 mb-2 select-none">
            <IconMessage class="w-full h-full" />
        </div>

        <h1 class="text-2xl 2xl:text-3xl text-black font-semibold mb-1">
            Welcome Back!
        </h1>

        <p class="text-sm 2xl:text-base text-neutral-700 mb-8 2xl:mb-11">Start chatting with your friends!</p>

        <FormInput label="Email" name="email" placeholder="Input your email" required class="mb-2 2xl:mb-4" />

        <FormInput label="Password" name="password" type="password" placeholder="Input your password" required class="mb-2 2xl:mb-4" />

        <FormButton type="submit" class="mb-2 2xl:mb-4">
            Login
        </FormButton>

        <p class="text-xs 2xl:text-sm text-neutral-800 text-center">
            Don't have an account? <a href="/auth/register"
                class="text-primary-500 hover:text-primary-600 font-semibold">Register</a>
        </p>
    </form>
</template>