import { redirect } from "@sveltejs/kit";
import { IS_ADMIN } from "$lib/app";
import { getComposer } from "$lib/server/composer";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ fetch, url }) => {
  const composer = await getComposer(fetch);

  if (!composer && IS_ADMIN && !url.pathname.startsWith("/setup")) {
    throw redirect(302, "/setup");
  }

  return { composer };
};
