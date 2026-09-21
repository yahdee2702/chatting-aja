<script setup lang="ts">

import { logout } from '@/features/auth/api';
import { useAuthStore } from '@/features/auth/store';
import ConversationItem from '@/features/chats/components/ConversationItem.vue';
import ConversationNav from '@/features/chats/components/ConversationNav.vue';
import router from '@/router';
import { ApiError } from '@/services/api_errors';
import { IconSearch, IconMessageUser, IconMessage, IconBubble, IconLogout } from '@tabler/icons-vue';

async function requestLogout() {
    try {
        const res = await logout();
        console.log(res.message);

        const auth = useAuthStore();
        await auth.initialize();

        router.push({ name: "login" });
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
    <main class="w-full h-screen flex flex-row bg-neutral-100 gap-4 p-4">
        <div
            class="flex flex-col h-full w-20 py-4 px-8 items-center justify-between bg-white border border-neutral-500 rounded-xl">
            <div></div>
            <nav class="flex flex-col items-center gap-4">
                <ConversationNav selected>
                    <IconMessageUser class="w-8 h-8" />
                </ConversationNav>

                <ConversationNav>
                    <IconBubble class="w-8 h-8" />
                </ConversationNav>

                <ConversationNav>
                    <IconMessage class="w-8 h-8" />
                </ConversationNav>
            </nav>
            <button @click="requestLogout"
                class="p-2 bg-red-50 text-red-400 hover:bg-red-100 hover:text-red-500 rounded-lg cursor-pointer">
                <IconLogout class="w-8 h-8" />
            </button>
        </div>

        <div class="flex flex-col w-md h-full gap-4">
            <div class="flex flex-col px-4 pt-6 pb-3 bg-white border border-neutral-500 rounded-xl">
                <h1 class="text-2xl font-semibold text-primary-700 pb-3">Chat</h1>
                <form class="relative w-full h-11" autocomplete="false" v-on:submit.prevent="">
                    <input type="text" placeholder="Search for someone"
                        class="absolute inset-y-0 left-0 w-full px-3 pl-11 py-2.5 bg-neutral-100 border border-primary-200 focus:border-primary-500 rounded-lg text-sm text-primary-700 placeholder:text-primary-400 outline-none appearance-none">
                    <div class="absolute inset-y-0 left-0 flex items-center px-3 pointer-events-none">
                        <IconSearch class="w-6 h-6 text-primary-500" />
                    </div>
                </form>
            </div>

            <div class="flex flex-col w-full flex-1 bg-white border border-neutral-500 rounded-xl gap-6 px-4 py-6">
                <ConversationItem />
            </div>
        </div>

        <div class="flex flex-col flex-1 h-full gap-4">
            <RouterView />
        </div>
    </main>
</template>