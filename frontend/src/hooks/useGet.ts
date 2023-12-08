import { useEffect, useState } from "react";
import { DependencyFunction, ErrorStatusInfo, GetHook } from "../hooks/hooktypes";

function useGet<T>(
  path: string,
  dependencies?: DependencyFunction<T>,
): GetHook<T> {
  const origin = import.meta.env.VITE_BACKEND_URL;
  const [data, setData] = useState<T>(null as T);
  const [pending, setPending] = useState<boolean>(true);
  const [err, setErr] = useState<ErrorStatusInfo>({} as ErrorStatusInfo);

  useEffect(() => {
    const abortCtrl = new AbortController();
    const mode = "cors";
    const headers = { "Access-Control-Allow-Origin": origin };

    fetch(origin + path, { signal: abortCtrl.signal, mode, headers })
      .then((res) => {
        if (!res.ok) {
          const message = "GET error";
          setErr({ code: res.status, message });
          throw Error(message);
        }
        return res.json();
      })
      .then((data) => {
        setData(data.data);
        setPending(false);
        setErr({} as ErrorStatusInfo);
        if (dependencies !== undefined) {
          dependencies(data.data);
        }
      })
      .catch((err) => {
        if (err.name !== "AbortError") {
          setPending(false);
        }
      });

    return () => abortCtrl.abort();
  }, [path, origin]);

  return { data, pending, err };
}

export default useGet;
