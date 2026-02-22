import { error } from "@sveltejs/kit";
import { getEvent } from "$lib/server/events";
import type { LayoutServerLoad } from "./$types";
import { getEventPieces } from "$lib/server/pieces";

export const load: LayoutServerLoad = async ({ fetch, params }) => {
  const currentEvent = await getEvent(fetch, params.id);

  if (!currentEvent) {
    throw error(500, {
      status: 500,
      message: "Error while finding the piece",
    });
  }

  const eventPieces = await getEventPieces(fetch, params.id);

  console.log(eventPieces);

  return {
    event: currentEvent,
    pieces: eventPieces,
  };
};
