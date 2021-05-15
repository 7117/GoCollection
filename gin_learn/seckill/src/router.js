import {createRouter, createWebHistory} from "vue-router"

const webHistory = createWebHistory()

const router = createRouter({
    history: webHistory,
    routes: [
        {
            path: "/home",
            name: "home",
            component: () => import('./components/Home'),
        },
        {
            path: "/",
            name: "index",
            component: () => import('./components/Index'),
        },
        {
            path: "/books",
            name: "books",
            component: () => import('./components/Books'),
        },
        {
            path: "/books_detail/:id",
            name: "books_detail",
            component: () => import('./components/BooksDetail'),
        },
        {
            path: "/books_detail/:id",
            name: "books_detail",
            component: () => import('./components/BooksDetail'),
        },
        {
            path: "/books_test",
            name: "books_test",
            component: () => import('./components/BooksTest'),
        }
    ]
})

export default router
