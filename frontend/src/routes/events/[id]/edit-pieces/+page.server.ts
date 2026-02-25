import type { Actions } from "@sveltejs/kit";
import { error, redirect } from "@sveltejs/kit";
import { createEventPiece } from "$lib/server/events";
import { getPieces } from "$lib/server/pieces";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch, parent }) => {
  const { pieces: eventPieces, event } = await parent();
  const pieces = await getPieces(fetch);
  const inEvent = new Set(eventPieces.map((p) => p.id));

  return {
    event,
    pieces: pieces.filter((p) => !inEvent.has(p.id)),
  };
};

export const actions: Actions = {
  createEventPieces: async ({ fetch, params, request }) => {
    if (!params.id) {
      throw error(500);
    }

    const eventId = params.id;

    const formData = await request.formData();
    const pieceIds = formData.getAll("piece_ids").map((v) => String(v));
    await Promise.all(pieceIds.map((pId) => createEventPiece(fetch, eventId, pId)));

    throw redirect(303, `/events/${eventId}`);
  },
};
