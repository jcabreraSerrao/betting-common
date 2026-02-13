package redis

// Groups representa el modelo de grupo para Redis, solo con ID y Name
type Groups struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	ChatInstanceID string `json:"chat_instance_id"`
}
