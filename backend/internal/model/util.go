package model

func (ft MusmgrFileType) Valid() bool {
	switch ft {
	case
		MusmgrFileTypeScoreFull,
		MusmgrFileTypeScorePart,
		MusmgrFileTypeAudioPreview,
		MusmgrFileTypeAudioRecording,
		MusmgrFileTypePicture,
		MusmgrFileTypeVideo,
		MusmgrFileTypeOther:
		return true
	default:
	}
	return false
}
