import TileList from "./multipurpose/tilelist";

type Event = {
  id: number;
  date: string;
  description: string;
  eventTypeName: string;
};

const EventFormatter = ({ data: event }: { data: Event }) => (
  <p>{event.date}</p>
);

const EventList = ({
  events,
}: {
  events: Event[];
  title: string;
  fromEvent: boolean;
}) => {
  return (
    <div className="Events">
      <TileList
        dataList={events}
        linkMaker={(event) => event.id.toString()}
        linkFormatter={EventFormatter}
      />
    </div>
  );
};

export default EventList;
