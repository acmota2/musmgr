import { SongStance } from "./song_text/structure";
import Verses from "./verses";
import "../styles/editor/stances.scss";
import {
  CurrentStanceState,
  SongState,
  addVerseToSong,
  removeVerseFromSong,
} from "./song_text/songState";
import Adder from "./atoms/adder";

interface StanceProp {
  songState: SongState;
  stanceState: CurrentStanceState;
  stance: SongStance;
  stanceIndex: number;
}

const Stances = ({ songState, stanceState, stanceIndex }: StanceProp) => {
  const [stance] = stanceState;
  return (
    <div className="stance" key={`stance:${stanceIndex}`}>
      <p>{stance.type === 0 ? "Refrão" : "Estrofe"}</p>
      {stance.text.map((v, i) => (
        <div className="verses" key={`stance:${stanceIndex}:verse:${i}`}>
          <Verses
            songState={songState}
            stanceState={stanceState}
            verse={v}
            stanceIndex={stanceIndex}
            verseIndex={i}
          />
          <Adder
            elemName="Verso"
            onAdd={() =>
              addVerseToSong({ songState, stanceState, stanceIndex })
            }
            onRemove={() =>
              removeVerseFromSong({ songState, stanceState, stanceIndex })
            }
          />
        </div>
      ))}
    </div>
  );
};

export default Stances;
