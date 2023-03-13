import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
    {
        path: "/",
        name: "Home",
        component: () => import('../views/Home.vue'),
    },
    {
        path: "/login",
        name: "Login",
        component: () => import('../views/Login.vue'),
    },
    {
        path: "/user",
        name: "User",
        component: () => import('../views/Person.vue'),
    },
    {
        path: '/:catchAll(.*)',
        name: '404',
        component: () => import('../views/404.vue')
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;
