import { Tracker } from "../tracker";
import { Plugin } from "./plugin";

export class ClickPlugin implements Plugin {

    private readonly listener = (event: MouseEvent) => {

        const target = event.target as HTMLElement | null;

        if (!target) {
            return;
        }

        this.tracker.track("Click", {
            tag: target.tagName,
            id: target.id,
            className: target.className,
            text: target.textContent?.trim().substring(0, 100),
            x: event.clientX,
            y: event.clientY
        });
    };

    private tracker!: Tracker;

    init(tracker: Tracker): void {

        this.tracker = tracker;

        document.addEventListener("click", this.listener);
    }

    destroy(): void {

        document.removeEventListener("click", this.listener);
    }

}