import { getApiUrl } from "./api";
import { handleFetchError } from "./utils";

export async function getFileTypes(fetch: typeof globalThis.fetch): Promise<string[]> {
  const url = `${getApiUrl()}/file_types`;
  const res = await fetch(url);

  if (!res.ok) {
    await handleFetchError(res, url);
  }

  return await res.json();
}

// Location header will remain unused for now
export async function createFile(
  fetch: typeof globalThis.fetch,
  pieceId: string,
  data: FormData,
): Promise<void> {
  const url = `${getApiUrl()}/pieces/${pieceId}/files`;
  const res = await fetch(url, {
    method: "POST",
    body: data,
  });

  if (!res.ok) {
    await handleFetchError(res, url);
  }
}
