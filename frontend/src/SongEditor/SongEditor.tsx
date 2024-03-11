import "../styles/SongEditor.scss";
import { useSearchParams } from "react-router-dom";
import TitlePage from "../multipurpose/titlepage";
import AnyPostForm from "../multipurpose/AnyPostForm";
import {
  SongStance,
  SongText,
  addStance,
  emptySong,
  emptyStance,
} from "./song_text/structure";
import useGet from "../hooks/useGet";
import ErrorPage from "../multipurpose/errorpage";
import Loading from "../multipurpose/loading";
import Stances from "./stances";
import { useRef, useState } from "react";
import { IntervalState, stringToNoteName } from "./song_text/songState";
import NoteSelector from "./NoteSelector";

export type text = 0;
export type score = 1;
export type FileType = text | score;
export type SongFile = {
  Name: string;
  Open: boolean;
  song_id: number;
  type: FileType;
  text_file: SongText;
};

const Interval = ({ interval }: { interval: IntervalState }) => {
  const {
    noteName: [noteName, setNoteName],
    semitones: [semitones, setSemitones],
  } = interval;

  return (
    <div className="intervalDetails">
      <NoteSelector
        noteValue={noteName}
        noteOnChange={(note) => {
          const semitones = stringToNoteName(note) - stringToNoteName(noteName);
          setNoteName(note);
          setSemitones(semitones);
        }}
        accidentValue={0}
        accidentOnChange={(a) => setSemitones(semitones + a)}
      />
    </div>
  );
};

const MainEditor = () => {
  const [params] = useSearchParams();
  const { data, err, pending } = params.get("create")
    ? { data: emptySong(), err: null, pending: false }
    : useGet<SongText>(`/files/song/text/${params.get("song_id")}`);

  const state = useState<SongText>(JSON.parse(JSON.stringify(data)));
  const [songState, setSongState] = state;

  const tonality = params.get("tonality");
  const curInterval = {
    noteName: useState<string>(tonality ? tonality : "Dó"),
    semitones: useState<number>(0),
  };

  let stanceNumberRef = useRef<number>(0);
  const stanceState = useState<SongStance>(data.stances[0]);
  const [currentStance, setCurrentStance] = stanceState;

  return (
    <TitlePage title={params.get("name") || ""}>
      <AnyPostForm
        path="/file"
        redirectTo={() => "/songs"}
        buttonText="Concluir"
        state={{ songState: state }}
        dataCreator={({ songState: [text_file] }) => ({
          name: params.get("name"),
          type: 0,
          open: false,
          song_id: params.has("id") ? params.get("id") : 0,
          text_file,
        })}
      >
        {err && <ErrorPage err={err} />}
        {pending && <Loading />}
        <div className="showSong">
          <Interval interval={curInterval} />
          {stanceNumberRef.current === 0 ? (
            <div className="placeHolder" />
          ) : (
            <button
              className="stanceChangerTop"
              onClick={() =>
                setCurrentStance(songState.stances[--stanceNumberRef.current])
              }
            >
              ^
            </button>
          )}
          <Stances
            songState={state}
            stanceState={stanceState}
            stance={currentStance}
            stanceIndex={stanceNumberRef.current}
          />
          <button
            className="stanceChangerBottom"
            onClick={() => {
              if (stanceNumberRef.current + 1 === songState.stances.length) {
                const newStance = emptyStance(stanceNumberRef.current + 1);
                setSongState(addStance(songState, newStance));
              }
              setCurrentStance(songState.stances[++stanceNumberRef.current]);
            }}
          >
            {stanceNumberRef.current + 1 === songState.stances.length
              ? "+"
              : "v"}
          </button>
        </div>
      </AnyPostForm>
    </TitlePage>
  );
};

export default MainEditor;
