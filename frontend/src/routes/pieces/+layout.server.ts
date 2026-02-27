import { getPieces } from "$lib/server/pieces";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ fetch }) => {
  const pieces = await getPieces(fetch);

  return { pieces };
};
