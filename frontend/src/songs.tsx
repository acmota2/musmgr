import { Song } from "./CreateSong";
import SongList from "./multipurpose/SongList";
import AnyList from "./multipurpose/anylist";
import TitlePage from "./multipurpose/titlepage";

const Songs = () => (
  <TitlePage title="Músicas de A-Z">
    <AnyList
      path="/songs"
      generator={(data: Song[]) => <SongList data={data} />}
    />
  </TitlePage>
);

export default Songs;
