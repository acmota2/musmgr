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

/*
export const load: PageServerLoad = () => {
  return {
    hasEvents: true,
    events: {
      "2025": {
        November: [
          {
            id: "123458",
            name: "Morte da Bezerra",
            happenedAt: "2025-11-30",
            eventType: "Other",
          },
          {
            id: "123456",
            name: "Aniversário da Bezerra",
            happenedAt: "2025-11-19",
            eventType: "Festival",
          },
        ],
        August: [
          {
            id: "123457",
            name: "O Rei vai nu",
            happenedAt: "2025-08-15",
            eventType: "Festival",
            description: "The king goes naked down the road",
          },
        ],
      },
    } satisfies PageMusmgrEvents,
  };
};
*/
