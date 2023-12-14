import { useState } from "react";
import TitlePage from "./multipurpose/titlepage";
import "./styles/CreateSong.scss";
import AnyPostForm, { FormState } from "./multipurpose/AnyPostForm";
import SubCategory from "./Category";
import SelectorButton from "./multipurpose/SelectorButton";
import AnyList from "./multipurpose/anylist";

export type Song = {
  id: number;
  name: string;
  tonality: string;
  subcategories: number[];
};

type SongState = {
  stateName: string;
  stateNote: string;
  stateTColor: string;
  stateSubCategories: Set<number>;
};

const CreateSong = () => {
  const songState: FormState<SongState> = {
    stateName: useState<string>(""),
    stateNote: useState<string>("Dó"),
    stateTColor: useState<string>("Maior"),
    stateSubCategories: useState<Set<number>>(new Set()),
  };

  const {
    stateName: [newSongName, setNewSongName],
    stateNote: [note, setNote],
    stateTColor: [tColor, setTColor],
    stateSubCategories: [subCategories, setSubCategories],
  }: FormState<SongState> = songState;

  return (
    <AnyPostForm
      path="/song"
      redirectTo={() => "/"}
      buttonText="Criar"
      state={songState}
      dataCreator={({
        stateName: [newSongName],
        stateNote: [note],
        stateTColor: [tColor],
        stateSubCategories: [subcategories],
      }: FormState<SongState>): Song => {
        return {
          id: 0,
          name: newSongName,
          tonality: `${note} ${tColor}`,
          subcategories: Array.from(subcategories),
        };
      }}
    >
      <div className="createSong">
        <TitlePage title="Nome:">
          <div className="newSongName">
            <input
              type="text"
              placeholder="Escrever aqui..."
              required
              value={newSongName}
              onChange={(e) => setNewSongName(e.target.value)}
            />
          </div>
        </TitlePage>
        <TitlePage title="Tonalidade:">
          <div className="tonality">
            <select value={note} onChange={(e) => setNote(e.target.value)}>
              <option value="Dó">Dó</option>
              <option value="Ré">Ré</option>
              <option value="Mi">Mi</option>
              <option value="Fá">Fá</option>
              <option value="Sol">Sol</option>
              <option value="Lá">Lá</option>
              <option value="Si">Si</option>
            </select>
            <select value={tColor} onChange={(e) => setTColor(e.target.value)}>
              <option value="Maior">Maior</option>
              <option value="menor">menor</option>
            </select>
          </div>
        </TitlePage>
        <TitlePage title="Subcategorias:">
          <AnyList
            path="/subcategories"
            generator={(data: SubCategory[]) => (
              <SelectorButton
                data={data}
                key={"create"}
                textSelector={(sc: SubCategory) => sc.name}
                idDefiner={(sc: SubCategory) => sc.id.toString()}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                  const id = parseInt(e.target.id, 10);
                  if (e.target.checked) {
                    setSubCategories(subCategories.add(id));
                  } else if (subCategories.delete(id)) {
                    setSubCategories(subCategories);
                  }
                  console.log(subCategories);
                }}
              />
            )}
          />
        </TitlePage>
      </div>
    </AnyPostForm>
  );
};

export default CreateSong;
