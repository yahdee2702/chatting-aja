<script setup lang="ts">
import { useRouter } from "vue-router";
import { IconMessage } from "@tabler/icons-vue";

import FormInput from "@/components/ui/FormInput.vue"
import FormButton from "@/components/ui/FormButton.vue"

import type { RegisterRequest } from "@/features/auth/types";
import { register } from "@/features/auth/api";
import { ApiError } from "@/services/api_errors";

const router = useRouter();

async function onSubmit(e: SubmitEvent) {
    const form = e.currentTarget as HTMLFormElement
    const formData = new FormData(form);
    
    const data: RegisterRequest = {
        name: formData.get('name') as string,
        email: formData.get('email') as string,
        password: formData.get('password') as string,
        password_confirm: formData.get('password_confirm') as string,
    }

    try {
        const res = await register(data);
        console.log(res.message);

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
    <form class="w-3xl px-9 pb-24 pt-32" v-on:submit.prevent="onSubmit">
        <div class="w-20 h-20 rounded-2xl p-2 bg-primary-50 text-primary-500 mb-1 select-none">
            <IconMessage class="w-full h-full" />
        </div>

        <h1 class="text-3xl text-black font-semibold mb-1">
            Create Your Account
        </h1>

        <p class="text-base text-neutral-700 mb-11">Identify yourself!</p>

        <FormInput label="Username" name="name" placeholder="Input your username" required class="mb-4" />

        <FormInput label="Email" name="email" placeholder="Input your email" required class="mb-4" />

        <FormInput label="Password" name="password" type="password" placeholder="Input your password" required class="mb-4" />

        <FormInput label="Password Confirm" name="password_confirm" type="password" placeholder="Confirm your password" required class="mb-4" />

        <FormButton type="submit" class="mb-4">
            Register
        </FormButton>

        <p class="text-sm text-neutral-800 text-center">
            Already have an account? <a href="/auth/login"
                class="text-primary-500 hover:text-primary-600 font-semibold">Login</a>
        </p>
    </form>
</template>