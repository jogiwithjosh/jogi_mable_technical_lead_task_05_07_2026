import { EventBatch } from '../queue/batch';
export interface Transport {
    send(batch: EventBatch): Promise<void>;
}
//# sourceMappingURL=transport.d.ts.map