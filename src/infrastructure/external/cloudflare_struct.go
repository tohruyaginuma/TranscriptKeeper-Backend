package external

type WhisperOutput struct {
	TranscriptionInfo TranscriptionInfo `json:"transcription_info"`
	Text              string            `json:"text"`
	WordCount         int               `json:"word_count"`
	Segments          []Segment         `json:"segments"`
	Vtt               string            `json:"vtt"`
}

type TranscriptionInfo struct {
	Language            string  `json:"language"`
	LanguageProbability float64 `json:"language_probability"`
	Duration            float64 `json:"duration"`
	DurationAfterVAD    float64 `json:"duration_after_vad"`
}

type Segment struct {
	Items []SegmentItem `json:"items"`
}

type SegmentItem struct {
	Start            int     `json:"start"`
	End              int     `json:"end"`
	Text             string  `json:"text"`
	Temperature      float64 `json:"temperature"`
	AvgLogprob       float64 `json:"avg_logprob"`
	CompressionRatio float64 `json:"compression_ratio"`
	NoSpeechProb     float64 `json:"no_speech_prob"`
	Words            []Word  `json:"words"`
}

type Word struct {
	Items []WordItem `json:"items"`
}

type WordItem struct {
	Word  string `json:"word"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type WhisperInput struct {
	Audio                         string  `json:"audio"`
	Task                          string  `json:"task"`
	Language                      string  `json:"language,omitempty"`
	VadFilter                     bool    `json:"vad_filter,omitempty"`
	InitialPrompt                 string  `json:"initial_prompt,omitempty"`
	Prefix                        string  `json:"prefix,omitempty"`
	BeamSize                      int     `json:"beam_size,omitempty"`
	ConditionOnPrevious           bool    `json:"condition_on_previous,omitempty"`
	NoSpeechThreshold             float64 `json:"no_speech_threshold,omitempty"`
	CompressionRatioThreshold     float64 `json:"compression_ratio_threshold,omitempty"`
	LogProbThreshold              float64 `json:"log_prob_threshold,omitempty"`
	HallucinationSilenceThreshold float64 `json:"hallucination_silence_threshold,omitempty"`
}
