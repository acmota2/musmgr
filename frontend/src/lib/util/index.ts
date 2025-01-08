export interface CategoryType {
  id: string;
  title: string;
}

export interface Category {
  id: string;
  name: string;
}

export interface EventType {
  id: string;
  type: string;
}

export interface Event {
  id: string;
  name: string;
  date: string;
}

export interface Song {
  id: string;
  title: string;
  key: string;
  file: string[];
}

export interface Posts<T> {
  data: T;
}
