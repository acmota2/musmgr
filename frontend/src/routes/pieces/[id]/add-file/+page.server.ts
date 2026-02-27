import type { Actions } from "@sveltejs/kit";
import { error, fail, redirect } from "@sveltejs/kit";
import { IS_ADMIN } from "$lib/app";
import { createFile, getFileTypes } from "$lib/server/files";
import { checkFileFormField, checkStringFormField } from "$lib/utils";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch, parent }) => {
  if (!IS_ADMIN) {
    throw error(404);
  }

  const { piece } = await parent();
  const fileTypes = await getFileTypes(fetch);

  return { piece, fileTypes };
};

export const actions: Actions = {
  createFile: async ({ fetch, params, request }) => {
    if (!params.id) {
      throw error(500);
    }

    const rawData = await request.formData();

    try {
      const file = checkFileFormField(rawData.get("file"), "file");
      if (file === null) {
        throw new Error("File is mandatory");
      }

      checkStringFormField(rawData.get("name"), "name");
      checkStringFormField(rawData.get("classification"), "classification");
      checkStringFormField(rawData.get("file_type"), "file_type");
    } catch (e) {
      const err = e as Error;
      return fail(400, err.message);
    }

    await createFile(fetch, params.id, rawData);

    throw redirect(303, `/pieces/${params.id}`);
  },
};
