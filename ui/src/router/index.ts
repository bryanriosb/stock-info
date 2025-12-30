import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/stocks'
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/auth/LoginView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('@/views/auth/RegisterView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/stocks',
      name: 'Stocks',
      component: () => import('@/views/stocks/StocksListView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/stocks/:id',
      name: 'StockDetail',
      component: () => import('@/views/stocks/StockDetailView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/recommendations',
      name: 'Recommendations',
      component: () => import('@/views/recommendations/RecommendationsView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/users',
      name: 'Users',
      component: () => import('@/views/users/UsersListView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/NotFoundView.vue')
    }
  ]
})

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('token')
  const isAuthenticated = !!token

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'Login', query: { redirect: to.fullPath } })
  } else if (to.meta.requiresGuest && isAuthenticated) {
    next({ name: 'Stocks' })
  } else {
    next()
  }
})

export default router
