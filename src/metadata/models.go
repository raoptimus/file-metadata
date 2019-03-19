package metadata

type (
	VideoMetadata struct {
		Width        int               `json:"width"`
		Height       int               `json:"height"`
		FileName     string            `json:"filename"`
		TotalStreams int               `json:"nb_streams"`
		Duration     float64           `json:"duration,string"`
		Size         int64             `json:"size,string"`
		BitRate      float64           `json:"bit_rate,string"`
		VideoStreams []VideoStream     `json:"videoStreams"`
		AudioStreams []AudioStream     `json:"audioStreams"`
		Format       string            `json:"format_name"`
		Tags         map[string]string `json:"tags"`
	}

	VideoStream struct {
		Index              int               `json:"index"`
		Width              int               `json:"width"`
		Height             int               `json:"height"`
		BFrames            int               `json:"b_frames"`
		BitRate            float64           `json:"bit_rate,string"`
		Duration           float64           `json:"duration,string"`
		SampleAspectRatio  string            `json:"sample_aspect_ratio"`
		DisplayAspectRatio string            `json:"display_aspect_ratio"`
		RFrames            int               `json:"r_frames"`
		TotalFrames        int               `json:"nb_frames,string"`
		Tags               map[string]string `json:"tags"`
		CodecName          string            `json:"codec_name"`
		AvgFrameRate       string            `json:"avg_frame_rate"`
		TimeBase           string            `json:"time_base"`
	}

	AudioStream struct {
		Index      int               `json:"index"`
		BitRate    float64           `json:"bit_rate,string"`
		SampleRate float64           `json:"sample_rate,string"`
		Channels   int               `json:"channels"`
		CodecName  string            `json:"codec_name"`
		Tags       map[string]string `json:"tags"`
	}
)
