import { EventBatch } from '../queue/batch';
import { Transport } from './transport';
export declare class RetryTransport implements Transport {
    private readonly transport;
    private readonly maxRetries;
    private readonly retryDelay;
    constructor(transport: Transport, maxRetries?: number, retryDelay?: number);
    send(batch: EventBatch): Promise<void>;
    private sleep;
}
//# sourceMappingURL=retry.d.ts.map