import Vue from 'vue'
import Router from 'vue-router'

// material
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

import Vuetify from 'vuetify'

// import Home from './views/Home.vue'
import Vocabulary from './views/Vocabulary'

Vue.use(Router)
Vue.use(VueMaterial)
Vue.use(Vuetify)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/products/vocabulary',
      component: Vocabulary
    }
    // {
    //   path: '/',
    //   name: 'home',
    //   component: Home
    // },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (about.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    // }
  ]
})
