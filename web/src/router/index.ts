import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
    {

        path: "/index",
        name: "Home",
        component: () => import('../views/Home.vue'),
        children: [
            {
                path: '',
                redirect: 'user',
            },

            {
                path: "user",
                name: "User",
                component: () => import('../views/Person.vue'),
            },
            {
                path: 'DataSet',
                name: 'Dataset',
                component: () => import('../views/DataSet.vue'),
            },
            {
                path: 'TagSet',
                name: 'TagSet',
                component: () => import('@/views/TagSet.vue')
            }
        ],
    },
    {
        path:"/annotation",
        name:"annotation",
        component: ()=> import('../components/Annotaion.vue')
    },
    {
        path: "/",
        name: "root",
        redirect: '/login',
        children: [
            {
                path: "login",
                name: "Login",
                component: () => import('../views/Login.vue'),
            }
        ]
    },
    {
        path: '/:catchAll(.*)',
        name: '404',
        component: () => import('../views/404.vue')
    },
    {
        path: '/algae/:riverId',
        name: 'Main',
        component: () => import('@/components/Main.vue')
    }

];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;
