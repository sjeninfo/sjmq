package events

type Permission struct {
	Name      string            `json:"name"`
	Domain    string            `json:"domain"`
	Languages map[string]string `json:"languages"`
}

type RegisterPermission []Permission
