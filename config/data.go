package config

type Data struct {
	Metric    string `json:"metric"`
	Endpoint  string `json:"endpoint"`
	Tags      string `json:"tags"`
	Value     int    `json:"value"`
	Timestamp int    `json:"timestamp"`
	Step      int    `json:"step"`
}