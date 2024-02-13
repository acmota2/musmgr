import { SyllableProps } from "../syllables";
import { SongStance, SongText, emptySyllable } from "./structure";
import { Interval, NOTESORDER, NoteName } from "./transposer";

export const NOTENAMES = ["Dó", "Ré", "Mi", "Fá", "Sol", "Lá", "Si"];

export type SongState = [
  SongText,
  React.Dispatch<React.SetStateAction<SongText>>,
];

export type CurrentStanceState = [
  SongStance,
  React.Dispatch<React.SetStateAction<SongStance>>,
];

type SyllableUpdater = SyllableProps; // sim, é um alias

export function setSongFromSyllable({
  songState: [song, setSong],
  stanceState: [, setStance],
  syllable,
  stanceIndex,
  verseIndex,
  wordIndex,
  syllableIndex,
}: SyllableUpdater) {
  song.stances[stanceIndex].text[verseIndex][wordIndex][syllableIndex] =
    syllable;
  setSong({ ...song });
  setStance({ ...song.stances[stanceIndex] });
}

export function addStanceToSong({
  songState: [song, setSong],
  stanceState: [, setStance],
}: SyllableUpdater) {
  const type = song.stances.length;
  song.stances.push({ type, text: [[[emptySyllable()]]] });
  setSong({ ...song });
  setStance({ ...song.stances[type] });
}

type StanceUpdater = {
  songState: SongState;
  stanceIndex: number;
  stance: SongStance;
};

export function modifyStanceInSong({
  songState: [song, setSong],
  stanceIndex,
  stance,
}: StanceUpdater) {
  song.stances[stanceIndex] = stance;
  console.log("song", song);
  setSong(song);
}

type VerseUpdater =
  | SyllableUpdater
  | {
      songState: SongState;
      stanceState: CurrentStanceState;
      stanceIndex: number;
    };

export function addVerseToSong({
  songState: [song, setSong],
  stanceState: [, setStance],
  stanceIndex,
}: VerseUpdater) {
  song.stances[stanceIndex].text.push([[emptySyllable()]]);
  setSong({ ...song });
  setStance({ ...song.stances[stanceIndex] });
}

export function removeVerseFromSong({
  songState: [song, setSong],
  stanceState: [, setStance],
  stanceIndex,
}: VerseUpdater) {
  song.stances[stanceIndex].text.pop();
  setSong({ ...song });
  setStance({ ...song.stances[stanceIndex] });
}

type WordUpdater =
  | SyllableUpdater
  | {
      songState: SongState;
      stanceState: CurrentStanceState;
      stanceIndex: number;
      verseIndex: number;
    };

export function addWordToSong({
  songState: [song, setSong],
  stanceState: [, setStance],
  stanceIndex,
  verseIndex,
}: WordUpdater) {
  song.stances[stanceIndex].text[verseIndex].push([emptySyllable()]);
  setSong({ ...song });
  setStance({ ...song.stances[stanceIndex] });
}

export function removeWordFromSong({
  songState: [song, setSong],
  stanceState: [, setStance],
  stanceIndex,
  verseIndex,
}: WordUpdater) {
  song.stances[stanceIndex].text[verseIndex].pop();
  setSong({ ...song });
  setStance({ ...song.stances[stanceIndex] });
}

type SyllableAdder =
  | SyllableUpdater
  | {
      songState: SongState;
      stanceState: CurrentStanceState;
      stanceIndex: number;
      verseIndex: number;
      wordIndex: number;
    };

export function addSyllableToSong({
  songState: [song, setSong],
  stanceState: [, setStance],
  stanceIndex,
  verseIndex,
  wordIndex,
}: SyllableAdder) {
  song.stances[stanceIndex].text[verseIndex][wordIndex].push(emptySyllable());
  setSong({ ...song });
  setStance({ ...song.stances[stanceIndex] });
}

export function removeSyllableFromSong({
  songState: [song, setSong],
  stanceState: [, setStance],
  stanceIndex,
  verseIndex,
  wordIndex,
}: SyllableAdder) {
  song.stances[stanceIndex].text[verseIndex][wordIndex].pop();
  setSong({ ...song });
  setStance({ ...song.stances[stanceIndex] });
}

export function stringToNoteName(noteName: string): NoteName {
  return NOTESORDER[NOTENAMES.indexOf(noteName)];
}

export type StringState = [
  string,
  React.Dispatch<React.SetStateAction<string>>,
];

export type NumberState = [
  number,
  React.Dispatch<React.SetStateAction<number>>,
];

export type IntervalState = {
  noteName: StringState;
  semitones: NumberState;
};

export function stateToInterval(i: IntervalState): Interval {
  const {
    noteName: [noteName],
    semitones: [semitones],
  } = i;
  return {
    distanceToNoteIndex: NOTENAMES.indexOf(noteName),
    semitoneDistance: semitones,
  };
}
