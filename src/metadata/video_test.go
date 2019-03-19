package metadata

import (
	"fmt"
	"testing"
)

func TestVideoMetadata(t *testing.T) {
	data, err := Metadata("testdata/test.mp4")
	if err != nil {
		t.Fatal(err)
	}

	expectedData := VideoMetadata{
		Size:         662895,
		TotalStreams: 2,
		FileName:     "testdata/test.mp4",
		Duration:     30.001,
		Format:       "mov,mp4,m4a,3gp,3g2,mj2",
		Width:        240,
		Height:       180,
		BitRate:      176766,
		VideoStreams: []VideoStream{
			{
				Index:              0,
				Width:              240,
				Height:             180,
				BFrames:            0,
				BitRate:            109091,
				Duration:           29.965,
				SampleAspectRatio:  "1:1",
				DisplayAspectRatio: "4:3",
				RFrames:            23,
				TotalFrames:        719,
				Tags: map[string]string{
					"creation_time": "2010-02-09T01:55:39.000000Z",
					"language":      "eng",
					"handler_name":  "VideoHandler",
				},
				CodecName:    "h264",
				AvgFrameRate: "103536000/4314821",
				TimeBase:     "1/288000",
			},
		},
		AudioStreams: []AudioStream{
			{
				Index:      1,
				BitRate:    64131,
				SampleRate: 22050,
				Channels:   2,
				CodecName:  "aac",
				Tags: map[string]string{
					"creation_time": "2010-02-09T01:55:39.000000Z",
					"language":      "eng",
					"handler_name":  "SoundHandler",
				},
			},
		},
		Tags: map[string]string{
			"compatible_brands": "isomiso2avc1mp41",
			"creation_time":     "2010-02-09T01:55:39.000000Z",
			"encoder":           "Lavf57.71.100",
			"major_brand":       "isom",
			"minor_version":     "512",
		},
	}

	assertMapEqual(expectedData.Tags, data.Tags, t)
	expectedData.Tags = make(map[string]string)
	data.Tags = make(map[string]string)

	if len(data.VideoStreams) > 0 {
		assertMapEqual(expectedData.VideoStreams[0].Tags, data.VideoStreams[0].Tags, t)
		expectedData.VideoStreams[0].Tags = make(map[string]string)
		data.VideoStreams[0].Tags = make(map[string]string)
	}
	if len(data.AudioStreams) > 0 {
		assertMapEqual(expectedData.AudioStreams[0].Tags, data.AudioStreams[0].Tags, t)
		expectedData.AudioStreams[0].Tags = make(map[string]string)
		data.AudioStreams[0].Tags = make(map[string]string)
	}

	if fmt.Sprintf("%v", expectedData) != fmt.Sprintf("%v", *data) {
		t.Fatalf("expected metadata `%v`, actual `%v`", expectedData, *data)
	}
}

func assertMapEqual(expected, actual map[string]string, t *testing.T) {
	if len(expected) != len(actual) {
		t.Fatalf("expected map `%v`, actual `%v`", expected, actual)
	}
	for k, v := range expected {
		if actual[k] != v {
			t.Fatalf("expected map `%v`, actual `%v`", expected, actual)
		}
	}
}
