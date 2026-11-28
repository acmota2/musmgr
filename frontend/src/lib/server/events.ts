import { error } from "@sveltejs/kit";
import { getApiUrl } from "./api";
import { handleFetchError } from "./utils";

interface GetEventResponse {
  id: string;
  name: string;
  happened_at: string;
  description?: string;
  event_type: string;
  created_at: Date;
  updated_at: Date;
}

export interface CreateEventPayload {
  event_type: string;
  happened_at: string;
  name: string;
  description?: string;
}

export interface MusmgrEvent {
  id: string;
  name: string;
  happenedAt: string;
  description?: string;
  eventType: string;
}

export type MusmgrEventTimetable = Record<string, Record<string, MusmgrEvent[]>>;

export function eventListToTimetable(events: MusmgrEvent[]): MusmgrEventTimetable {
  const splitHappenedAt = (event: MusmgrEvent) => event.happenedAt.split("-");
  return events.reduce<MusmgrEventTimetable>((acc, e) => {
    const [year] = splitHappenedAt(e);
    if (!acc[year]) {
      acc[year] = {};
    }
    const month = new Date(e.happenedAt).toLocaleDateString("en-US", { month: "long" });
    if (!acc[year][month]) {
      acc[year][month] = [];
    }
    acc[year][month].push(e);
    return acc;
  }, {});
}

export function timetableToEventList(events: MusmgrEventTimetable): MusmgrEvent[] {
  const res: MusmgrEvent[] = [];
  for (const months of Object.values(events)) {
    for (const events of Object.values(months)) {
      res.push(...events);
    }
  }
  return res;
}

export async function getEventTypes(fetch: typeof globalThis.fetch): Promise<string[]> {
  const url = `${getApiUrl()}/event_types`;
  const res = await fetch(url);

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  return await res.json();
}

export async function createEvent(
  fetch: typeof globalThis.fetch,
  data: CreateEventPayload,
): Promise<string> {
  const url = `${getApiUrl()}/events`;
  const res = await fetch(url, {
    method: "POST",
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  const location = res.headers.get("Location");
  if (location === null) {
    throw error(500, {
      code: "INVALID_LOCATION_HEADER",
      message: "It's not you, it's us",
      status: 500,
    });
  }

  return location;
}

export async function getEvent(fetch: typeof globalThis.fetch, id: string): Promise<MusmgrEvent> {
  const url = `${getApiUrl()}/events/${id}`;
  const res = await fetch(url);

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  const event: GetEventResponse = await res.json();

  return {
    id: event.id,
    description: event.description,
    eventType: event.event_type,
    happenedAt: event.happened_at,
    name: event.name,
  };
}

async function getEventsWithUrl(
  fetch: typeof globalThis.fetch,
  url: string,
): Promise<MusmgrEvent[]> {
  const res = await fetch(url);

  const eventsData: GetEventResponse[] = (await res.json()) || [];

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  return eventsData.map((event) => ({
    id: event.id,
    name: event.name,
    happenedAt: event.happened_at,
    description: event.description,
    eventType: event.event_type,
  }));
}

export async function getEvents(fetch: typeof globalThis.fetch): Promise<MusmgrEvent[]> {
  return await getEventsWithUrl(fetch, `${getApiUrl()}/events`);
}

export async function createEventPiece(
  fetch: typeof globalThis.fetch,
  eventId: string,
  pieceId: string,
): Promise<void> {
  const url = `${getApiUrl()}/pieces/${pieceId}/events/${eventId}`;
  const res = await fetch(url, {
    method: "POST",
  });

  if (!res.ok) {
    await handleFetchError(res, url);
  }
}

export async function getPieceEvents(
  fetch: typeof globalThis.fetch,
  pieceId: string,
): Promise<MusmgrEvent[]> {
  return await getEventsWithUrl(fetch, `${getApiUrl()}/pieces/${pieceId}/events`);
}
