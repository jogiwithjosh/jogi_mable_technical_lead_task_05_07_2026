import { tracker } from "./analytics.client";

export const Analytics = {

    identify(
        id: string,
        email: string
    ) {

        tracker.identify({
            id,
            email
        });

    },

    logout() {

        tracker.reset();

    },

    page() {

        tracker.page();

    },

    addToCart(
        productId: string,
        quantity: number,
        price: number
    ) {

        tracker.track("AddToCart", {
    productId: productId,
    quantity: quantity,
    price: price
});

    },

    checkout(
        cartId: string,
        total: number
    ) {

        tracker.track( "Checkout", {
            cartId,
            total
        });

    },

    paymentInfoAdded(
        paymentMethod: string
    ) {

        tracker.track("paymentInfoAdded", {
            paymentMethod
        });

    },

    purchase(
        orderId: string,
        total: number
    ) {

        tracker.track("purchase",{
            orderId,
            total
    });

    },

    event(
        name: string,
        properties: Record<string, unknown>
    ) {

        tracker.track(
            name,
            properties
        );

    }

};