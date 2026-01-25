import { getPieces } from "$lib/server/pieces";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  const pieces = await getPieces();

  return { pieces };
};
