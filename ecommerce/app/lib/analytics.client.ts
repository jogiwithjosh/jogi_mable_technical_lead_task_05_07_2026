import {
    ClickPlugin,
    FormPlugin,
    PagePlugin,
    Tracker
} from "@mable/tracking-sdk";

const BASE_URL = import.meta.env.VITE_API_URL ?? "http://localhost:8080";

export const tracker = new Tracker();

tracker.register(new PagePlugin());

tracker.register(new ClickPlugin());

tracker.register(new FormPlugin());

tracker.init({

    endpoint: BASE_URL + "/api/events",

    batchSize: 20,

    flushInterval: 5000,

    debug: true

});