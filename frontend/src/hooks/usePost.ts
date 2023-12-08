import { useState } from "react";
import { ErrorStatusInfo, PostHook } from "./hooktypes";
import { To, useNavigate } from "react-router-dom";

function usePost<T>(path: string, redirectTo: To | null): PostHook<T> {
  const origin = import.meta.env.VITE_BACKEND_URL;
  const [pending, setPending] = useState<boolean>(false);
  const [err, setErr] = useState<ErrorStatusInfo>({} as ErrorStatusInfo);
  const navigate = useNavigate();

  const poster = <T>(data: T) => {
    const requestOptions = {
      method: "POST",
      headers: { "Access-Control-Allow-Origin": origin },
      body: JSON.stringify(data),
    };
    setPending(true);
    setTimeout(() => {
      console.log(data, `Future path: ${path}`);
      setPending(false);
      if (redirectTo != null) {
        navigate(redirectTo);
      }
    }, 1000);
    // fetch(origin + path, requestOptions)
    //   .then((res) => {
    //     if (!res.ok) {
    //       setPending(false);
    //       const message = "POST error";
    //       setErr({ code: res.status, message });
    //     }
    //     setPending(true);
    //     return res.json;
    //   })
    //   .then(() => {
    //     setPending(false);
    //     setErr({} as ErrorStatusInfo);
    //   });
  };

  return { pending, err, poster };
}

export default usePost;
