<script setup lang="ts">
    const { selected = false } = defineProps<{
        selected?: boolean,
    }>();
</script>

<template>
    <a class="p-2 bg-white text-neutral-600 rounded-lg cursor-pointer"
        :class="selected ? 'bg-primary-50! text-primary-500!' : ''">
        <slot/>
    </a>
</template>