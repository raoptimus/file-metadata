package metadata

import (
	"encoding/json"
	"fmt"
	"math"
	"os/exec"
	"strings"
)

const binFfprobe = "ffprobe"
const codecTypeVideo = "video"
const codecTypeAudio = "audio"

type (
	body struct {
		Format  json.RawMessage
		Streams []json.RawMessage
	}
	stream struct {
		CodecType string `json:"codec_type"`
	}
)

func Metadata(videoFile string) (*VideoMetadata, error) {
	out, err := execCmd(videoFile)
	if err != nil {
		return nil, err
	}
	b := body{}
	if err = json.Unmarshal(out, &b); err != nil {
		return nil, err
	}

	m := VideoMetadata{}
	if err = json.Unmarshal(b.Format, &m); err != nil {
		return nil, err
	}

	var videoStreams []VideoStream

	for _, raw := range b.Streams {
		var s stream
		if err := json.Unmarshal(raw, &s); err != nil {
			return nil, err
		}

		switch s.CodecType {
		case codecTypeAudio:
			var as AudioStream
			if err := json.Unmarshal(raw, &as); err != nil {
				return nil, err
			}
			m.AudioStreams = append(m.AudioStreams, as)
		case codecTypeVideo:
			var vs VideoStream
			if err := unmarshalVideoStream(raw, &vs); err != nil {
				return nil, err
			}

			if vs.Duration == 0 {
				vs.Duration = m.Duration
			}

			switch {
			case strings.ToLower(vs.CodecName) == "unknown":
				continue
			case vs.Duration == 0:
				continue
			case vs.RFrames <= 0:
				continue
			}

			m.Width = vs.Width
			m.Height = vs.Height

			videoStreams = append(videoStreams, vs)
		}
	}

	for _, vs := range videoStreams {
		if vs.BitRate == 0 && len(m.AudioStreams) > 0 {
			vs.BitRate = m.BitRate - m.AudioStreams[0].BitRate
		}
		if vs.BitRate != 0 {
			m.VideoStreams = append(m.VideoStreams, vs)
		}
	}

	if m.Duration == 0 || len(m.VideoStreams) == 0 {
		return nil, fmt.Errorf("wrong metadata by file %s", videoFile)
	}

	return &m, nil
}

func unmarshalVideoStream(raw json.RawMessage, vs *VideoStream) error {
	if err := json.Unmarshal(raw, vs); err != nil {
		return err
	}

	var r1, r2 float64
	n, err := fmt.Sscanf(vs.AvgFrameRate, "%f/%f", &r1, &r2)
	if err != nil {
		return err
	}
	if n != 2 {
		return fmt.Errorf("wrong avg_frame_rate: %s", vs.AvgFrameRate)
	}
	if r2 > 0 {
		vs.RFrames = int(math.Floor(r1 / r2))
	}

	if vs.RFrames <= 0 && vs.TimeBase != "" {
		n, err := fmt.Sscan(vs.TimeBase, &vs.RFrames)
		if err != nil {
			return err
		}
		if n != 1 {
			return fmt.Errorf("wrong time_base: %s", vs.TimeBase)
		}
	}

	return nil
}

func execCmd(videoFile string) ([]byte, error) {
	// ffprobe -show_streams -show_format -v quiet -of json gotham.121.hdtv-lol.mp4
	args := []string{
		"-show_streams",
		"-show_format",
		"-v", "quiet",
		"-of", "json",
		videoFile,
	}

	cmd := exec.Command(binFfprobe, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("%v; Output: %s", err, output)
	}

	return output, err
}
