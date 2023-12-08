export interface ErrorStatusInfo {
  code: number;
  message: string;
}

export type DependencyFunction<T> = (v: T) => void;

export type GetHook<T> = {
  data: T;
  pending: boolean;
  err: ErrorStatusInfo;
};

export type PostHook<T> = {
  pending: boolean;
  err: ErrorStatusInfo;
  poster: (data: T) => void;
};
