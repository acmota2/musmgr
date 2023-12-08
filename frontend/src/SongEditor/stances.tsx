type text = 0
type score = 1
type FileType = text | score

type SongFile = {
	Path: string
	Name: string
	Open: boolean
	Type: FileType
}

const Stances = () => {
  return {}
};

export default Stances;
