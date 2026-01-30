package ytdlp

const (
	AudioAAC    string = "aac"
	AudioALAC   string = "alac"
	AudioFLAC   string = "flac"
	AudioM4A    string = "m4a"
	AudioMP3    string = "mp3"
	AudioOPUS   string = "opus"
	AudioVORBIS string = "vorbis"
	AudioWAV    string = "wav"

	FilenameTitle         string = "%(title)s.%(ext)s"
	FilenameUploaderTitle string = "%(uploader)s_%(title)s.%(ext)s"

	FlagHelp                        string = "-h"
	FlagVersion                     string = "--version"
	FlagUpdate                      string = "-U"
	FlagIgnoreErrors                string = "-i"
	FlagNoAbortOnError              string = "--no-abort-on-error"
	FlagAbortOnError                string = "--abort-on-error"
	FlagDumpUserAgent               string = "--dump-user-agent"
	FlagListExtractors              string = "--list-extractors"
	FlagExtractorArgs               string = "--extractor-args"
	FlagExtractorOptions            string = "--extractor-descriptions"
	FlagFlatPlaylist                string = "--flat-playlist"
	FlagNoPlaylist                  string = "--no-playlist"
	FlagYesPlaylist                 string = "--yes-playlist"
	FlagSkipUnavailableFragments    string = "--skip-unavailable-fragments"
	FlagAbortOnUnavailableFragments string = "--abort-on-unavailable-fragments"
	FlagKeepFragments               string = "--keep-fragments"
	FlagLazyPlaylist                string = "--lazy-playlist"
	FlagRestrictFilenames           string = "--restrict-filenames"
	FlagWindowsFilenames            string = "--windows-filenames"
	FlagWriteThumbnail              string = "--write-thumbnail"
	FlagWriteAllThumbnails          string = "--write-all-thumbnails"
	FlagListThumbnails              string = "--list-thumbnails"
	FlagQuiet                       string = "-q"
	FlagIgnoreWarnings              string = "--no-warnings"
	FlagSimulate                    string = "-s"
	FlagNoSimulate                  string = "--no-simulate"
	FlagSkipDownload                string = "--skip-download"
	FlagDumpJson                    string = "-j"
	FlagDumpSingleJson              string = "-J"
	FlagNewLine                     string = "--newline"
	FlagNoProgress                  string = "--no-progress"
	FlagProgress                    string = "--progress"
	FlagConsoleTitle                string = "--console-title"
	FlagVerbose                     string = "-v"
	FlagExtractAudio                string = "-x"
	FlagEmbedThumbnail              string = "--embed-thumbnail"
	FlagEmbedMetadata               string = "--embed-metadata"
	FlagNoDownloadArchive           string = "--no-download-archive"
	FlagConfigLocations             string = "--config-locations"
	FlagPluginDirs                  string = "--plugin-dirs"
	FlagDownloadArchive             string = "--download-archive"
	FlagBatchFile                   string = "-a"
	FlagLoadInfoJson                string = "--load-info-json"
	FlagCookieFile                  string = "--cookies"
	FlagCacheDir                    string = "--cache-dir"
	FlagFfmpegLocation              string = "--ffmpeg-location"
	FlagProgressTemplate            string = "--progress-template"
	flagFilenameTemplate            string = "-o"
	FlagPrint                       string = "-O"
	FlagAudioQuality                string = "--audio-quality"
	flagAudioFormat                 string = "--audio-format"
	FlagVideoFormat                 string = "--recode-video"
	FlagEncoding                    string = "--encoding"
	FlagPlaylistItems               string = "-I"
)

var (
	ParamAudio    = param(flagAudioFormat)
	ParamFilename = param(flagFilenameTemplate)
)

func param(flag string) func(string) (string, string) {
	return func(value string) (string, string) {
		return flag, value
	}
}
