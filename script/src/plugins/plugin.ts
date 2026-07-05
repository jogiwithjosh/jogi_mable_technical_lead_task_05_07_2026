import { Tracker } from "../tracker";

export interface Plugin {
    init(tracker: Tracker): void;
    destroy(): void;
}