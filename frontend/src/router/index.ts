import { createRouter, createWebHistory } from 'vue-router'
import AlbumView from '@/views/AlbumView.vue'
import HomeView from '@/views/HomeView.vue'
import ReviewView from '@/views/ReviewView.vue'
import UserProfile from '@/views/UserProfile.vue'


const router = createRouter({
  // @ts-ignore
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/album/:id',
      name: 'album',
      component: AlbumView
    },
    {
      path: '/track/:id',
      name: 'track',
      component: AlbumView // TODO separete view for track or one for both ?
    },
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/review/:id',
      name: 'review',
      component: ReviewView
    },
    {
      path: '/user/:id',
      name: 'user',
      component: UserProfile
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('@/views/AboutView.vue')
    }
  ]
})

export default router
