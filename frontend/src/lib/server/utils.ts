import { error } from "@sveltejs/kit";

export const generalServerError = {
  code: "SERVER_ERROR",
  message: "It's not you, it's us",
  status: 500,
};

/**
 * Logs fetch failure details to the server terminal and throws a SvelteKit error.
 * Use this to replace bare `throw error(500)` calls so you can see actual backend responses.
 */
export async function handleFetchError(res: Response, url: string): Promise<never> {
  let body = "(no body)";
  try {
    body = await res.text();
  } catch {
    // ignore parse failures
  }

  console.error(`[API ERROR] ${res.status} ${res.statusText} fetching ${url}\nResponse: ${body}`);
  throw error(res.status === 404 ? 404 : 500, generalServerError);
}
