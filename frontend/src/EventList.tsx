import { Link } from "react-router-dom";

type Event = {
  id: number;
  date: string;
  description: string;
  eventTypeName: string;
};

const EventList = ({
  events,
  title,
  fromEvent,
}: {
  events: Event[];
  title: string;
  fromEvent: boolean;
}) => {
  return (
    <div className="Events">
      <p>{title}</p>
      {events.map((event) => {
        const eventRef = `/events/${event.eventTypeName}-${event.date}`;

        return (
          <div>
            <Link to={eventRef}>
              {fromEvent
                ? `${event.date}`
                : `${event.eventTypeName}
                ${event.date}`}
            </Link>
          </div>
        );
      })}
    </div>
  );
};

export default EventList;
