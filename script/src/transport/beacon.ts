import { EventBatch } from "../queue/batch";
import { Transport } from "./transport";

export class BeaconTransport implements Transport {

    constructor(
        private readonly endpoint: string
    ) {}

    async send(batch: EventBatch): Promise<void> {

    const blob = new Blob(
        [JSON.stringify(batch)],
        {
            type: "application/json"
        }
    );

    const ok = navigator.sendBeacon(
        this.endpoint,
        blob
    );

    if (!ok) {
        throw new Error("sendBeacon failed");
    }
}
}