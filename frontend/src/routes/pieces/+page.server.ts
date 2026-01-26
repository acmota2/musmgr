import { getPieces } from "$lib/server/pieces";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const pieces = await getPieces(fetch);

  return { pieces };
};
