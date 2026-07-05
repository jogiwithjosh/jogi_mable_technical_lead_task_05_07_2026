import { Cart } from "../types/cart";

interface Props {
    cart: Cart;
}

export default function OrderSummary({
    cart
}: Props) {

    return (

        <div>

            <h2>Order Summary</h2>

            <p>
                Items: {cart.totalItems}
            </p>

            <h3>
                Total: ${cart.totalAmount}
            </h3>

        </div>

    );

}