import { error } from "@sveltejs/kit";
import { getEvent } from "$lib/server/events";
import { getEventPieces } from "$lib/server/pieces";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ fetch, params }) => {
  const event = await getEvent(fetch, params.id);

  if (!event) {
    throw error(500, {
      status: 500,
      message: "Error while finding the piece",
    });
  }

  const eventPieces = await getEventPieces(fetch, params.id);

  return {
    event,
    pieces: eventPieces,
  };
};
