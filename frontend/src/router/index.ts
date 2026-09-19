import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/features/auth/store';

import loginImage from "@/assets/login-page.webp"
import registerImage from "@/assets/register-page.webp"

import AuthLayout from '@/components/layouts/AuthLayout.vue';
import HomeView from '@/pages/HomeView.vue';
import LoginView from '@/pages/LoginView.vue';
import RegisterView from '@/pages/RegisterView.vue';


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
    name: "home",
    component: HomeView,
    meta: {
      requiredAuth: true
    }
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
