import { TrackingEvent } from "../types/event";

export class EventQueue {

    private readonly events: TrackingEvent[] = [];

    constructor(
        private readonly maxSize: number
    ) {}

    enqueue(event: TrackingEvent): boolean {

        if (this.events.length >= this.maxSize) {
            return false;
        }

        this.events.push(event);

        return true;
    }

    dequeueBatch(size: number): TrackingEvent[] {

        return this.events.splice(0, size);
    }

    clear(): void {
        this.events.length = 0;
    }

    size(): number {
        return this.events.length;
    }

    isEmpty(): boolean {
        return this.events.length === 0;
    }

    peek(): readonly TrackingEvent[] {
        return this.events;
    }
}