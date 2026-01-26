import { redirect } from "@sveltejs/kit";
import { isAdmin } from "$lib/app";
import { getComposer } from "$lib/server/composer";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ fetch, url }) => {
  const composer = await getComposer(fetch);

  if (!composer && isAdmin() && !url.pathname.startsWith("/setup")) {
    throw redirect(302, "/setup");
  }

  return { composer };
};
