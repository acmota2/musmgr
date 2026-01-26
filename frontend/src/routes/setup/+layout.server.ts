import { error, redirect } from "@sveltejs/kit";
import { isAdmin } from "$lib/app";
import { getComposer } from "$lib/server/composer";
import type { LayoutServerLoad } from "./$types";

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
