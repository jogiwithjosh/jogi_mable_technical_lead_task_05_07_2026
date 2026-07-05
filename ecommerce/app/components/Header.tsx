import { Link } from "@remix-run/react";

export default function Header() {
    return (
        <header
            style={{
                display: "flex",
                justifyContent: "space-between",
                padding: "1rem",
                borderBottom: "1px solid #ddd"
            }}
        >
            <h2>Mable Store</h2>

            <nav
                style={{
                    display: "flex",
                    gap: "1rem"
                }}
            >
                <Link to="/products">Products</Link>
                <Link to="/cart">Cart</Link>
                <Link to="/logout">Logout</Link>
            </nav>
        </header>
    );
}