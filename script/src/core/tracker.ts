import { Plugin } from "../plugins/plugin";
import { createBatch, EventBatch } from "../queue/batch";
import { EventQueue } from "../queue/queue";
import { FlushScheduler } from "../queue/scheduler";
import { SessionManager } from "../session/sessionManager";
import { createTransport } from "../transport";
import { NetworkMonitor } from "../transport/network";
import { Transport } from "../transport/transport";
import { TrackerConfig } from "../types/config";
import { TrackingEvent } from "../types/event";
import { ITracker } from "../types/tracker";
import { User } from "../types/user";
import { UserManager } from "../user/userManager";
import { registerUnload } from "./browser";

import { createConfig } from "../config/config";

import { now } from "../utils/clock";
import { Logger } from "../utils/logger";
import { uuid } from "../utils/uuid";

import { TrackerState } from "./lifecycle";

export class Tracker implements ITracker {

    private state = TrackerState.CREATED;

    private config!: Required<TrackerConfig>;

    private logger!: Logger;

    private user?: User;

    private queue!: EventQueue;

    private scheduler!: FlushScheduler;

    private transport!: Transport;

    private pending: EventBatch[] = [];

    private network!: NetworkMonitor;

    private readonly plugins: Plugin[] = [];

    private session = new SessionManager();

    private userManager = new UserManager();

    register(plugin: Plugin): void {
        this.plugins.push(plugin);
    }

    init(config: TrackerConfig): void {

        this.config = createConfig(config);

        this.logger = new Logger(this.config.debug);

        this.state = TrackerState.INITIALIZED;

        this.queue = new EventQueue(1000);

        this.transport = createTransport(
            this.config.endpoint
        );

        this.scheduler = new FlushScheduler(
            this.config.flushInterval,
            () => this.flush()
        );

        this.scheduler.start();
        registerUnload(() => {
            void this.flush();
            void this.flushPending();
        });

        this.network = new NetworkMonitor();

        window.addEventListener("online", () => {
            void this.flushPending();
        });

        for (const plugin of this.plugins) {
            plugin.init(this);
        }

        this.logger.debug("Tracker initialized");
    }

    identify(user: User): void {

        this.userManager.identify(user);
        this.user = user;

        this.logger.debug("User identified", user);
    }

    page(properties: Record<string, unknown> = {}): void {

        this.track("PageView", properties);
    }

    track(
        type: string,
        properties: Record<string, unknown> = {}
    ): void {

        this.ensureInitialized();

        const event: TrackingEvent = {

            id: uuid(),

            type,

            timestamp: now(),

            page: document.title,

            url: window.location.href,

            referrer: document.referrer,

            userId: this.user?.id,

            anonymousId: this.user?.anonymousId,

            properties
        };

        const eventEnriced = this.enrichEvent(event);

        this.logger.debug(event);
        this.logger.debug(eventEnriced);

        // Queue
        const accepted = this.queue.enqueue(eventEnriced);

        if (!accepted) {

            this.logger.error(
                "Queue full. Dropping event."
            );
        }

        if (
            this.queue.size() >= this.config.batchSize
        ) {

            void this.flush();
        }
    }

    async flush(): Promise<void> {

        if (!this.network.isOnline()) {

            const events = this.queue.dequeueBatch(
                this.config.batchSize
            );

            if (events.length > 0) {
                this.pending.push(createBatch(events));
            }

            return;
        }

        if (this.queue.isEmpty()) {
            return;
        }

        const events = this.queue.dequeueBatch(
            this.config.batchSize
        );

        const batch = createBatch(events);

        try {

            await this.transport.send(batch);

            this.logger.debug(
                `Sent ${events.length} events`
            );

        } catch (err) {

            this.logger.error(err);
            this.pending.push(batch);
            // Retry queue will be implemented next.
        }
    }

    reset(): void {

        this.user = undefined;
    }

    destroy(): void {
        this.scheduler.stop();

        void this.flush();
        for (const plugin of this.plugins) {
            plugin.destroy();
        }

        this.state = TrackerState.DESTROYED;
    }

    private ensureInitialized() {

        if (this.state !== TrackerState.INITIALIZED) {
            throw new Error("Tracker not initialized");
        }
    }

    private async flushPending(): Promise<void> {

        if (this.pending.length === 0) {
            return;
        }

        const failed: EventBatch[] = [];

        for (const batch of this.pending) {

            try {

                await this.transport.send(batch);

            } catch {

                failed.push(batch);

            }

        }

        this.pending = failed;
    }

    private enrichEvent(event: TrackingEvent): TrackingEvent {
    const session = this.session.current();

    return {
        ...event,
        sessionId: session.id,
        anonymousId: session.anonymousId,
        userId: this.user?.id,
        properties: {
            ...event.properties,
            sessionDuration: this.session.duration(),
            eventNumber: session.eventCount,
            pageViews: session.pageViews,
            userAgent: navigator.userAgent,
            language: navigator.language,
            timezone: Intl.DateTimeFormat().resolvedOptions().timeZone
        }
    };
}
}