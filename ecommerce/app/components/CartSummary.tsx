import { Link } from "@remix-run/react";
import { Cart } from "../types/cart";

interface Props {
    cart: Cart;
}

export default function CartSummary({
    cart
}: Props) {

    return (

        <div>

            <h2>

                Total

            </h2>

            <h3>

                ${cart.totalAmount}

            </h3>

            <Link to="/checkout">

                Checkout

            </Link>

        </div>

    );

}