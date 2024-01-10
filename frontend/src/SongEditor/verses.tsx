import { SongState } from "./song_text/songState";
import { Verse, Word } from "./song_text/structure";
import Syllables from "./syllables";

type VerseProps = {
  songState: SongState;
  verse: Verse;
  stanceIndex: number;
  verseIndex: number;
};

const Verses = ({ songState, verse, verseIndex, stanceIndex }: VerseProps) => {
  return (
    <div className="verse">
      {verse.map((w: Word, i: number) => (
        <div className="verses">
          {w.map((s, j) => (
            <Syllables
              songState={songState}
              stanceIndex={stanceIndex}
              verseIndex={verseIndex}
              wordIndex={i}
              syllableIndex={j}
              syllable={s}
            />
          ))}
        </div>
      ))}
    </div>
  );
};

export default Verses;
