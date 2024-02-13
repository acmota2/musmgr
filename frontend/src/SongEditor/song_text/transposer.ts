import { NOTENAMES } from "./songState";

export enum NoteName {
  C = 0,
  D = 2,
  E = 4,
  F = 5,
  G = 7,
  A = 9,
  B = 11,
}

export function noteNameToString(note: NoteName) {
  switch (note) {
    case NoteName.C:
      return "Dó";
    case NoteName.D:
      return "Ré";
    case NoteName.E:
      return "Mi";
    case NoteName.F:
      return "Fá";
    case NoteName.G:
      return "Sol";
    case NoteName.A:
      return "Lá";
    case NoteName.B:
      return "Si";
  }
}

export type Interval = {
  readonly distanceToNoteIndex: number;
  readonly semitoneDistance: number;
};

export type Note = {
  index: number;
  alteration: number;
  absoluteCount: (n: Note) => number;
  transpose: (n: Note, i: Interval) => Note;
  toString: (n: Note) => string;
};

export const NOTESORDER: NoteName[] = [
  NoteName.C,
  NoteName.D,
  NoteName.E,
  NoteName.F,
  NoteName.G,
  NoteName.A,
  NoteName.B,
];

export function newNote(note: NoteName, alteration: number): Note {
  let i = 0;
  while (NOTESORDER[i] !== note) ++i;
  return {
    index: i,
    alteration,
    absoluteCount,
    transpose,
    toString: noteToString,
  };
}

function noteToString(n: Note): string {
  let alteration = "";
  if (n.alteration < 0) {
    alteration += "♭".repeat(n.alteration * -1);
  } else {
    alteration += "#".repeat(n.alteration);
  }
  return NOTENAMES[n.index] + alteration;
}

const absoluteCount = (n: Note) => {
  return NOTESORDER[n.index] + n.alteration;
};

function transpose(n: Note, i: Interval): Note {
  const semitonesAfterIndexing =
    NOTESORDER[i.distanceToNoteIndex] - NOTESORDER[n.index];
  n.index = i.distanceToNoteIndex - n.index;
  n.alteration = semitonesAfterIndexing - NOTESORDER[n.index];
  return n;
}

function transposeNote(from: Note, to: Note): Interval {
  const distanceToNote = from.index - to.index;
  const semitones = absoluteCount(from) - absoluteCount(to);
  return {
    distanceToNoteIndex: distanceToNote,
    semitoneDistance: semitones,
  };
}

export function transposeSemitones(from: Note, st: number): Interval {
  const target = sensicalTranspose[absoluteCount(from) + st];
  return transposeNote(from, target);
}

const sensicalTranspose: Note[] = [
  newNote(NoteName.C, 0), // Dó
  newNote(NoteName.D, -1), // Réb
  newNote(NoteName.D, 0), // Ré
  newNote(NoteName.E, -1), // Mib
  newNote(NoteName.E, 0), // Mi
  newNote(NoteName.F, 0), // Fá
  newNote(NoteName.F, 1), // Fá#
  newNote(NoteName.G, 0), // Sol
  newNote(NoteName.A, -1), // Láb
  newNote(NoteName.A, 0), // Lá
  newNote(NoteName.B, -1), // Sib
  newNote(NoteName.B, 0), // Si
];

export type Tonality = {
  details: string;
  note: Note;
};
