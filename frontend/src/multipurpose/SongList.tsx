import { Link } from "react-router-dom";
import "../styles/SongList.scss";
import { Song } from "../CreateSong";

const SongList = ({ data: songList }: { data: Song[] }) => {
  return (
    <ul className="SongList">
      {songList.map((song, index) => {
        const songRef = `/songs?name=${song.name}&id=${song.id}`;

        return (
          <li className="songEntry" key={`link_${song.id}`}>
            <Link className="songLink" to={songRef} key={`link_${index}`}>
              <h2>{song.name}</h2>
              <p>Tonalidade: {song.tonality}</p>
            </Link>
            <button>Editar</button>
          </li>
        );
      })}
    </ul>
  );
};

export default SongList;
