import { TrackingEvent } from '../types/event';
export interface EventBatch {
    events: TrackingEvent[];
    createdAt: string;
}
export declare function createBatch(events: TrackingEvent[]): EventBatch;
//# sourceMappingURL=batch.d.ts.map