export class Logger {

    constructor(
        private enabled: boolean
    ) {}

    debug(...args: unknown[]) {

        if (!this.enabled) {
            return;
        }

        console.log("[TrackingSDK]", ...args);
    }

    error(...args: unknown[]) {

        console.error("[TrackingSDK]", ...args);
    }
}