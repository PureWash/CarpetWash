package models

type CreateClientReq struct {
	FullName    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}

type CreateClientResp struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type CreateOrderReq struct {
	ClientID   string  `json:"client_idt"`
	Area       float32 `json:"area"`
	TotalPrice float32 `json:"total_price"`
	ServiceId  string  `json:"service_id"`
}

type CreateOrderResp struct {
	ID         string  `json:"id"`
	Area       float32 `json:"area"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
}

type UpdateClientReq struct {
	ID          string  `json:"id"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	PhoneNumber string  `json:"phone_number"`
}

type UpdateOrderReq struct {
	ID         string  `json:"id"`
	Area       float32 `json:"area"`
	TotalPrice float32 `json:"total_price"`
	Status     string  `json:"status"`
}

type UpdateOrderResp struct {
	ID         string  `json:"id"`
	ClientID   string  `json:"client_id"`
	Area       float32 `json:"area"`
	TotalPrice float32 `json:"total_price"`
	Status     string  `json:"status"`
	UpdatedAt  string  `json:"updated_at"`
}
