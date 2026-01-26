import { getComposer } from "$lib/server/composer";
import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import { isAdmin } from "$lib/app";

export const load: LayoutServerLoad = async ({ fetch, url }) => {
  const composer = await getComposer(fetch);

  if (!composer && isAdmin() && !url.pathname.startsWith("/setup")) {
    throw redirect(302, "/setup");
  }

  return { composer };
};
