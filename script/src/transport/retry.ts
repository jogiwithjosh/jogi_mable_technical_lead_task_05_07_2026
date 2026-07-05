import { EventBatch } from "../queue/batch";
import { Transport } from "./transport";

export class RetryTransport implements Transport {

    constructor(
        private readonly transport: Transport,
        private readonly maxRetries = 3,
        private readonly retryDelay = 1000
    ) {}

    async send(batch: EventBatch): Promise<void> {

        let lastError: unknown;

        for (let attempt = 1; attempt <= this.maxRetries; attempt++) {

            try {

                await this.transport.send(batch);

                return;

            } catch (err) {

                lastError = err;

                if (attempt === this.maxRetries) {
                    break;
                }

                await this.sleep(this.retryDelay * attempt);
            }
        }

        throw lastError;
    }

    private sleep(ms: number): Promise<void> {

        return new Promise(resolve => setTimeout(resolve, ms));
    }

}