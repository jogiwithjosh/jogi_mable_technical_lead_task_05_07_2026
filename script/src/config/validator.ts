import { TrackerConfig } from "../types/config";

export function validate(config: TrackerConfig): void {

    if (!config.endpoint) {
        throw new Error("endpoint is required");
    }

    if (config.batchSize !== undefined && config.batchSize <= 0) {
        throw new Error("batchSize must be > 0");
    }

    if (config.flushInterval !== undefined && config.flushInterval <= 0) {
        throw new Error("flushInterval must be > 0");
    }
}