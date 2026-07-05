import { User } from "../types/auth";
import { TOKEN_KEY } from "./constants";

const USER_KEY = "current-user";
const CART_KEY = "shopping-cart";

const isBrowser = typeof window !== "undefined";

function getLocalStorage(): Storage | null {
    if (!isBrowser) {
        return null;
    }

    return window.localStorage;
}

export const Storage = {

    // -----------------------
    // Token
    // -----------------------

    getToken(): string | null {
        return getLocalStorage()?.getItem(TOKEN_KEY) ?? null;
    },

    setToken(token: string): void {
        getLocalStorage()?.setItem(TOKEN_KEY, token);
    },

    removeToken(): void {
        getLocalStorage()?.removeItem(TOKEN_KEY);
    },

    // -----------------------
    // User
    // -----------------------

    setUser(user: User): void {
        getLocalStorage()?.setItem(
            USER_KEY,
            JSON.stringify(user)
        );
    },

    getUser(): User | null {

        const value =
            getLocalStorage()?.getItem(USER_KEY);

        if (!value) {
            return null;
        }

        try {
            return JSON.parse(value);
        } catch {
            return null;
        }
    },

    removeUser(): void {
        getLocalStorage()?.removeItem(USER_KEY);
    },

    // -----------------------
    // Cart
    // -----------------------

    getCart(): string | null {
        return getLocalStorage()?.getItem(CART_KEY) ?? null;
    },

    setCart(value: string): void {
        getLocalStorage()?.setItem(CART_KEY, value);
    },

    clearCart(): void {
        getLocalStorage()?.removeItem(CART_KEY);
    },

    // -----------------------
    // Generic helpers
    // -----------------------

    clear(): void {
        getLocalStorage()?.clear();
    }

};