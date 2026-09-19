<script setup lang="ts">
import { IconEye, IconEyeClosed } from "@tabler/icons-vue";
import { ref } from "vue";

const showPassword = ref(false);

defineProps<{
    name: string,
    label: string,
    placeholder?: string,
    type?: string,
    required?: boolean,
}>();

const getType = (type?: string) =>
    type === 'password'
        ? (showPassword.value ? 'text' : type)
        : (type ?? 'text');

</script>

<template>
    <div>
        <label :id="'input-label-' + name" :for="'input-' + name" class="block text-base text-black font-medium mb-1">{{ label }}</label>
        <div class="relative w-full h-11">
            <input :id="'input-' + name" :type="getType(type)" :name="name" :placeholder="placeholder" :required="required"
                class="absolute inset-y-0 left-0 w-full h-11 px-3 py-2.5 bg-white autofill:bg-red-400 border border-neutral-400 focus:border-primary-500 rounded-lg text-sm text-black placeholder:text-neutral-700 outline-none appearance-none"
                :class="type == 'password' ? 'pr-12' : ''">
            <div v-if="type === 'password'"
                class="absolute inset-y-0 right-0 flex items-center px-3 pointer-events-none">
                <div class="w-6 h-6 text-neutral-700 pointer-events-auto cursor-pointer select-none">
                    <IconEye class="w-full h-full" v-if="showPassword" @click="showPassword = false" />
                    <IconEyeClosed class="w-full h-full" v-if="!showPassword" @click="showPassword = true" />
                </div>
            </div>
        </div>
    </div>
</template>