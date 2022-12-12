package requests

type DeliveryDocumentItem struct {
	InvoiceDocument               *int    `json:"InvoiceDocument"`
	DeliveryDocument              int     `json:"DeliveryDocument"`
	ConfirmedDeliveryDateFrom     *string `json:"ConfirmedDeliveryDateFrom"`
	ConfirmedDeliveryDateTo       *string `json:"ConfirmedDeliveryDateTo"`
	ActualGoodsIssueDateFrom      *string `json:"ActualGoodsIssueDateFrom"`
	ActualGoodsIssueDateTo        *string `json:"ActualGoodsIssueDateTo"`
	ConfirmedDeliveryDate         *string `json:"ConfirmedDeliveryDate"`
	ActualGoodsIssueDate          *string `json:"ActualGoodsIssueDate"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string  `json:"ItemDeliveryStatus"`
	ItemBillingStatus             string  `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        *bool   `json:"ItemBillingBlockStatus"`
	DeliveryDocumentItem          int     `json:"DeliveryDocumentItem"`
}
