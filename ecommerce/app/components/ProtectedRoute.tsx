import { useNavigate } from "@remix-run/react";
import { useEffect } from "react";
import { isAuthenticated } from "../services/auth";

export default function ProtectedRoute({
    children
}: {
    children: React.ReactNode;
}) {

    const navigate = useNavigate();

    useEffect(() => {

        if (!isAuthenticated()) {

            navigate("/login");

        }

    }, []);

    if (!isAuthenticated()) {

        return null;

    }

    return <>{children}</>;
}