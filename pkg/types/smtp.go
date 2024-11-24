package types

type CreateSmtpDto struct {
	Name     string `json:"name" binding:"required"`
	Slug     string `json:"slug" binding:"required"`
	Server   string `json:"server" binding:"required"`
	Port     uint16 `json:"port" binding:"required"`
	Email    string `json:"email" binding:"required"`
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
	Default  bool   `json:"default"`
}

type UpdateSmtpDto struct {
	Name     *string `json:"name" binding:"required"`
	Slug     *string `json:"slug" binding:"required"`
	Server   *string `json:"server" binding:"required"`
	Port     *uint16 `json:"port" binding:"required"`
	Email    *string `json:"email" binding:"required"`
	User     *string `json:"user" binding:"required"`
	Password *string `json:"password" binding:"required"`
	Default  *bool   `json:"default"`
}

type SmtpDto struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Server  string `json:"server"`
	Port    uint16 `json:"port"`
	Email   string `json:"email"`
	User    string `json:"user"`
	Default bool   `json:"default"`
}
