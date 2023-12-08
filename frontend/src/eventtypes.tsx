import AnyList from "./multipurpose/anylist";
import TileList from "./multipurpose/tilelist";
import TitlePage from "./multipurpose/titlepage";

export type EventType = {
  name: string;
};

const EventTypeFormatter = ({ data: event }: { data: EventType }) => (
  <p>{event.name}</p>
);

export const EventTypes = () => (
  <TitlePage title="Tipos de eventos">
    <AnyList
      path="/event_types"
      generator={(data: EventType[]) => (
        <TileList
          dataList={data}
          linkMaker={(et: EventType) => et.name}
          linkFormatter={EventTypeFormatter}
        />
      )}
    />
  </TitlePage>
);

export default EventTypes;
