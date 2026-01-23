import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
  const userType = process.env.USER_TYPE || "public";

  if (!(userType in ["public", "admin"])) {
    return new Response("You must set USER_TYPE to 'public' or 'admin'", {
      status: 500,
    });
  }

  event.locals.user = userType;

  return await resolve(event);
};
