import { type Actions, fail, redirect } from "@sveltejs/kit";
import { type CreatePiecePayload, createPiece, getInstrumentationNames } from "$lib/server/pieces";
import { checkStringFormField } from "$lib/utils";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  return { instrumentationNames: (await getInstrumentationNames(fetch)).sort() };
};

export const actions: Actions = {
  create: async ({ fetch, request }) => {
    const rawData = await request.formData();

    console.log(rawData);

    let data: CreatePiecePayload;
    try {
      data = {
        composed_at: checkStringFormField(rawData.get("composed_at"), "composed_at"),
        description: checkStringFormField(rawData.get("description"), "description"),
        instrumentation: checkStringFormField(rawData.get("instrumentation"), "instrumentation"),
        title: checkStringFormField(rawData.get("title"), "title"),
      };
    } catch (e) {
      return fail(400, e);
    }

    const newLocation = await createPiece(fetch, data);

    throw redirect(303, newLocation);
  },
};
