import Home from '../views/Home.vue'
import Main from '../layout/index'

const routers = [
    {
        path: '/login',
        name: 'Login',
        meta: {
            title: '登录',
            icon: ''
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/login/login.vue')
    },
    {
      path: '/',
      name: 'Main',
      redirect: '/home',
      meta: {
        name: 'Main',
        title: 'Main'
      },
      component: Main,
      children: [
          {
              path: '/home',
              name: 'Home',
              component: Home
          }
      ]
      
    },
    {
      path: '/about',
      name: 'About',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    }
]


export default routers