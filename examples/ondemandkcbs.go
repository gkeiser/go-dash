package main

import (
	"fmt"

	"github.com/zencoder/go-dash/mpd"
)

// func main() {
// 	exampleSyncbakOndemand()
// }

func exampleSyncbakOndemand() {
	m := mpd.NewMPD(mpd.DASH_PROFILE_ONDEMAND, "PT90S", "PT1.97S")

	// audioAS, _ := m.AddNewAdaptationSetAudio(mpd.DASH_MIME_TYPE_AUDIO_MP4, true, 1, "und")
	// _, _ = audioAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	// _, _ = audioAS.AddNewContentProtectionSchemeWidevineWithPSSH([]byte{})
	// _, _ = audioAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	// audioRep, _ := audioAS.AddNewRepresentationAudio(44100, 128558, "mp4a.40.5", "800k/audio-und")
	// _ = audioRep.SetNewBaseURL("800k/output-audio-und.mp4")
	// _, _ = audioRep.AddNewSegmentBase("629-756", "0-628")

	// videoAS, _ := m.AddNewAdaptationSetVideo(mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)
	// _, _ = videoAS.AddNewContentProtectionRoot("08e367028f33436ca5dd60ffe5571e60")
	// _, _ = videoAS.AddNewContentProtectionSchemeWidevineWithPSSH([]byte{})
	// _, _ = videoAS.AddNewContentProtectionSchemePlayready("mgIAAAEAAQCQAjwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuAG0AaQBjAHIAbwBzAG8AZgB0AC4AYwBvAG0ALwBEAFIATQAvADIAMAAwADcALwAwADMALwBQAGwAYQB5AFIAZQBhAGQAeQBIAGUAYQBkAGUAcgAiACAAdgBlAHIAcwBpAG8AbgA9ACIANAAuADAALgAwAC4AMAAiAD4APABEAEEAVABBAD4APABQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsARQBZAEwARQBOAD4AMQA2ADwALwBLAEUAWQBMAEUATgA+ADwAQQBMAEcASQBEAD4AQQBFAFMAQwBUAFIAPAAvAEEATABHAEkARAA+ADwALwBQAFIATwBUAEUAQwBUAEkATgBGAE8APgA8AEsASQBEAD4AQQBtAGYAagBDAFQATwBQAGIARQBPAGwAMwBXAEQALwA1AG0AYwBlAGMAQQA9AD0APAAvAEsASQBEAD4APABDAEgARQBDAEsAUwBVAE0APgBCAEcAdwAxAGEAWQBaADEAWQBYAE0APQA8AC8AQwBIAEUAQwBLAFMAVQBNAD4APABMAEEAXwBVAFIATAA+AGgAdAB0AHAAOgAvAC8AcABsAGEAeQByAGUAYQBkAHkALgBkAGkAcgBlAGMAdAB0AGEAcABzAC4AbgBlAHQALwBwAHIALwBzAHYAYwAvAHIAaQBnAGgAdABzAG0AYQBuAGEAZwBlAHIALgBhAHMAbQB4ADwALwBMAEEAXwBVAFIATAA+ADwALwBEAEEAVABBAD4APAAvAFcAUgBNAEgARQBBAEQARQBSAD4A")

	videoAS, _ := m.AddNewAdaptationSetVideo(mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)
	// 	the init body length 675
	// the body w/o init File:
	//  styp: size=24, Major=msdh, Minor=0, Compatible=msdh,msix
	//  sidx: size=52, version=1, flags=0x000000, refID=1, timescale=90000, earliestPTS=0,firstOffset=0, refCount=1
	//sidx range =(676+24),(676+24+52)
	videoRep1, _ := videoAS.AddNewRepresentationVideo(3731000, "avc1.4d401e", "3000k/video-1", "30000/1001", 1920, 1280)
	_ = videoRep1.SetNewBaseURL("3000k/output-video-1.mp4")
	_, _ = videoRep1.AddNewSegmentBase("700-752", "0-675")

	// videoRep2, _ := videoAS.AddNewRepresentationVideo(1633516, "avc1.4d401f", "1200k/video-1", "30000/1001", 960, 540)
	// _ = videoRep2.SetNewBaseURL("1200k/output-video-1.mp4")
	// _, _ = videoRep2.AddNewSegmentBase("686-813", "0-685")

	// subtitleAS, _ := m.AddNewAdaptationSetSubtitle(mpd.DASH_MIME_TYPE_SUBTITLE_VTT, "en")
	// subtitleRep, _ := subtitleAS.AddNewRepresentationSubtitle(256, "captions_en")
	// _ = subtitleRep.SetNewBaseURL("http://example.com/content/sintel/subtitles/subtitles_en.vtt")

	mpdStr, _ := m.WriteToString()
	fmt.Println(mpdStr)
}
