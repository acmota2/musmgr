import { error } from "@sveltejs/kit";

export interface CreatePiecePayload {
  composed_at: string;
  instrumentation: string;
  title: string;
  description: string;
}

export interface GetPieceResponse {
  id: string;
  composed_at: string;
  instrumentation: string;
  title: string;
  description: string;
  created_at: Date;
  updated_at: Date;
}

export interface Piece {
  id: string;
  composedAt: string;
  instrumentation: string;
  title: string;
  description: string;
}

export async function createPiece(
  fetch: typeof globalThis.fetch,
  data: CreatePiecePayload,
): Promise<string> {
  const res = await fetch("/pieces", {
    method: "POST",
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    throw error(500);
  }

  const location = res.headers.get("Location");
  if (location === null) {
    throw error(500);
  }

  return `/api${location}`;
}

export async function getPieces(fetch: typeof globalThis.fetch): Promise<Piece[]> {
  const res = await fetch("/api/pieces");

  const piecesData: GetPieceResponse[] = await res.json();

  if (!res.ok) {
    throw error(500);
  }

  return piecesData.map((piece) => ({
    id: piece.id,
    composedAt: piece.composed_at,
    instrumentation: piece.instrumentation,
    title: piece.title,
    description: piece.description,
  }));
}
