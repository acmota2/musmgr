import { error } from "@sveltejs/kit";

export async function getFileTypes(fetch: typeof globalThis.fetch): Promise<string[]> {
  const res = await fetch("/api/file_types");

  if (!res.ok) {
    throw error(500);
  }

  return await res.json();
}

// Location header will remain unused for now
export async function createFile(
  fetch: typeof globalThis.fetch,
  pieceId: string,
  data: FormData,
): Promise<void> {
  const res = await fetch(`/api/pieces/${pieceId}/files`, {
    method: "POST",
    body: data,
  });

  if (!res.ok) {
    throw error(500);
  }
}
