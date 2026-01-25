import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ parent, url }) => {
  const { composer } = await parent();

  if (!composer && url.pathname !== "/setup") {
    throw redirect(302, "/setup");
  }

  return {};
};
