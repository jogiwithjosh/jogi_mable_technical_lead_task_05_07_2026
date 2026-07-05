import { CartItem } from "./cart";

export interface OrderRequest {
    items: CartItem[];
    paymentMethod: string;
}

export interface OrderResponse {
    orderId: string;
    total: number;
    status: string;
}