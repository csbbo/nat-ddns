import Vue from 'vue'
import Router from 'vue-router'
import Login from './components/Login'
import Regist from './components/Regist'
import Main from './components/Main'
import {CheckAuthAPI} from "@/common/api";

Vue.use(Router)

const router = new Router({
    mode: 'history',
    routes: [
        {path: '/login', name: 'login', component: Login},
        {path: '/regist', name: 'regist', component: Regist},
        {path: '/', name: 'main', component: Main, meta:{requireAuth: true}},
    ]
})

router.beforeEach((to, from, next) => {
    if (to.meta.requireAuth === true){
        CheckAuthAPI().then(resp=>{
            if (resp.err === null) {
                next()
            } else {
                next({ path: '/login' })
            }
        }).catch(() => {
            next({path: '/login'})
        })
    } else {
        next()
    }
})

export default router