import { error, redirect } from "@sveltejs/kit";
import { IS_ADMIN } from "$lib/app";
import { getComposer } from "$lib/server/composer";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ fetch }) => {
  const composer = await getComposer(fetch);

  if (!IS_ADMIN) {
    error(404, {
      code: "NOT_FOUND",
      message: "This page doesn't exist",
      status: 404,
    });
  }

  if (composer !== null) {
    throw redirect(302, "/");
  }

  return {
    layoutWidth: "narrow",
  };
};
