interface GetPieceResponse {
  id: string;
  composed_at: string;
  instrumentation: string;
  title: string;
  description: string;
  audio_content_type: string;
  audio: string;
  created_at: Date;
  updated_at: Date;
}

export interface Piece {
  id: string;
  composedAt: string;
  title: string;
  description: string;
}

export async function getPieces(): Promise<Piece[]> {
  const res = await fetch("/pieces");

  const piecesData: GetPieceResponse[] = await res.json();

  return piecesData.map((piece) => ({
    id: piece.id,
    composedAt: piece.composed_at,
    title: piece.title,
    description: piece.description,
  }));
}
