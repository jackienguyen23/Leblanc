import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/Home.vue'),
    },
    {
      path: '/menu',
      name: 'menu',
      component: () => import('@/views/Menu.vue'),
    },
    {
      path: '/mood-booker',
      name: 'mood',
      component: () => import('@/views/MoodBooker.vue'),
    },
    {
      path: '/booking',
      name: 'booking',
      component: () => import('@/views/Booking.vue'),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/Login.vue'),
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('@/views/Register.vue'),
    },
  ],
})

export default router
