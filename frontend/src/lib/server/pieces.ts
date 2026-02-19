import { error } from "@sveltejs/kit";

export interface CreatePiecePayload {
  composed_at: string;
  instrumentation: string;
  title: string;
  description: string;
}

export interface PieceResponse {
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

// revise this in the backend, for now this works
// classification, fileType, parentId, createdAt, updatedAt are not necessary
export interface PieceFileResponse {
  id: string;
  classification: number;
  content_type: string;
  description: string;
  name: string;
  origin: string;
  file_type: string;
  parent_id: string;
  piece_id: string;
  created_at: Date;
  updated_at: Date;
}

export interface PieceFile {
  id: string;
  contentType: string;
  description: string;
  fileType: string;
  name: string;
  pieceId: string;
}

export async function createPiece(
  fetch: typeof globalThis.fetch,
  data: CreatePiecePayload,
): Promise<string> {
  const res = await fetch("/api/pieces", {
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

  return location;
}

export async function getPieces(fetch: typeof globalThis.fetch): Promise<Piece[]> {
  const res = await fetch("/api/pieces");

  const piecesData: PieceResponse[] = (await res.json()) || [];

  if (!res.ok) {
    throw error(500);
  }

  return piecesData.map((piece) => ({
    ...piece,
    composedAt: piece.composed_at,
  }));
}

export async function getPiece(fetch: typeof globalThis.fetch, id: string): Promise<Piece> {
  const res = await fetch(`/api/piece/${id}`);

  if (!res.ok) {
    throw error(res.status === 404 ? 404 : 500);
  }

  const piece: PieceResponse = await res.json();

  return {
    ...piece,
    composedAt: piece.composed_at,
  };
}

export async function getInstrumentationNames(fetch: typeof globalThis.fetch): Promise<string[]> {
  const res = await fetch("/api/instrumentation_names");

  if (!res.ok) {
    throw error(500);
  }

  return await res.json();
}

export async function getPieceFiles(
  fetch: typeof globalThis.fetch,
  id: string,
): Promise<PieceFile[]> {
  const res = await fetch(`/api/pieces/${id}/files`);

  const pieceFiles = (await res.json()) || [];
  if (!res.ok) {
    throw error(500);
  }

  return pieceFiles.map((file: PieceFileResponse) => ({
    id: file.id,
    description: file.description,
    name: file.name,
    contentType: file.content_type,
    fileType: file.file_type,
    pieceId: file.piece_id,
  }));
}
