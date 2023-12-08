import { useSearchParams } from "react-router-dom";
import TitlePage from "../multipurpose/titlepage";

const MainEditor = () => {
  const [params] = useSearchParams();
  return <TitlePage title={params.get("name") || ""}>{}</TitlePage>;
};

export default MainEditor;
