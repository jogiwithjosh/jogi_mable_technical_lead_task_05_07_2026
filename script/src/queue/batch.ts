import { TrackingEvent } from "../types/event";

export interface EventBatch {

    events: TrackingEvent[];

    createdAt: string;
}

export function createBatch(
    events: TrackingEvent[]
): EventBatch {

    return {

        events,

        createdAt: new Date().toISOString()
    };
}