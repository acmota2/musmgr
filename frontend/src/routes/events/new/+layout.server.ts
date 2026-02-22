import { IS_ADMIN } from "$lib/app";
import { error } from "node:console";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async () => {
  if (!IS_ADMIN) {
    throw error(404, {
      message: "Couldn't find the specific resource",
      status: 404,
    });
  }
};
