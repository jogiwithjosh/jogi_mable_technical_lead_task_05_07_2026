import { Api } from "../lib/api";
import {
    OrderRequest,
    OrderResponse
} from "../types/order";

const BASE_URL = import.meta.env.VITE_API_URL ?? "http://localhost:8080";

export async function placeOrder(
    request: OrderRequest
): Promise<OrderResponse> {
    return Api.post<OrderResponse>(
        BASE_URL+"/api/order",
        request
    );
}