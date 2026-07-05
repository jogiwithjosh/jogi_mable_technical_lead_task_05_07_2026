import { Cart, CartItem } from "../types/cart";
import { Product } from "../types/products";

const CART_KEY = "shopping-cart";

function calculate(items: CartItem[]): Cart {

    const totalItems = items.reduce(
        (sum, item) => sum + item.quantity,
        0
    );

    const totalAmount = items.reduce(
        (sum, item) =>
            sum + item.product.price * item.quantity,
        0
    );

    return {
        items,
        totalItems,
        totalAmount
    };
}

export function getCart(): Cart {

    const value = localStorage.getItem(CART_KEY);

    if (!value) {
        return calculate([]);
    }

    return calculate(JSON.parse(value));
}

export function saveCart(items: CartItem[]) {
    localStorage.setItem(
        CART_KEY,
        JSON.stringify(items)
    );
}

export function addToCart(product: Product) {

    const cart = getCart();

    const existing = cart.items.find(
        i => i.product.id === product.id
    );

    if (existing) {

        existing.quantity++;

    } else {

        cart.items.push({
            product,
            quantity: 1
        });

    }

    saveCart(cart.items);
}

export function removeFromCart(productId: string) {

    const cart = getCart();

    saveCart(
        cart.items.filter(
            item => item.product.id !== productId
        )
    );
}

export function updateQuantity(
    productId: string,
    quantity: number
) {

    const cart = getCart();

    const item = cart.items.find(
        i => i.product.id === productId
    );

    if (!item) {
        return;
    }

    item.quantity = quantity;

    saveCart(cart.items);
}

export function clearCart() {

    saveCart([]);

}