import AuthLayout from '@/components/layouts/AuthLayout.vue';
import { useAuthStore } from '@/features/auth/store';
import HomeView from '@/pages/HomeView.vue';
import LoginView from '@/pages/LoginView.vue';
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/auth',
    component: AuthLayout,
    children: [
      {
        path: 'login',
        name: 'login',
        component: LoginView,
      },
    ],
  },
  {
    path: "",
    name: "home",
    component: HomeView,
    meta: {requiredAuth: true}
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})

router.beforeEach((to, _, next) => {
  const auth = useAuthStore();

  if (!auth.initialized) {
    auth.initialize();
  }

  const requiredAuth = to.matched.some(record => {
    console.log(record.name);
    return record.meta.requiredAuth;
  });

  if (requiredAuth && !auth.isAuthenticated) {
    next({name: "Login"})
  } else if (to.name === "Login" && auth.isAuthenticated) {
    next({name: "Home"})
  } else {
    next();
  }
})

export default router
