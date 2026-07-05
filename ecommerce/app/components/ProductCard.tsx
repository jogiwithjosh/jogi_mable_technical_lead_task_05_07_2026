import { Analytics } from "../lib/tracker";
import { Product } from "../types/products";

interface Props {
    product: Product;
    onAdd: (product: Product) => void;
}

export default function ProductCard({
    product,
    onAdd
}: Props) {

    function handleAdd() {

        Analytics.addToCart(
            product.id,
            1,
            product.price
        );

        onAdd(product);
    }

    return (
        <div
            style={{
                border: "1px solid #ddd",
                padding: 20,
                borderRadius: 8
            }}
        >
            <img
                src={product.image}
                width={180}
                alt={product.title}
            />

            <h3>{product.title}</h3>

            <p>{product.description}</p>

            <strong>
                ${product.price}
            </strong>

            <br />

            <button
                onClick={handleAdd}
            >
                Add to Cart
            </button>
        </div>
    );
}