package cartdto

type CartRequest struct {
	UserID    int   ` json:"user_id" `
	ProductID int   ` json:"product_id" `
	TopingID  []int `json:"toping_id" `
	Qty       int   ` json:"qty" `
	SubAmount int   ` json:"sub_amount"`
}
