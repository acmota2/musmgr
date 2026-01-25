import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ parent }) => {
  const { composer } = await parent();

  if (!!composer) {
    throw redirect(302, "/");
  }

  return {};
};
