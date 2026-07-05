export class NetworkMonitor {

    private online = navigator.onLine;

    constructor() {

        window.addEventListener("online", () => {

            this.online = true;

        });

        window.addEventListener("offline", () => {

            this.online = false;

        });

    }

    isOnline(): boolean {

        return this.online;
    }

}