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

export async function getEvents(fetch: typeof globalThis.fetch): Promise<MusmgrEvent[]> {
  const res = await fetch("/api/events");

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
