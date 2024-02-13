import "../styles/editor/verses.scss";
import Adder from "./atoms/adder";
import {
  CurrentStanceState,
  SongState,
  addSyllableToSong,
  addWordToSong,
  removeSyllableFromSong,
  removeWordFromSong,
} from "./song_text/songState";
import { Verse, Word } from "./song_text/structure";
import Syllables from "./syllables";

type VerseProps = {
  songState: SongState;
  stanceState: CurrentStanceState;
  verse: Verse;
  stanceIndex: number;
  verseIndex: number;
};

const Verses = ({
  songState,
  stanceState,
  verse,
  verseIndex,
  stanceIndex,
}: VerseProps) => {
  return (
    <div className="verse" key={`verse_${verseIndex}:${stanceIndex}`}>
      {verse.map((w: Word, i: number) => (
        <div className="word" key={`word${i}`}>
          {w.map((s, j) => (
            <div className="syllable" key={`${j}`}>
              <Syllables
                songState={songState}
                stanceState={stanceState}
                stanceIndex={stanceIndex}
                verseIndex={verseIndex}
                wordIndex={i}
                syllableIndex={j}
                syllable={s}
              />
            </div>
          ))}
          <Adder
            elemName="Sílaba"
            onAdd={() =>
              addSyllableToSong({
                songState,
                stanceState,
                stanceIndex,
                verseIndex,
                wordIndex: i,
              })
            }
            onRemove={() =>
              removeSyllableFromSong({
                songState,
                stanceState,
                stanceIndex,
                verseIndex,
                wordIndex: i,
              })
            }
          />
        </div>
      ))}
      <Adder
        elemName="Palavra"
        onAdd={() =>
          addWordToSong({ songState, stanceState, stanceIndex, verseIndex })
        }
        onRemove={() =>
          removeWordFromSong({
            songState,
            stanceState,
            stanceIndex,
            verseIndex,
          })
        }
      />
    </div>
  );
};

export default Verses;
