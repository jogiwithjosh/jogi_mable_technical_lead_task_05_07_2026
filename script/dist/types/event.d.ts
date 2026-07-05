export interface TrackingEvent {
    id: string;
    type: string;
    timestamp: string;
    page: string;
    referrer: string;
    url: string;
    sessionId?: string;
    userId?: string;
    anonymousId?: string;
    properties: Record<string, unknown>;
}
//# sourceMappingURL=event.d.ts.map