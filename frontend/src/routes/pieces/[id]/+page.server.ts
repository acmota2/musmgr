import { error } from "@sveltejs/kit";
import { eventListToTimetable, getPieceEvents } from "$lib/server/events";
import { getPiece, getPieceFiles } from "$lib/server/pieces";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch, parent, params }) => {
  const { pieces } = await parent();
  // control when it comes from a redirect
  const currentPiece = !pieces
    ? await getPiece(fetch, params.id)
    : pieces.find(({ id }) => params.id === id);

  if (!currentPiece) {
    throw error(500, {
      status: 500,
      message: "Error while finding the piece",
    });
  }

  const pieceFiles = await getPieceFiles(fetch, params.id);
  const pieceEvents = await getPieceEvents(fetch, params.id);

  return {
    scores: pieceFiles.filter(({ fileType }) => fileType.startsWith("score")),
    audios: pieceFiles.filter(({ fileType }) => fileType.startsWith("audio")),
    images: pieceFiles.filter(
      ({ fileType }) => !fileType.startsWith("audio") || !fileType.startsWith("score"),
    ),
    piece: currentPiece,
    events: eventListToTimetable(pieceEvents),
  };
};
