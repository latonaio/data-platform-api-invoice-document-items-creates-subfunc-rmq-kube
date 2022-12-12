package requests

type OrderItem struct {
	InvoiceDocument               *int   `json:"InvoiceDocument"`
	OrderID                       int    `json:"OrderID"`
	ItemCompleteDeliveryIsDefined *bool  `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string `json:"ItemDeliveryStatus"`
	ItemBillingStatus             string `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        *bool  `json:"ItemBillingBlockStatus"`
	OrderItem                     int    `json:"OrderItem"`
}
