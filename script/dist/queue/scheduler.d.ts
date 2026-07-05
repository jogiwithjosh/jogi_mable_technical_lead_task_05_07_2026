export declare class FlushScheduler {
    private readonly interval;
    private readonly callback;
    private timer?;
    constructor(interval: number, callback: () => Promise<void>);
    start(): void;
    stop(): void;
}
//# sourceMappingURL=scheduler.d.ts.map