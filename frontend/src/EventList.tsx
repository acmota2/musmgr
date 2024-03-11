import TileList from "./multipurpose/tilelist";
import TitlePage from "./multipurpose/titlepage";
import { useSearchParams } from "react-router-dom";
import AnyList from "./multipurpose/anylist";
import { useState } from "react";
import CreateNode from "./multipurpose/CreateNode";

type Event = {
  id: number;
  date: string;
  description: string;
  eventTypeName: string;
};

type DateSelectorElement = {
  text: string;
  className: string;
  min?: number;
  size: number;
  current: number;
  onChange: (n: number) => void;
};

const EventFormatter = ({ data: event }: { data: Event }) => (
  <p>{event.date}</p>
);

const DateSelector = ({
  text,
  className,
  min = 1,
  size,
  current,
  onChange,
}: DateSelectorElement) => {
  return (
    <div className="dateElement">
      <p>{text}</p>
      <select
        key={className}
        value={current}
        className={className}
        onChange={(e) => onChange(parseInt(e.target.value, 10))}
      >
        {[...Array(size).keys()].map((i: number) => {
          i += min;
          return (
            <option value={`${i}`} key={`${className}_${i}`}>{`${i}`}</option>
          );
        })}
      </select>
    </div>
  );
};

const EventList = () => {
  const [params] = useSearchParams();
  const currentDate = new Date();
  const eventState = {
    stateYear: useState<number>(currentDate.getFullYear()),
    stateMonth: useState<number>(currentDate.getMonth()),
    stateDay: useState<number>(currentDate.getDay()),
    stateDescription: useState<string>(""),
  };
  const [year, setYear] = eventState.stateYear;
  const [month, setMonth] = eventState.stateMonth;
  const [day, setDay] = eventState.stateDay;
  const [description, setDescription] = eventState.stateDescription;

  return (
    <TitlePage title={`Eventos de ${params.get("name")}`} renderRight={1}>
      <CreateNode
        path={"/event"}
        redirectTo={() => "0"}
        buttonText="Criar evento..."
        state={eventState}
        dataCreator={({
          stateYear: [year],
          stateMonth: [month],
          stateDay: [day],
          stateDescription: [description],
        }) => {
          const padding = (n: number) => (n < 10 ? `0${n}` : `${n}`);
          return {
            id: 0,
            date: `${year}-${padding(month)}-${padding(day)}`,
            description,
            event_type_name: params.get("name"),
          };
        }}
      >
        <div className="date" style={{ display: "flex", flexDirection: "row" }}>
          <DateSelector
            text="Ano"
            className="year"
            min={1970}
            size={year - 1970 + 1}
            current={year}
            onChange={setYear}
          />
          <DateSelector
            text="Mês"
            className="month"
            size={12}
            current={month}
            onChange={setMonth}
          />
          <DateSelector
            text="Dia"
            className="day"
            size={31}
            current={day}
            onChange={setDay}
          />
        </div>
        <input
          autoFocus
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="Descrição do evento..."
        />
      </CreateNode>
      <AnyList
        path={`/events/event_type/${params.get("name")}`}
        generator={(data: Event[]) => (
          <TileList
            dataList={data}
            linkMaker={(event) =>
              `/event-songs?name=${params.get("name")}&date=${event.date}&id=${event.id}`
            }
            linkFormatter={EventFormatter}
          />
        )}
      />
    </TitlePage>
  );
};

export default EventList;
