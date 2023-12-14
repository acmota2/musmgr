import { useState } from "react";
import { SongStance } from "./song_text/structure";

interface StanceProp {
  stance: SongStance;
}

const Stance = ({ stance }: StanceProp) => {
  const [type, setType] = useState<string>(stance.type.toString());
  return (
    <div className="stance">
      <input
        type="text"
        placeholder="Número da estrofe..."
        value={type}
        onChange={(e) => setType(e.target.value)}
      />
      <div className="stanceText">
        
      </div>
    </div>
  );
};

export default Stance;
