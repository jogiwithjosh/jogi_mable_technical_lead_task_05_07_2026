import {
    ClickPlugin,
    FormPlugin,
    PagePlugin,
    Tracker
} from "@mable/tracking-sdk";

export const tracker = new Tracker();

tracker.register(new PagePlugin());

tracker.register(new ClickPlugin());

tracker.register(new FormPlugin());

tracker.init({

    endpoint: "http://localhost:8080/api/events",

    batchSize: 20,

    flushInterval: 5000,

    debug: true

});