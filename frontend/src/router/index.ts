import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/features/auth/store';

import loginImage from "@/assets/images/login-page.webp"
import registerImage from "@/assets/images/register-page.webp"

import AuthLayout from '@/components/layouts/AuthLayout.vue';
import ChatLayout from '@/components/layouts/ChatLayout.vue';

import LoginView from '@/pages/auth/LoginView.vue';
import RegisterView from '@/pages/auth/RegisterView.vue';
import EmptyConversation from '@/pages/chats/EmptyConversation.vue';
import ChatsConversation from '@/pages/chats/ChatsConversation.vue';


const routes = [
  {
    path: '/auth',
    component: AuthLayout,
    children: [
      {
        path: 'login',
        name: 'login',
        meta: {
          image: loginImage,
        },
        component: LoginView,
      },
      {
        path: 'register',
        name: 'register',
        meta: {
          image: registerImage,
        },
        component: RegisterView,
      },
    ],
  },
  {
    path: "/",
    component: ChatLayout,
    meta: {
      requiredAuth: true
    },
    children: [
      {
        path: "/",
        name: "home",
        component: EmptyConversation,
      },
      {
        path: "/",
        name: "chats",
        component: ChatsConversation,
      },
      {
        path: "/",
        name: "servers",
        component: EmptyConversation,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})

router.beforeEach(async (to, _) => {
  const auth = useAuthStore();

  if (!auth.initialized) {
    await auth.initialize();
  }

  const requiredAuth = to.matched.some(record => {
    return record.meta.requiredAuth;
  });

  if (requiredAuth && !auth.isAuthenticated) {
    return {name: "login"};
  } else if ((to.name === "login" || to.name === "register") && auth.isAuthenticated) {
    return {name: "home"};
  } else {
    return true;
  }
})

export default router
