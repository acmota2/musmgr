import { useSearchParams } from "react-router-dom";
import { Song } from "./CreateSong";
import SongList from "./multipurpose/SongList";
import AnyList from "./multipurpose/anylist";
import TitlePage from "./multipurpose/titlepage";

type SongElement = {
  path: string;
  title: string;
  searchParam?: string[];
};

export default function Songs({ path, searchParam }: SongElement) {
  const [params] = useSearchParams();

  const makeTitle = () => {
    let title = `Músicas de A-Z`;
    if (searchParam) {
      for (let i = 0; i < searchParam.length; ++i) {
        title += params.get(searchParam[i]) + " ";
      }
    }
  };

  return (
    <TitlePage title={`Músicas de `}>
      <AnyList
        path={searchParam ? `${path}/${params.get(searchParam)}` : path}
        generator={(data: Song[]) => <SongList data={data} />}
      />
    </TitlePage>
  );
}
