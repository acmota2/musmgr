import { error } from "@sveltejs/kit";
import { generalServerError } from "./utils";

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

export async function getEventTypes(fetch: typeof globalThis.fetch): Promise<string[]> {
  const res = await fetch("/api/event_types");

  if (!res.ok) {
    throw error(500);
  }

  return await res.json();
}

export async function createEvent(
  fetch: typeof globalThis.fetch,
  data: CreateEventPayload,
): Promise<string> {
  const res = await fetch("/api/events", {
    method: "POST",
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    throw error(500, generalServerError);
  }

  const location = res.headers.get("Location");
  if (location === null) {
    throw error(500, generalServerError);
  }

  return location;
}

export async function getEvent(fetch: typeof globalThis.fetch, id: string): Promise<MusmgrEvent> {
  const res = await fetch(`/api/events/${id}`);

  if (!res.ok) {
    throw error(res.status === 404 ? 404 : 500);
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
    throw error(500, generalServerError);
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
  return await getEventsWithUrl(fetch, "/api/events");
}

export async function getPieceEvents(
  fetch: typeof globalThis.fetch,
  pieceId: string,
): Promise<MusmgrEvent[]> {
  return await getEventsWithUrl(fetch, `/api/pieces/${pieceId}/events`);
}
