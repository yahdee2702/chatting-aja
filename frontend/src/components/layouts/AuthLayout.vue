<script setup>
import { useRoute } from 'vue-router'

const route = useRoute();
</script>

<template>
    <main class="w-full h-screen flex flex-row">
        <img class="flex-1 relative overflow-hidden m-4 rounded-2xl object-cover" :src="route.meta.image" alt="authentication image" />
        <RouterView />
    </main>
</template>