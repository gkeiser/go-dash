package main

import (
	"fmt"

	"github.com/zencoder/go-dash/mpd"
)

func liveNew() {
	m := mpd.NewMPD(mpd.DASH_PROFILE_LIVE, "PT0M12S", "PT1.97S")

	//the below two lines are not required
	// pdDur := mpd.Duration(12 * time.Second)
	// m.Periods[0].Duration = pdDur
	// audioAS, _ := m.AddNewAdaptationSetAudio(mpd.DASH_MIME_TYPE_AUDIO_MP4, true, 1, "und")
	// _, _ = audioAS.SetNewSegmentTemplate(1968, "$RepresentationID$/audio/en/init.mp4", "$RepresentationID$/audio/en/seg-$Number$.m4f", 0, 1000)
	// _, _ = audioAS.AddNewRepresentationAudio(48000, 67095, "mp4a.40.2", "800")
	// (duration / timescale = segment length in seconds)
	const timescale = int64(90000)
	duration := 2 * timescale
	videoAS, _ := m.AddNewAdaptationSetVideo(mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)
	_, _ = videoAS.SetNewSegmentTemplate(duration, "$RepresentationID$/video_init1.m4s", "$RepresentationID$/video$Number$.m4s", 1, timescale)
	_, _ = videoAS.AddNewRepresentationVideo(3569000, "avc1.4d401f", "3000k", "30000/1001", 1920, 1280)
	// _, _ = videoAS.AddNewRepresentationVideo(1911775, "avc1.4d401f", "1000", "30000/1001", 1024, 576)
	// _, _ = videoAS.AddNewRepresentationVideo(2295158, "avc1.4d401f", "1200", "30000/1001", 1024, 576)
	// _, _ = videoAS.AddNewRepresentationVideo(2780732, "avc1.4d401f", "1500", "30000/1001", 1280, 720)

	// subtitleAS, _ := m.AddNewAdaptationSetSubtitle(mpd.DASH_MIME_TYPE_SUBTITLE_VTT, "en")
	// subtitleRep, _ := subtitleAS.AddNewRepresentationSubtitle(256, "subtitle_en")
	// _ = subtitleRep.SetNewBaseURL("http://example.com/content/sintel/subtitles/subtitles_en.vtt")

	schemeIDURI := "urn:mpeg:dash:utc:http-iso:2014"
	//or, write out current server time?
	// To comply with the requirement described in paragraph 4.7.2 of version 4.0 of these guidelines, we add the UTCTiming element to the MPD when live streaming (as of release 1.7.27):
	// from https://docs.unified-streaming.com/documentation/live/recommended-settings.html
	value := "https://time.akamai.com/?iso"
	m.UTCTiming = &mpd.DescriptorType{
		SchemeIDURI: &schemeIDURI,
		Value:       &value,
	}

	mpdStr, _ := m.WriteToString()
	fmt.Println(mpdStr)
}

func main() {
	liveNew()
}
