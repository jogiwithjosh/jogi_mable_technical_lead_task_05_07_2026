import { Api } from "../lib/api";
import {
    OrderRequest,
    OrderResponse
} from "../types/order";

export async function placeOrder(
    request: OrderRequest
): Promise<OrderResponse> {
    return Api.post<OrderResponse>(
        "http://localhost:8080/api/order",
        request
    );
}