import type { PageLoad } from "./$types";

export const load: PageLoad = () => {
  return {
    posts: [
      { id: "1", type: "III Domingo" },
      { id: "2", type: "Acampamento" },
    ],
  };
};
