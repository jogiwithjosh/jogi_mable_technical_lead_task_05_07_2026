import { useNavigate } from "@remix-run/react";
import { useState } from "react";

import CheckoutForm from "../components/CheckoutForm";
import Layout from "../components/Layout";
import OrderSummary from "../components/OrderSummary";
import PaymentMethod from "../components/PaymentMethod";

import { clearCart, getCart } from "../services/cart";
import { placeOrder } from "../services/checkout";

import { Analytics } from "../lib/tracker";

export default function Checkout() {

    const navigate = useNavigate();

    const cart = getCart();

    const [loading, setLoading] =
        useState(false);

    const [paymentMethod, setPaymentMethod] =
        useState("Credit Card");

    async function handleCheckout() {

        setLoading(true);

        try {

            Analytics.checkout(
                "cart-1",
                cart.totalAmount
            );

            Analytics.paymentInfoAdded(
                paymentMethod
            );

            const order =
                await placeOrder({

                    items: cart.items,

                    paymentMethod

                });

            Analytics.purchase(
                order.orderId,
                order.total
            );

            clearCart();

            navigate("/success");

        } finally {

            setLoading(false);

        }

    }

    return (

        <Layout>

            <h1>Checkout</h1>

            <OrderSummary
                cart={cart}
            />

            <PaymentMethod
                value={paymentMethod}
                onChange={setPaymentMethod}
            />

            <CheckoutForm
                loading={loading}
                onSubmit={handleCheckout}
            />

        </Layout>

    );

}