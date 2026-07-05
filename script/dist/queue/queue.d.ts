import { TrackingEvent } from '../types/event';
export declare class EventQueue {
    private readonly maxSize;
    private readonly events;
    constructor(maxSize: number);
    enqueue(event: TrackingEvent): boolean;
    dequeueBatch(size: number): TrackingEvent[];
    clear(): void;
    size(): number;
    isEmpty(): boolean;
    peek(): readonly TrackingEvent[];
}
//# sourceMappingURL=queue.d.ts.map