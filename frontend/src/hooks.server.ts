import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
  const frontendVisibilityMode = process.env.FRONTEND_VISIBILITY_MODE || "public";

  if (!["public", "admin"].includes(frontendVisibilityMode)) {
    return new Response("You must set FRONTEND_VISIBILITY_MODE to 'public' or 'admin'", {
      status: 500,
    });
  }

  return await resolve(event);
};
