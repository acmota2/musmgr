import { type Actions, fail, redirect } from "@sveltejs/kit";
import { type CreateComposerPayload, createComposer } from "$lib/server/composer";
import { checkStringFormField } from "$lib/utils";

export const actions: Actions = {
  create: async ({ fetch, request }) => {
    const rawData = await request.formData();

    let data: CreateComposerPayload;
    try {
      data = {
        full_name: checkStringFormField(rawData.get("full_name"), "full_name"),
        biography: checkStringFormField(rawData.get("biography"), "biography"),
      };
    } catch (e) {
      const err = e as Error;
      return fail(400, err.message);
    }

    const newLocation = await createComposer(fetch, data);

    throw redirect(303, newLocation);
  },
};
