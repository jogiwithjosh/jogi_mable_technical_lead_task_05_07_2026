var l = Object.defineProperty;
var g = (s, e, t) => e in s ? l(s, e, { enumerable: !0, configurable: !0, writable: !0, value: t }) : s[e] = t;
var i = (s, e, t) => g(s, typeof e != "symbol" ? e + "" : e, t);
function c(s) {
  return {
    events: s,
    createdAt: (/* @__PURE__ */ new Date()).toISOString()
  };
}
class p {
  constructor(e) {
    i(this, "events", []);
    this.maxSize = e;
  }
  enqueue(e) {
    return this.events.length >= this.maxSize ? !1 : (this.events.push(e), !0);
  }
  dequeueBatch(e) {
    return this.events.splice(0, e);
  }
  clear() {
    this.events.length = 0;
  }
  size() {
    return this.events.length;
  }
  isEmpty() {
    return this.events.length === 0;
  }
  peek() {
    return this.events;
  }
}
class f {
  constructor(e, t) {
    i(this, "timer");
    this.interval = e, this.callback = t;
  }
  start() {
    this.timer || (this.timer = window.setInterval(async () => {
      await this.callback();
    }, this.interval));
  }
  stop() {
    this.timer && (clearInterval(this.timer), this.timer = void 0);
  }
}
function a() {
  return crypto.randomUUID();
}
const d = "tracking-sdk-session", v = 1800 * 1e3;
class w {
  constructor() {
    i(this, "session");
    this.loadOrCreate();
  }
  current() {
    return this.touch(), this.session;
  }
  incrementEvent() {
    this.session.eventCount++, this.touch();
  }
  incrementPageView() {
    this.session.pageViews++, this.touch();
  }
  duration() {
    return Date.now() - this.session.startedAt;
  }
  touch() {
    this.session.lastActivity = Date.now(), this.save();
  }
  loadOrCreate() {
    const e = localStorage.getItem(d);
    if (!e) {
      this.create();
      return;
    }
    const t = JSON.parse(e);
    if (Date.now() - t.lastActivity > v) {
      this.create();
      return;
    }
    this.session = t;
  }
  create() {
    const e = localStorage.getItem("tracking-anonymous-id") ?? a();
    localStorage.setItem(
      "tracking-anonymous-id",
      e
    ), this.session = {
      id: a(),
      anonymousId: e,
      startedAt: Date.now(),
      lastActivity: Date.now(),
      pageViews: 0,
      eventCount: 0
    }, this.save();
  }
  save() {
    localStorage.setItem(
      d,
      JSON.stringify(this.session)
    );
  }
}
class m {
  constructor(e) {
    this.endpoint = e;
  }
  async send(e) {
    const t = new Blob(
      [JSON.stringify(e)],
      {
        type: "application/json"
      }
    );
    if (!navigator.sendBeacon(
      this.endpoint,
      t
    ))
      throw new Error("sendBeacon failed");
  }
}
class y {
  constructor(e, t = {}) {
    this.endpoint = e, this.headers = t;
  }
  async send(e) {
    const t = await fetch(this.endpoint, {
      method: "POST",
      keepalive: !0,
      headers: {
        "Content-Type": "application/json",
        ...this.headers
      },
      body: JSON.stringify(e)
    });
    if (!t.ok)
      throw new Error(
        `HTTP ${t.status}`
      );
  }
}
class E {
  constructor(e, t = 3, n = 1e3) {
    this.transport = e, this.maxRetries = t, this.retryDelay = n;
  }
  async send(e) {
    let t;
    for (let n = 1; n <= this.maxRetries; n++)
      try {
        await this.transport.send(e);
        return;
      } catch (o) {
        if (t = o, n === this.maxRetries)
          break;
        await this.sleep(this.retryDelay * n);
      }
    throw t;
  }
  sleep(e) {
    return new Promise((t) => setTimeout(t, e));
  }
}
function I(s) {
  const e = typeof navigator.sendBeacon == "function" ? new m(s) : new y(s);
  return new E(e);
}
class b {
  constructor() {
    i(this, "online", navigator.onLine);
    window.addEventListener("online", () => {
      this.online = !0;
    }), window.addEventListener("offline", () => {
      this.online = !1;
    });
  }
  isOnline() {
    return this.online;
  }
}
class k {
  constructor() {
    i(this, "user");
  }
  identify(e) {
    this.user = e;
  }
  current() {
    return this.user;
  }
  reset() {
    this.user = void 0;
  }
  isAuthenticated() {
    return this.user !== void 0;
  }
}
function S(s) {
  window.addEventListener("beforeunload", s), document.addEventListener("visibilitychange", () => {
    document.visibilityState === "hidden" && s();
  });
}
const D = {
  endpoint: "",
  batchSize: 20,
  flushInterval: 5e3,
  maxRetries: 3,
  retryDelay: 1e3,
  headers: {},
  debug: !1,
  plugins: !0
};
function T(s) {
  if (!s.endpoint)
    throw new Error("endpoint is required");
  if (s.batchSize !== void 0 && s.batchSize <= 0)
    throw new Error("batchSize must be > 0");
  if (s.flushInterval !== void 0 && s.flushInterval <= 0)
    throw new Error("flushInterval must be > 0");
}
function L(s) {
  return T(s), {
    ...D,
    ...s
  };
}
function C() {
  return (/* @__PURE__ */ new Date()).toISOString();
}
class O {
  constructor(e) {
    this.enabled = e;
  }
  debug(...e) {
    this.enabled && console.log("[TrackingSDK]", ...e);
  }
  error(...e) {
    console.error("[TrackingSDK]", ...e);
  }
}
var r = /* @__PURE__ */ ((s) => (s[s.CREATED = 0] = "CREATED", s[s.INITIALIZED = 1] = "INITIALIZED", s[s.DESTROYED = 2] = "DESTROYED", s))(r || {});
class z {
  constructor() {
    i(this, "state", r.CREATED);
    i(this, "config");
    i(this, "logger");
    i(this, "user");
    i(this, "queue");
    i(this, "scheduler");
    i(this, "transport");
    i(this, "pending", []);
    i(this, "network");
    i(this, "plugins", []);
    i(this, "session", new w());
    i(this, "userManager", new k());
  }
  register(e) {
    this.plugins.push(e);
  }
  init(e) {
    this.config = L(e), this.logger = new O(this.config.debug), this.state = r.INITIALIZED, this.queue = new p(1e3), this.transport = I(
      this.config.endpoint
    ), this.scheduler = new f(
      this.config.flushInterval,
      () => this.flush()
    ), this.scheduler.start(), S(() => {
      this.flush(), this.flushPending();
    }), this.network = new b(), window.addEventListener("online", () => {
      this.flushPending();
    });
    for (const t of this.plugins)
      t.init(this);
    this.logger.debug("Tracker initialized");
  }
  identify(e) {
    this.userManager.identify(e), this.user = e, this.logger.debug("User identified", e);
  }
  page(e = {}) {
    this.track("PageView", e);
  }
  track(e, t = {}) {
    var h, u;
    this.ensureInitialized();
    const n = {
      id: a(),
      type: e,
      timestamp: C(),
      page: document.title,
      url: window.location.href,
      referrer: document.referrer,
      userId: (h = this.user) == null ? void 0 : h.id,
      anonymousId: (u = this.user) == null ? void 0 : u.anonymousId,
      properties: t
    }, o = this.enrichEvent(n);
    this.logger.debug(n), this.logger.debug(o), this.queue.enqueue(o) || this.logger.error(
      "Queue full. Dropping event."
    ), this.queue.size() >= this.config.batchSize && this.flush();
  }
  async flush() {
    if (!this.network.isOnline()) {
      const n = this.queue.dequeueBatch(
        this.config.batchSize
      );
      n.length > 0 && this.pending.push(c(n));
      return;
    }
    if (this.queue.isEmpty())
      return;
    const e = this.queue.dequeueBatch(
      this.config.batchSize
    ), t = c(e);
    try {
      await this.transport.send(t), this.logger.debug(
        `Sent ${e.length} events`
      );
    } catch (n) {
      this.logger.error(n), this.pending.push(t);
    }
  }
  reset() {
    this.user = void 0;
  }
  destroy() {
    this.scheduler.stop(), this.flush();
    for (const e of this.plugins)
      e.destroy();
    this.state = r.DESTROYED;
  }
  ensureInitialized() {
    if (this.state !== r.INITIALIZED)
      throw new Error("Tracker not initialized");
  }
  async flushPending() {
    if (this.pending.length === 0)
      return;
    const e = [];
    for (const t of this.pending)
      try {
        await this.transport.send(t);
      } catch {
        e.push(t);
      }
    this.pending = e;
  }
  enrichEvent(e) {
    var n;
    const t = this.session.current();
    return {
      ...e,
      sessionId: t.id,
      anonymousId: t.anonymousId,
      userId: (n = this.user) == null ? void 0 : n.id,
      properties: {
        ...e.properties,
        sessionDuration: this.session.duration(),
        eventNumber: t.eventCount,
        pageViews: t.pageViews,
        userAgent: navigator.userAgent,
        language: navigator.language,
        timezone: Intl.DateTimeFormat().resolvedOptions().timeZone
      }
    };
  }
}
class N {
  constructor() {
    i(this, "tracker");
    i(this, "onPageChange", () => {
      this.tracker.page();
    });
  }
  init(e) {
    this.tracker = e, e.page(), window.addEventListener("popstate", this.onPageChange), window.addEventListener("hashchange", this.onPageChange);
  }
  destroy() {
    window.removeEventListener("popstate", this.onPageChange), window.removeEventListener("hashchange", this.onPageChange);
  }
}
class q {
  constructor() {
    i(this, "listener", (e) => {
      var n;
      const t = e.target;
      t && this.tracker.track("Click", {
        tag: t.tagName,
        id: t.id,
        className: t.className,
        text: (n = t.textContent) == null ? void 0 : n.trim().substring(0, 100),
        x: e.clientX,
        y: e.clientY
      });
    });
    i(this, "tracker");
  }
  init(e) {
    this.tracker = e, document.addEventListener("click", this.listener);
  }
  destroy() {
    document.removeEventListener("click", this.listener);
  }
}
class R {
  constructor() {
    i(this, "tracker");
    i(this, "listener", (e) => {
      const t = e.target;
      if (!t)
        return;
      const n = t.querySelector(
        "input[type=email]"
      );
      this.tracker.track("Lead", {
        formId: t.id,
        action: t.action,
        method: t.method,
        hasEmail: !!n
      });
    });
  }
  init(e) {
    this.tracker = e, document.addEventListener("submit", this.listener);
  }
  destroy() {
    document.removeEventListener("submit", this.listener);
  }
}
export {
  q as ClickPlugin,
  R as FormPlugin,
  N as PagePlugin,
  z as Tracker
};
//# sourceMappingURL=index.js.map
