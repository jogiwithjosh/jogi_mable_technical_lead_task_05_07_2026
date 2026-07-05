import { useNavigate } from "@remix-run/react";
import { useEffect } from "react";

import { useAuth } from "../hooks/useAuth";

export default function Logout() {

    const navigate = useNavigate();

    const { signOut } = useAuth();

    useEffect(() => {

        signOut();

        navigate("/login", {
            replace: true
        });

    }, [navigate, signOut]);

    return <p>Signing out...</p>;

}