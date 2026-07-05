import { BeaconTransport } from "./beacon";
import { FetchTransport } from "./fetch";
import { RetryTransport } from "./retry";
import { Transport } from "./transport";

export function createTransport(endpoint: string): Transport {

    const base =
        typeof navigator.sendBeacon === "function"
            ? new BeaconTransport(endpoint)
            : new FetchTransport(endpoint);

    return new RetryTransport(base);
}