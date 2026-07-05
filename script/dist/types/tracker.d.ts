import { TrackerConfig } from './config';
import { User } from './user';
export interface ITracker {
    init(config: TrackerConfig): void;
    identify(user: User): void;
    page(properties?: Record<string, unknown>): void;
    track(type: string, properties?: Record<string, unknown>): void;
    flush(): Promise<void>;
    reset(): void;
    destroy(): void;
}
//# sourceMappingURL=tracker.d.ts.map