import { error, redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import { isAdmin } from "$lib/app";
import { getComposer } from "$lib/server/composer";

export const load: LayoutServerLoad = async ({ fetch }) => {
  const composer = await getComposer(fetch);

  if (!isAdmin()) {
    throw error(404);
  }

  if (composer !== null) {
    throw redirect(302, "/");
  }

  return {
    layoutWidth: "narrow",
  };
};
