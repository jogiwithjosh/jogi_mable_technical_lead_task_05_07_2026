import { useEffect, useState } from "react";

import Layout from "../components/Layout";
import ProductGrid from "../components/ProductGrid";

import { Product } from "../types/products";

import { useNavigate } from "@remix-run/react";
import { Analytics } from "../lib/tracker";
import * as Auth from "../services/auth";
import { addToCart } from "../services/cart";
import { getProducts } from "../services/product";



export default function Products() {

    const navigate = useNavigate();

    const [products, setProducts] =
        useState<Product[]>([]);

    const [loading, setLoading] =
        useState(true);

    useEffect(() => {

        Analytics.page();

        if (!Auth.isAuthenticated()) {

        navigate("/login", {
            replace: true
        });

        return;
    }
        loadProducts();

    }, [navigate]);

    async function loadProducts() {

        try {

            const result =
                await getProducts();

            setProducts(result);

        } finally {

            setLoading(false);

        }

    }

    function handleAdd(
        product: Product
    ) {
        addToCart(product);

        console.log(
            "Added",
            product.name
        );

    }

    if (loading) {

        return (
            <Layout>

                <p>Loading...</p>

            </Layout>
        );

    }

    return (

        <Layout>

            <h1>Products</h1>

            <ProductGrid

                products={products}

                onAdd={handleAdd}

            />

        </Layout>

    );

}