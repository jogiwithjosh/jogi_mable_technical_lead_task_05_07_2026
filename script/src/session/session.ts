export interface Session {

    id: string;

    anonymousId: string;

    startedAt: number;

    lastActivity: number;

    pageViews: number;

    eventCount: number;
}