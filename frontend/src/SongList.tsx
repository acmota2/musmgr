type Song = {
  id: number;
  name: string;
  tonality: string;
};

const SongList = ({ songs, title }: { songs: Song[]; title: string }) => {
  return (
    <div className="SongList">
      <p>{title}</p>
      {songs.map((song) => {
        const songRef = `/songs/${song.name}`;

        return (
          <div className="songEntry">
            <a href={songRef}></a>
            <p>{song.tonality}</p>
            <div className="songDetails">
              <div>
                <button></button>
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default SongList;
