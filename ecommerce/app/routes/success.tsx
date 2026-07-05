import { Link } from "@remix-run/react";
import Layout from "../components/Layout";

export default function Success() {

    return (

        <Layout>

            <h1>
                🎉 Thank You!
            </h1>

            <p>
                Your order has been placed successfully.
            </p>

            <Link to="/products">
                Continue Shopping
            </Link>

        </Layout>

    );

}