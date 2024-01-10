import { useState } from "react";
import { SongStance } from "./song_text/structure";
import Verses from "./verses";
import { SongState } from "./song_text/songState";

interface StanceProp {
  songState: SongState;
  stance: SongStance;
  stanceIndex: number;
}

const Stances = ({ songState, stance, stanceIndex }: StanceProp) => {
  const [type, setType] = useState<string>(stance.type.toString());
  return (
    <div className="stance" key={`stance:${stanceIndex}`}>
      <input
        type="text"
        placeholder="Número da estrofe..."
        value={type}
        onChange={(e) => setType(e.target.value)}
      />
      <div className="stanceText">
        {stance.text.map((v, i) => (
          <div className="stances" key={`verse:${i}`}>
            <Verses
              songState={songState}
              verse={v}
              stanceIndex={stanceIndex}
              verseIndex={i}
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default Stances;
