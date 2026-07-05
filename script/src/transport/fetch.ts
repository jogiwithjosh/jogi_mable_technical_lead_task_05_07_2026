import { EventBatch } from "../queue/batch";
import { Transport } from "./transport";

export class FetchTransport implements Transport {

    constructor(
        private readonly endpoint: string,
        private headers: Record<string,string>={}
    ) {}

    async send(batch: EventBatch): Promise<void> {

        const response = await fetch(this.endpoint, {

            method: "POST",

            keepalive: true,

            headers: {
                "Content-Type": "application/json", 
                ...this.headers
            },

            body: JSON.stringify(batch)
        });

        if (!response.ok) {
            throw new Error(
                `HTTP ${response.status}`
            );
        }
    }
}