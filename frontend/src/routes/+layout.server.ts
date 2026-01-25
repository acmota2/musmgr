import { getComposer } from "$lib/server/composer";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async () => {
  const composer = await getComposer();
  return { composer };
};
