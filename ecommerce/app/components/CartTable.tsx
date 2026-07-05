import { Cart } from "../types/cart";
import CartItem from "./CartItem";

interface Props {

    cart: Cart;

    onRemove(id: string): void;

    onQuantityChange(
        id: string,
        quantity: number
    ): void;

}

export default function CartTable({

    cart,

    onRemove,

    onQuantityChange

}: Props) {

    return (

        <table width="100%">

            <thead>

                <tr>

                    <th>Product</th>

                    <th>Price</th>

                    <th>Qty</th>

                    <th>Total</th>

                    <th></th>

                </tr>

            </thead>

            <tbody>

                {cart.items.map(item => (

                    <CartItem

                        key={item.product.id}

                        item={item}

                        onRemove={onRemove}

                        onQuantityChange={
                            onQuantityChange
                        }

                    />

                ))}

            </tbody>

        </table>

    );

}