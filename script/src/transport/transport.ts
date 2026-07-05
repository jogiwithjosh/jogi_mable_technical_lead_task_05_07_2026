import { EventBatch } from "../queue/batch";

export interface Transport {

    send(batch: EventBatch): Promise<void>;
}