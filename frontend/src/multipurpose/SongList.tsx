import { Link } from "react-router-dom";
import "../styles/SongList.scss";
import { Song } from "../CreateSong";
import { noteNameToString } from "../SongEditor/song_text/transposer";

const SongList = ({ data: songList }: { data: Song[] }) => {
  return (
    <ul className="SongList">
      {songList.map((song: Song, index) => {
        const songUrl = `/song-editor?name=${song.name}&tonality=${noteNameToString(song.tonality_root)}`;
        const tonality = `${noteNameToString(song.tonality_root)} ${song.tonality_details}`;
        return (
          <li className="songEntry" key={`link_${song.id}`}>
            <Link className="songLink" to={songUrl} key={`link_${index}`}>
              <h2>{song.name}</h2>
              <p>Tonalidade: {tonality}</p>
            </Link>
          </li>
        );
      })}
    </ul>
  );
};

export default SongList;
