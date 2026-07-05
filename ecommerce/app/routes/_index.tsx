import { useNavigate } from "@remix-run/react";
import { useEffect } from "react";

import * as Auth from "../services/auth";

export default function Index() {

    const navigate = useNavigate();

    useEffect(() => {

        if (Auth.isAuthenticated()) {

            navigate("/products", {
                replace: true
            });

        } else {

            navigate("/login", {
                replace: true
            });

        }

    }, [navigate]);

    return <p>Loading...</p>;

}