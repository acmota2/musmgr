import "../styles/SelectorButton.scss";

interface SelectorButtonProps<T> {
  data: T[];
  key: string,
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  textSelector: TextSelector<T>;
  idDefiner: IdDefiner<T>;
}

type TextSelector<T> = (t: T) => string;
type IdDefiner<T> = (t: T) => string;

const SelectorButton = <T,>({
  data,
  key,
  onChange,
  textSelector,
  idDefiner,
}: SelectorButtonProps<T>) => {
  return (
    <ul key={key}>
      {data.map((t: T) => {
        const text = textSelector(t);
        const id = idDefiner(t);
        return (
          <li key={`link_${id}`} id={`${id}_li`}>
            <input type="checkbox" id={id} onChange={onChange} />
            <label htmlFor={id}>{text}</label>
          </li>
        );
      })}
    </ul>
  );
};

export default SelectorButton;
