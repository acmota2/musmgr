interface GetEventResponse {
  id: string;
  happened_at: string;
  description: string;
  event_type: string;
  created_at: Date;
  updated_at: Date;
}

export interface Event {
  id: string;
  happenedAt: string;
  description: string;
  eventType: string;
}

export async function getEvents(): Promise<Event[]> {
  const res = await fetch("/events");

  const eventsData: GetEventResponse[] = await res.json();

  return eventsData.map((event) => ({
    id: event.id,
    happenedAt: event.happened_at,
    description: event.description,
    eventType: event.event_type,
  }));
}
