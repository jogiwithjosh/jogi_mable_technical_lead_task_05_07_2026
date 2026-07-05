import { Plugin } from '../plugins/plugin';
import { TrackerConfig } from '../types/config';
import { ITracker } from '../types/tracker';
import { User } from '../types/user';
export declare class Tracker implements ITracker {
    private state;
    private config;
    private logger;
    private user?;
    private queue;
    private scheduler;
    private transport;
    private pending;
    private network;
    private readonly plugins;
    private session;
    private userManager;
    register(plugin: Plugin): void;
    init(config: TrackerConfig): void;
    identify(user: User): void;
    page(properties?: Record<string, unknown>): void;
    track(type: string, properties?: Record<string, unknown>): void;
    flush(): Promise<void>;
    reset(): void;
    destroy(): void;
    private ensureInitialized;
    private flushPending;
    private enrichEvent;
}
//# sourceMappingURL=tracker.d.ts.map