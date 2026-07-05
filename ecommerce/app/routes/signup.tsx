import { Link, useNavigate } from "@remix-run/react";
import { FormEvent, useState } from "react";

import { signup } from "../services/auth";

export default function Signup() {

    const navigate = useNavigate();

    const [name, setName] = useState("");

    const [email, setEmail] = useState("");

    const [password, setPassword] = useState("");

    const [loading, setLoading] = useState(false);

    const [error, setError] = useState("");

    async function handleSubmit(
        e: FormEvent<HTMLFormElement>,
    ) {

        e.preventDefault();

        setError("");

        setLoading(true);

        try {

            await signup({

                name,
                email,
                password,
            });

            navigate("/login");

        } catch (err) {

            if (err instanceof Error) {

                setError(err.message);

            } else {

                setError("Signup failed");
            }

        } finally {

            setLoading(false);
        }
    }

    return (

        <main
            style={{
                maxWidth: 420,
                margin: "80px auto",
                padding: 24,
                border: "1px solid #ddd",
                borderRadius: 8,
            }}
        >

            <h1>Create Account</h1>

            <form onSubmit={handleSubmit}>

                <div style={{ marginBottom: 16 }}>

                    <label>Name</label>

                    <input
                        value={name}
                        required
                        onChange={(e) =>
                            setName(e.target.value)
                        }
                        style={{
                            width: "100%",
                            padding: 8,
                            marginTop: 4,
                        }}
                    />

                </div>

                <div style={{ marginBottom: 16 }}>

                    <label>Email</label>

                    <input
                        type="email"
                        value={email}
                        required
                        onChange={(e) =>
                            setEmail(e.target.value)
                        }
                        style={{
                            width: "100%",
                            padding: 8,
                            marginTop: 4,
                        }}
                    />

                </div>

                <div style={{ marginBottom: 16 }}>

                    <label>Password</label>

                    <input
                        type="password"
                        value={password}
                        required
                        minLength={8}
                        onChange={(e) =>
                            setPassword(e.target.value)
                        }
                        style={{
                            width: "100%",
                            padding: 8,
                            marginTop: 4,
                        }}
                    />

                </div>

                {error && (

                    <p
                        style={{
                            color: "red",
                        }}
                    >
                        {error}
                    </p>

                )}

                <button
                    type="submit"
                    disabled={loading}
                    style={{
                        width: "100%",
                        padding: 10,
                    }}
                >
                    {loading
                        ? "Creating..."
                        : "Create Account"}
                </button>

            </form>

            <p
                style={{
                    marginTop: 20,
                }}
            >
                Already have an account?{" "}
                <Link to="/login">
                    Login
                </Link>
            </p>

        </main>
    );
}