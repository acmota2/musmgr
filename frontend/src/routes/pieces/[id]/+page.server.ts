import { IS_ADMIN } from "$lib/app";
import { getPiece, getPieceFiles } from "$lib/server/pieces";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch, parent, params }) => {
  const { pieces } = await parent();

  const pieceFiles = await getPieceFiles(fetch, params.id);
  const currentPiece = !pieces
    ? await getPiece(fetch, params.id)
    : pieces.find(({ id }) => params.id === id);

  return {
    scores: pieceFiles.filter(({ fileType }) => fileType.startsWith("score")),
    audio: pieceFiles.find(({ fileType }) => fileType === "audio"),
    piece: currentPiece,
  };
};
