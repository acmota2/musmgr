import { useEffect, useState } from "react";

export type ErrorInfo = { code: number; message: string };
type GetHook<T> = { data: T; pending: boolean; err: ErrorInfo };

export default function useGet<T>(url: string): GetHook<T> {
  const [data, setData] = useState<T>(null as T);
  const [pending, setPending] = useState<boolean>(true);
  const [err, setErr] = useState<ErrorInfo>({} as ErrorInfo);

  useEffect(() => {
    const abortCtrl = new AbortController();

    fetch(url, { signal: abortCtrl.signal })
      .then((res) => {
        if (!res.ok) {
          const message = "GET error";
          console.log(message);
          setErr({ code: res.status, message });
          throw Error(message); // melhorar isto, eventualmente
        }
        return res.json();
      })
      .then((data) => {
        setData(data);
        setPending(false);
      })
      .catch((err) => {
        if (err.name !== "AbortError") {
          setPending(false);
        }
      });

    return () => abortCtrl.abort();
  }, [url]);

  return { data, pending, err };
}
