import { Session } from './session';
export declare class SessionManager {
    private session;
    constructor();
    current(): Session;
    incrementEvent(): void;
    incrementPageView(): void;
    duration(): number;
    private touch;
    private loadOrCreate;
    private create;
    private save;
}
//# sourceMappingURL=sessionManager.d.ts.map