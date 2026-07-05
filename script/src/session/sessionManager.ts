import { uuid } from "../utils/uuid";
import { Session } from "./session";

const STORAGE_KEY = "tracking-sdk-session";

const SESSION_TIMEOUT = 30 * 60 * 1000;

export class SessionManager {

    private session!: Session;

    constructor() {

        this.loadOrCreate();
    }

    current(): Session {

        this.touch();

        return this.session;
    }

    incrementEvent(): void {

        this.session.eventCount++;

        this.touch();
    }

    incrementPageView(): void {

        this.session.pageViews++;

        this.touch();
    }

    duration(): number {

        return Date.now() - this.session.startedAt;
    }

    private touch(): void {

        this.session.lastActivity = Date.now();

        this.save();
    }

    private loadOrCreate(): void {

        const stored = localStorage.getItem(STORAGE_KEY);

        if (!stored) {

            this.create();

            return;
        }

        const session = JSON.parse(stored) as Session;

        if (Date.now() - session.lastActivity > SESSION_TIMEOUT) {

            this.create();

            return;
        }

        this.session = session;
    }

    private create(): void {

        const anonymousId =
            localStorage.getItem("tracking-anonymous-id")
            ?? uuid();

        localStorage.setItem(
            "tracking-anonymous-id",
            anonymousId
        );

        this.session = {

            id: uuid(),

            anonymousId,

            startedAt: Date.now(),

            lastActivity: Date.now(),

            pageViews: 0,

            eventCount: 0
        };

        this.save();
    }

    private save(): void {

        localStorage.setItem(
            STORAGE_KEY,
            JSON.stringify(this.session)
        );
    }

}