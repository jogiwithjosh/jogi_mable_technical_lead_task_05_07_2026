export class FlushScheduler {

    private timer?: number;

    constructor(
        private readonly interval: number,
        private readonly callback: () => Promise<void>
    ) {}

    start(): void {

        if (this.timer) {
            return;
        }

        this.timer = window.setInterval(async () => {

            await this.callback();

        }, this.interval);
    }

    stop(): void {

        if (!this.timer) {
            return;
        }

        clearInterval(this.timer);

        this.timer = undefined;
    }
}