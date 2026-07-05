export interface TrackerConfig {
    endpoint: string;
    batchSize?: number;
    flushInterval?: number;
    maxRetries?: number;
    retryDelay?: number;
    headers?: Record<string, string>;
    debug?: boolean;
    plugins?: boolean;
}
//# sourceMappingURL=config.d.ts.map