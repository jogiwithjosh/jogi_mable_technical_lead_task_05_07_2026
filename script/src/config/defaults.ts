import { TrackerConfig } from "../types/config";

export const DEFAULT_CONFIG: Required<TrackerConfig> = {
    endpoint: "",
    batchSize: 20,
    flushInterval: 5000,
    maxRetries: 3,
    retryDelay: 1000,
    headers: {},
    debug: false,
    plugins: true
};