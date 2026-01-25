import { getEvents } from "$lib/server/events";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  const events = await getEvents();

  return { events };
};
