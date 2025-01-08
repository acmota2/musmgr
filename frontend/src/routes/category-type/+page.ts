import type { PageLoad } from "./$types";

export const load: PageLoad = () => {
  return {
    posts: [
      { title: "cenas", id: "1" },
      { title: "bla", id: 2 },
    ],
  };
};
