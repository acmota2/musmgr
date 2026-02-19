import { eventListToTimetable, getEvents, type MusmgrEventTimetable } from "$lib/server/events";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const rawEvents = await getEvents(fetch);

  const events: MusmgrEventTimetable = eventListToTimetable(rawEvents);

  return { events, hasEvents: rawEvents.length > 0 };
};
