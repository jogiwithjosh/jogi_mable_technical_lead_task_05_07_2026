export function registerUnload(
    callback: () => void
): void {

    window.addEventListener("beforeunload", callback);

    document.addEventListener("visibilitychange", () => {

        if (document.visibilityState === "hidden") {

            callback();

        }

    });

}