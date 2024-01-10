import { SyllableProps } from "../syllables";
import { SongText, emptySyllable } from "./structure";
import { Interval, NOTESORDER, NoteName } from "./transposer";

export const NOTENAMES = ["Dó", "Ré", "Mi", "Fá", "Sol", "Lá", "Si"];

export type SongState = [
  SongText,
  React.Dispatch<React.SetStateAction<SongText>>,
];

type SyllableUpdater = SyllableProps; // sim, é um alias

export function setSongFromSyllable({
  songState: [song, setSong],
  syllable,
  stanceIndex,
  verseIndex,
  wordIndex,
  syllableIndex,
}: SyllableUpdater) {
  song.stances[stanceIndex].text[verseIndex][wordIndex][syllableIndex] =
    syllable;
  setSong({ ...song });
}

export function addStanceToSong({
  songState: [song, setSong],
  stanceIndex,
}: SyllableUpdater) {
  song.stances.push({ type: stanceIndex, text: [[[emptySyllable()]]] });
  setSong({ ...song });
}

export function addVerseToSong({
  songState: [song, setSong],
  stanceIndex,
}: SyllableUpdater) {
  song.stances[stanceIndex].text.push([[emptySyllable()]]);
  setSong({ ...song });
}

export function addWordToSong({
  songState: [song, setSong],
  stanceIndex,
  verseIndex,
}: SyllableUpdater) {
  song.stances[stanceIndex].text[verseIndex].push([emptySyllable()]);
  setSong({ ...song });
}

export function addSyllableToSong({
  songState: [song, setSong],
  stanceIndex,
  verseIndex,
  wordIndex,
}: SyllableUpdater) {
  song.stances[stanceIndex].text[verseIndex][wordIndex].push(emptySyllable());
  setSong({ ...song });
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
