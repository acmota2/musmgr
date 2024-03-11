import { useSearchParams } from "react-router-dom";
import SongList from "./multipurpose/SongList";
import AnyList from "./multipurpose/anylist";
import TitlePage from "./multipurpose/titlepage";
import { Song } from "./CreateSong";
import { useState } from "react";
import Modal from "./multipurpose/Modal";
import useGet from "./hooks/useGet";
import Loading from "./multipurpose/loading";
import ErrorPage from "./multipurpose/errorpage";
import AnyPostForm from "./multipurpose/AnyPostForm";

type SongEvent = {
  event_id: number;
  song_id: number;
};

export default function EventSongs() {
  const [params] = useSearchParams();
  const [modalOpen, setModalOpen] = useState<boolean>(false);
  const [curSong, setCurSong] = useState<string>(
    JSON.stringify({ id: 0, name: "" }),
  );
  const songEventState = {
    stateSongID: useState<number>(0), // not sure if not useRef
  };
  const [, setSongID] = songEventState.stateSongID;
  const { data, pending, err } = useGet<Song[]>("/songs");

  return (
    <TitlePage
      title={`Canções de ${params.get("name")} | ${params.get("date")}`}
      renderRight={1}
    >
      <button
        onClick={(e) => {
          e.preventDefault();
          setModalOpen(true);
        }}
      >
        Adicionar canção...
      </button>
      <Modal isOpened={modalOpen} onClose={() => setModalOpen(false)}>
        <AnyPostForm
          path="/song_event"
          redirectTo={() => ""}
          buttonText="Adicionar"
          state={songEventState}
          dataCreator={({ stateSongID: [song_id] }): SongEvent => ({
            song_id,
            event_id: parseInt(params.get("id")!, 10),
          })}
        >
          {err && <ErrorPage err={err} />}
          {pending && <Loading />}
          {!pending && (
            <select
              value={curSong}
              onChange={(e) => {
                setCurSong(e.target.value);
                setSongID(JSON.parse(e.target.value).id);
              }}
            >
              {data.map((s: Song) => (
                <option
                  value={JSON.stringify({ id: s.id, name: s.name })}
                  key={s.id}
                >
                  {s.name}
                </option>
              ))}
            </select>
          )}
        </AnyPostForm>
      </Modal>
      <AnyList
        path={`/songs/event/${params.get("id")}`}
        generator={(data: Song[]) => <SongList data={data} />}
      />
    </TitlePage>
  );
}
