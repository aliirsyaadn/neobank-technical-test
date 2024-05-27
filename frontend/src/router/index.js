import * as React from "react";
import * as ReactDOM from "react-dom";
import {
    createBrowserRouter
} from "react-router-dom";

import Homepage from "../pages/Homepage";
import Login from "../pages/Login";
import Register from "../pages/Register";
import Layout from "../components/Layout";
import TransactionCreate from "../pages/TransactionCreate";
  
const router = createBrowserRouter([
    {
        path: "/",
        Component: Layout,
        children: [
            {
                path: "/",
                Component: Homepage,
            },
            {
                path: "/create-transaction",
                Component: TransactionCreate,
            }
        ]
    },
    {
        path: "/login",
        Component: Login,
    },
    {
        path: "/register",
        Component: Register,
    }
]);

export default router