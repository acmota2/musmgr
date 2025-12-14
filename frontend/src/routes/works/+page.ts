import type { PageLoad } from "./$types";

export const load: PageLoad = () => {
  return {
    posts: [
      {
        id: "1",
        title: "O senhor é meu pastor",
        key: "C",
        file: ["./hello.txt"],
      },
    ],
  };
};
