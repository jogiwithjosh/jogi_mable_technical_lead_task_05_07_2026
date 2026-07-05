import { EventBatch } from '../queue/batch';
import { Transport } from './transport';
export declare class BeaconTransport implements Transport {
    private readonly endpoint;
    constructor(endpoint: string);
    send(batch: EventBatch): Promise<void>;
}
//# sourceMappingURL=beacon.d.ts.map