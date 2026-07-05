import { Tracker } from "../tracker";
import { Plugin } from "./plugin";

export class PagePlugin implements Plugin {

    private tracker!: Tracker;

    private readonly onPageChange = () => {
        this.tracker.page();
    };

    init(tracker: Tracker): void {

        this.tracker = tracker;

        tracker.page();

        window.addEventListener("popstate", this.onPageChange);

        window.addEventListener("hashchange", this.onPageChange);
    }

    destroy(): void {

        window.removeEventListener("popstate", this.onPageChange);

        window.removeEventListener("hashchange", this.onPageChange);
    }

}