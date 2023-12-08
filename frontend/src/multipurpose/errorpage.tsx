import { ErrorInfo } from "../hooks/useGet";

const messageHandler = (status: number): string => {
  switch (status) {
    case 404:
      return "Tens a certeza que procuraste na barra certa?";
    case 400:
      return "Esse autocorrect...";
    case 500:
      return "...";
    default:
      return "";
  }
};

const ErrorPage = ({ err }: { err: ErrorInfo }) => {
  return (
    <div className="error">
      <div>{err.code}</div>
      <p>{messageHandler(err.code)}</p>
    </div>
  );
};

export default ErrorPage;
