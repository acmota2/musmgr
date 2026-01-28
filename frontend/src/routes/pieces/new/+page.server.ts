import { type Actions, fail, redirect } from "@sveltejs/kit";
import { type CreatePiecePayload, createPiece } from "$lib/server/pieces";
import { checkStringFormField } from "$lib/utils";

export const actions: Actions = {
  create: async ({ fetch, request }) => {
    const rawData = await request.formData();

    let data: CreatePiecePayload;
    try {
      data = {
        composed_at: checkStringFormField(rawData.get("compose_at"), "compose_at"),
        description: checkStringFormField(rawData.get("description"), "description"),
        instrumentation: checkStringFormField(rawData.get("instrumentation"), "instrumentation"),
        title: checkStringFormField(rawData.get("title"), "title"),
      };
    } catch (e) {
      throw fail(400, e);
    }

    const newLocation = await createPiece(fetch, data);

    throw redirect(303, newLocation);
  },
};
