import { CartItem as Item } from "../types/cart";

interface Props {

    item: Item;

    onRemove(id: string): void;

    onQuantityChange(
        id: string,
        quantity: number
    ): void;

}

export default function CartItem({

    item,

    onRemove,

    onQuantityChange

}: Props) {

    return (

        <tr>

            <td>{item.product.title}</td>

            <td>

                ${item.product.price}

            </td>

            <td>

                <input

                    type="number"

                    min={1}

                    value={item.quantity}

                    onChange={e =>
                        onQuantityChange(
                            item.product.id,
                            Number(e.target.value)
                        )
                    }

                />

            </td>

            <td>

                $

                {item.product.price *
                    item.quantity}

            </td>

            <td>

                <button

                    onClick={() =>
                        onRemove(
                            item.product.id
                        )
                    }

                >

                    Remove

                </button>

            </td>

        </tr>

    );

}