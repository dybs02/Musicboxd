import { createRouter, createWebHistory } from 'vue-router'
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
      component: () => import('@/views/AlbumView.vue')
    },
    {
      path: '/track/:id',
      name: 'track',
      component: () => import('@/views/TrackView.vue')
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
    },
    {
      path: '/error/:message/:cause',
      name: 'error',
      component: () => import('@/views/ErrorView.vue')
    },
    {
      path: '/loginredirect',
      name: 'loginredirect',
      component: () => import('@/views/LoginRedirectView.vue')
    },
    {
      path: '/admin/reported',
      name: 'reported',
      component: () => import('@/views/admin/ReportedView.vue')
    },
    {
      path: '/diary/:userId',
      name: 'diary',
      component: () => import('@/views/DiaryView.vue')
    },
    {
      path: '/posts',
      name: 'posts',
      component: () => import('@/views/PostsView.vue')
    },
    {
      path: '/posts/:reviewId',
      name: 'post-share-review',
      component: () => import('@/views/PostsView.vue')
    },
    {
      path: '/chat',
      name: 'chat',
      component: () => import('@/views/ChatView.vue')
    },
  ]
})

export default router
