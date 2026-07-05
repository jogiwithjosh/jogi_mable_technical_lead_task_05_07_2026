import { Product } from "../types/products";
import ProductCard from "./ProductCard";

interface Props {
    products: Product[];
    onAdd(product: Product): void;
}

export default function ProductGrid({
    products,
    onAdd
}: Props) {
    return (
        <div
            style={{
                display: "grid",
                gridTemplateColumns:
                    "repeat(auto-fill,minmax(250px,1fr))",
                gap: "1rem"
            }}
        >
            {products.map(product => (
                <ProductCard
                    key={product.id}
                    product={product}
                    onAdd={onAdd}
                />
            ))}
        </div>
    );
}