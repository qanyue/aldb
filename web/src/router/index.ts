import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
    {

        path: "/index",
        name: "Home",
        component: () => import('../views/Home.vue'),
        children: [
            {
                path: "user",
                name: "User",
                component: () => import('../views/Person.vue'),
            },
            {
                path: 'DataSets',
                name: 'Datasets',
                component: () => import('../components/Main.vue'),
            }
        ],
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
        path: '/Main',
        name: 'Main',
        component: () => import('@/components/Main.vue')
    }

];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;
