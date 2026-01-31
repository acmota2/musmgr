import type { Actions } from "@sveltejs/kit";
import { updateComposer, updateComposerPicture } from "$lib/server/composer";
import { checkFileFormField, checkStringFormField } from "$lib/utils";

export const actions: Actions = {
  updateComposer: async ({ fetch, request }) => {
    const rawData = await request.formData();

    const picture = checkFileFormField(rawData.get("picture"), "picture");
    const bio = checkStringFormField(rawData.get("biography"), "biography");

    await updateComposer(fetch, { biography: bio });
    if (picture !== null) {
      await updateComposerPicture(fetch, picture);
    }

    return { success: true };
  },
};
