import { TrackerConfig } from "../types/config";
import { DEFAULT_CONFIG } from "./defaults";
import { validate } from "./validator";

export function createConfig(
    config: TrackerConfig
): Required<TrackerConfig> {

    validate(config);

    return {

        ...DEFAULT_CONFIG,

        ...config
    };
}