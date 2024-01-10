import "../styles/SongEditor.scss";
import { useSearchParams } from "react-router-dom";
import TitlePage from "../multipurpose/titlepage";
import AnyPostForm from "../multipurpose/AnyPostForm";
import { SongStance, SongText, emptySong } from "./song_text/structure";
import useGet from "../hooks/useGet";
import ErrorPage from "../multipurpose/errorpage";
import Loading from "../multipurpose/loading";
import Stances from "./stances";
import { useState } from "react";
import { IntervalState, stringToNoteName } from "./song_text/songState";
import NoteSelector from "./NoteSelector";

export type text = 0;
export type score = 1;
export type FileType = text | score;
export type SongFile = {
  Name: string;
  Type: FileType;
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

  const songState = useState<SongText>(JSON.parse(JSON.stringify(data)));
  const [song] = songState;

  const tonality = params.get("tonality");
  const curInterval = {
    noteName: useState<string>(tonality ? tonality : "Dó"),
    semitones: useState<number>(0),
  };

  return (
    <TitlePage title={params.get("name") || ""}>
      <AnyPostForm
        path="/file/text"
        redirectTo={() => "/songs"}
        buttonText="Concluir"
      >
        {err && <ErrorPage err={err} />}
        {pending && <Loading />}
        <div className="showSong">
          <Interval interval={curInterval} />
          {song.stances.map((s: SongStance, i) => (
            <Stances songState={songState} stance={s} stanceIndex={i} />
          ))}
        </div>
      </AnyPostForm>
    </TitlePage>
  );
};

export default MainEditor;
