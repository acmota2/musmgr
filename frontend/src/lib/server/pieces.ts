import { error } from "@sveltejs/kit";
import { getApiUrl } from "./api";
import { handleFetchError } from "./utils";

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
  const url = `${getApiUrl()}/pieces`;
  const res = await fetch(url, {
    method: "POST",
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  const location = res.headers.get("Location");
  if (location === null) {
    throw error(500, {
      code: "INVALID_LOCATION_HEADER",
      message: "It's not you, it's us",
      status: 500,
    });
  }

  return location;
}

export async function getEventPieces(fetch: typeof globalThis.fetch, id: string): Promise<Piece[]> {
  const url = `${getApiUrl()}/events/${id}/pieces`;
  const res = await fetch(url);

  if (res.status !== 200) {
    await handleFetchError(res, url);
  }

  const pieces: PieceResponse[] | null = await res.json();

  return (pieces || []).map((piece) => ({
    id: piece.id,
    composedAt: piece.composed_at,
    instrumentation: piece.instrumentation,
    title: piece.title,
    description: piece.description,
  }));
}

export async function getPieces(fetch: typeof globalThis.fetch): Promise<Piece[]> {
  const url = `${getApiUrl()}/pieces`;
  const res = await fetch(url);

  const piecesData: PieceResponse[] = (await res.json()) || [];

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  return piecesData.map((piece) => ({
    ...piece,
    composedAt: piece.composed_at,
  }));
}

export async function getPiece(fetch: typeof globalThis.fetch, id: string): Promise<Piece> {
  const url = `${getApiUrl()}/pieces/${id}`;
  const res = await fetch(url);

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  const piece: PieceResponse = await res.json();

  return {
    ...piece,
    composedAt: piece.composed_at,
  };
}

export async function getInstrumentationNames(fetch: typeof globalThis.fetch): Promise<string[]> {
  const url = `${getApiUrl()}/instrumentation_names`;
  const res = await fetch(url);

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  return await res.json();
}

export async function getPieceFiles(
  fetch: typeof globalThis.fetch,
  id: string,
): Promise<PieceFile[]> {
  const url = `${getApiUrl()}/pieces/${id}/files`;
  const res = await fetch(url);

  const pieceFiles = (await res.json()) || [];
  if (!res.ok) {
    await handleFetchError(res, url);
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
