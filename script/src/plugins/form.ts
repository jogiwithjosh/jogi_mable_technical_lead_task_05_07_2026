import { Tracker } from "../tracker";
import { Plugin } from "./plugin";

export class FormPlugin implements Plugin {

    private tracker!: Tracker;

    private readonly listener = (event: SubmitEvent) => {

        const form = event.target as HTMLFormElement | null;

        if (!form) {
            return;
        }

        const email =
            form.querySelector<HTMLInputElement>(
                "input[type=email]"
            );

        this.tracker.track("Lead", {

            formId: form.id,

            action: form.action,

            method: form.method,

            hasEmail: !!email
        });
    };

    init(tracker: Tracker): void {

        this.tracker = tracker;

        document.addEventListener("submit", this.listener);
    }

    destroy(): void {

        document.removeEventListener("submit", this.listener);
    }

}