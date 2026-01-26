import { getEvents } from "$lib/server/events";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const events = await getEvents(fetch);

  return { events };
};
