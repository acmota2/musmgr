import { error } from "@sveltejs/kit";

interface GetEventResponse {
  id: string;
  name: string;
  happened_at: string;
  description?: string;
  event_type: string;
  created_at: Date;
  updated_at: Date;
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

async function getEventsWithUrl(
  fetch: typeof globalThis.fetch,
  url: string,
): Promise<MusmgrEvent[]> {
  const res = await fetch(url);

  const eventsData: GetEventResponse[] = (await res.json()) || [];

  if (!res.ok) {
    throw error(500);
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
