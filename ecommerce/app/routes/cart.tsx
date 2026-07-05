import { useEffect, useState } from "react";

import CartSummary from "../components/CartSummary";
import CartTable from "../components/CartTable";
import Layout from "../components/Layout";

import {
    getCart,
    removeFromCart,
    updateQuantity
} from "../services/cart";

import { Analytics } from "../lib/tracker";
import { Cart } from "../types/cart";

export default function CartPage() {

    const [cart, setCart] =
        useState<Cart>(getCart());

    useEffect(() => {

        Analytics.event(
            "CartViewed",
            {
                total: cart.totalAmount,
                items: cart.totalItems
            }
        );

    }, []);

    function refresh() {

        setCart(
            getCart()
        );

    }

    function handleRemove(
        id: string
    ) {

        removeFromCart(id);

        refresh();

    }

    function handleQuantity(
        id: string,
        quantity: number
    ) {

        updateQuantity(
            id,
            quantity
        );

        refresh();

    }

    return (

        <Layout>

            <h1>

                Shopping Cart

            </h1>

            <CartTable

                cart={cart}

                onRemove={handleRemove}

                onQuantityChange={
                    handleQuantity
                }

            />

            <CartSummary

                cart={cart}

            />

        </Layout>

    );

}