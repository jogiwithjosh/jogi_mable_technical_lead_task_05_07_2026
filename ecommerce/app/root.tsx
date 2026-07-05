import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration
} from "@remix-run/react";

import { AuthProvider } from "../app/context/AuthProvider";

export default function App() {

    return (

        <html lang="en">

            <head>

                <Meta />

                <Links />

            </head>

            <body>

                <AuthProvider>

                    <Outlet />

                </AuthProvider>

                <ScrollRestoration />

                <Scripts />

                <LiveReload />

            </body>

        </html>

    );

}