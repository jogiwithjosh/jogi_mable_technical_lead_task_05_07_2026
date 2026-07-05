// -----------------------------------------------------------------------------
// Application
// -----------------------------------------------------------------------------

export const APP_NAME = "Mable Ecommerce";

export const APP_VERSION = "1.0.0";

// -----------------------------------------------------------------------------
// Environment
// -----------------------------------------------------------------------------

export const API_URL =
    import.meta.env.VITE_API_URL ??
    "http://localhost:8080";

export const USE_MOCK_API =
    import.meta.env.VITE_USE_MOCK_API === "true";

// -----------------------------------------------------------------------------
// Authentication
// -----------------------------------------------------------------------------

export const TOKEN_KEY = "auth-token";

export const USER_KEY = "current-user";

// -----------------------------------------------------------------------------
// Shopping Cart
// -----------------------------------------------------------------------------

export const CART_KEY = "shopping-cart";

// -----------------------------------------------------------------------------
// Analytics
// -----------------------------------------------------------------------------

export const ANALYTICS_ENDPOINT = "/events";

export const ANALYTICS_FLUSH_INTERVAL = 5000;

export const ANALYTICS_BATCH_SIZE = 20;

export const ANALYTICS_MAX_RETRIES = 3;

export const ANALYTICS_DEBUG =
    import.meta.env.DEV;

// -----------------------------------------------------------------------------
// Session
// -----------------------------------------------------------------------------

export const SESSION_TIMEOUT_MS =
    30 * 60 * 1000; // 30 minutes