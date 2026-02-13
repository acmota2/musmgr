import { getEvents, type MusmgrEvent } from "$lib/server/events";
import { extractMonthFromDate } from "$lib/utils";
import type { PageServerLoad } from "./$types";

export type PageMusmgrEvents = Record<string, Record<string, MusmgrEvent[]>>;

export const load: PageServerLoad = async ({ fetch }) => {
  const rawEvents = await getEvents(fetch);

  const events: PageMusmgrEvents = {};
  for (const event of rawEvents) {
    const [year] = event.happenedAt.split("-");
    if (!(year in events)) {
      events[year] = {};
    }

    const month = extractMonthFromDate(event.happenedAt);
    if (!(month in events[year])) {
      events[year][month] = [];
    }

    const monthEvents = events[year][month];
    monthEvents.push(event);
  }

  return { events, hasEvents: rawEvents.length > 0 };
};
