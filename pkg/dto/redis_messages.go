package dto

type PayloadMessageRedisDto struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type RedisMessages struct {
	GroupID string `json:"group_id"`
	Payload PayloadMessageRedisDto
}
