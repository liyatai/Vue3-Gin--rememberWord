import { createRouter, createWebHashHistory,} from 'vue-router'

const routes = [
    {
        path:'/',
        redirect:'/word'
    },
    {
        path:'/word',
        component:()=>import('../view/word.vue')
    },
    {
        path:'/admin',
        component:()=>import('../view/admin.vue'),
    },
    {
        path:'/add',
        component:()=>import('../view/add.vue')
    },
    {
        path:'/manage',
        component:()=>import('../view/manage.vue')
    },
]

const router = createRouter({
    history:createWebHashHistory(),
    routes:routes
})

export default router