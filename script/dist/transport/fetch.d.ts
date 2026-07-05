import { EventBatch } from '../queue/batch';
import { Transport } from './transport';
export declare class FetchTransport implements Transport {
    private readonly endpoint;
    private headers;
    constructor(endpoint: string, headers?: Record<string, string>);
    send(batch: EventBatch): Promise<void>;
}
//# sourceMappingURL=fetch.d.ts.map