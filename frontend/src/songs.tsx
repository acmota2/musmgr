import { useSearchParams } from "react-router-dom";
import { Song } from "./CreateSong";
import SongList from "./multipurpose/SongList";
import AnyList from "./multipurpose/anylist";
import TitlePage from "./multipurpose/titlepage";

export default function Songs() {
  const [params] = useSearchParams();

  const makeTitle = () => {
    let title = "Canções";
    if (params.has("name")) {
      title += ` de ${params.get("name")}`;
    }
    if (params.has("date")) {
      title += ` | ${params.get("date")}`;
    } else if (!params.has("subcategory")) {
      title += " de A-Z";
    }
    return title;
  };
  const makePath = () => {
    if (params.has("event")) {
      return `/songs/event/${params.get("id")}`;
    } else if (params.has("subcategory")) {
      return `/songs/subcategory/${params.get("id")}`;
    }
    return "/songs";
  };

  return (
    <TitlePage title={makeTitle()}>
      <AnyList
        path={makePath()}
        generator={(data: Song[]) => <SongList data={data} />}
      />
    </TitlePage>
  );
}
