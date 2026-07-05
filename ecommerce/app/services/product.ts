import { Api } from "../lib/api";
import { Product } from "../types/products";

export async function getProducts(): Promise<Product[]> {
    return Api.get<Product[]>("https://fakestoreapi.com/products");
}

export async function getProduct(
    id: string
): Promise<Product> {
    return Api.get<Product>(
        `https://fakestoreapi.com/products/${id}`
    );
}