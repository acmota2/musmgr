import { useSearchParams } from "react-router-dom";
import TitlePage from "../multipurpose/titlepage";
import AnyPostForm from "../multipurpose/AnyPostForm";

export type text = 0;
export type score = 1;
export type FileType = text | score;
export type SongFile = {
  Name: string;
  Type: FileType;
};

const MainEditor = () => {
  const [params] = useSearchParams();
  return (
    <TitlePage title={params.get("name") || ""}>
      <AnyPostForm
        path=""
        redirectTo={() => ""}
        buttonText="Concluir"
        state={{}}
        dataCreator={() => {}}
      >
        {}
      </AnyPostForm>
    </TitlePage>
  );
};

export default MainEditor;
