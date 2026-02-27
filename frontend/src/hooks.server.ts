import type { Handle, HandleServerError } from "@sveltejs/kit";
import { error } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export const handle: Handle = async ({ event, resolve }) => {
  const API_BASE_URL = API_URL;
  const PROXY_PATH = "/api";
  const frontendVisibilityMode = process.env.VITE_FRONTEND_VISIBILITY_MODE;

  if (!API_BASE_URL) {
    throw error(500, {
      message: "You must set API_URL",
      status: 500,
      code: "SERVER_ERROR",
    });
  }

  if (frontendVisibilityMode && !["public", "admin"].includes(frontendVisibilityMode)) {
    throw error(500, {
      message: "You must set VITE_FRONTEND_VISIBILITY_MODE to 'public' or 'admin'",
      status: 500,
      code: "SERVER_ERROR",
    });
  }

  if (event.url.pathname.startsWith(PROXY_PATH)) {
    const strippedPath = event.url.pathname.substring(PROXY_PATH.length);
    const proxiedUrl = new URL(`${API_BASE_URL}${strippedPath}${event.url.search}`);

    const headers = new Headers(event.request.headers);
    headers.delete("content-length");
    let response: Response | null = null;
    try {
      response = await fetch(proxiedUrl.toString(), {
        method: event.request.method,
        headers,
        body: event.request.body,
        duplex: "half",
      } as Request & { duplex: string }); // necessary beause of linter

      return new Response(response.body, {
        status: response.status,
        headers: response.headers,
      });
    } catch (err) {
      console.error("Could not proxy with: ", err);
    }
  }

  return await resolve(event);
};

export const handleError: HandleServerError = async ({ status, message }) => ({
  message,
  status,
  code: status === 404 ? "NOT_FOUND" : "",
});
