import type { PageLoad } from "./$types";

export const load: PageLoad = () => {
  return {
    posts: [
      { id: "1", type: "Liturgia" },
      { id: "2", type: "Acampamento" },
    ],
  };
};
