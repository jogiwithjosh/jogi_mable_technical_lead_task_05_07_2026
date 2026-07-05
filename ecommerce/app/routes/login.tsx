import { Link, useNavigate } from "@remix-run/react";
import { FormEvent, useEffect, useState } from "react";

import { useAuth } from "../hooks/useAuth";

export default function Login() {

    const navigate = useNavigate();

    const {
        signIn,
        authenticated,
        loading
    } = useAuth();

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const [submitting, setSubmitting] = useState(false);
    const [error, setError] = useState("");

    // Redirect authenticated users
    useEffect(() => {

        if (!loading && authenticated) {

            navigate("/products", {
                replace: true
            });

        }

    }, [
        authenticated,
        loading,
        navigate
    ]);

    async function handleSubmit(
        e: FormEvent<HTMLFormElement>
    ) {

        e.preventDefault();

        setError("");
        setSubmitting(true);

        try {

            await signIn(
                email,
                password
            );

            navigate("/products", {
                replace: true
            });

        } catch (err) {

            if (err instanceof Error) {

                setError(err.message);

            } else {

                setError("Login failed");

            }

        } finally {

            setSubmitting(false);

        }

    }

    if (loading) {

        return (
            <main
                style={{
                    maxWidth: 420,
                    margin: "80px auto",
                    textAlign: "center"
                }}
            >
                <p>Loading...</p>
            </main>
        );

    }

    return (

        <main
            style={{
                maxWidth: 420,
                margin: "80px auto",
                padding: 24,
                border: "1px solid #ddd",
                borderRadius: 8
            }}
        >

            <h1>Login</h1>

            <form onSubmit={handleSubmit}>

                <div style={{ marginBottom: 16 }}>

                    <label htmlFor="email">

                        Email

                    </label>

                    <input
                        id="email"
                        type="email"
                        required
                        value={email}
                        onChange={(e) =>
                            setEmail(e.target.value)
                        }
                        style={{
                            width: "100%",
                            padding: 8,
                            marginTop: 4
                        }}
                    />

                </div>

                <div style={{ marginBottom: 16 }}>

                    <label htmlFor="password">

                        Password

                    </label>

                    <input
                        id="password"
                        type="password"
                        required
                        value={password}
                        onChange={(e) =>
                            setPassword(e.target.value)
                        }
                        style={{
                            width: "100%",
                            padding: 8,
                            marginTop: 4
                        }}
                    />

                </div>

                {error && (

                    <p
                        style={{
                            color: "red",
                            marginBottom: 16
                        }}
                    >
                        {error}
                    </p>

                )}

                <button
                    type="submit"
                    disabled={submitting}
                    style={{
                        width: "100%",
                        padding: 10,
                        cursor: submitting
                            ? "not-allowed"
                            : "pointer"
                    }}
                >

                    {submitting
                        ? "Signing in..."
                        : "Login"}

                </button>

            </form>

            <p style={{marginTop: 20,}}>

                Don't have an account?{" "}

                <Link to="/signup">Create one</Link>

            </p>

        </main>


    );

}