import { Product } from "./products";

export interface CartItem {
    product: Product;
    quantity: number;
}

export interface Cart {
    items: CartItem[];
    totalItems: number;
    totalAmount: number;
}