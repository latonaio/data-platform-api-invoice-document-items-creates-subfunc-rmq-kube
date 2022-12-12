package subfunction

import (
	"context"
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc-rmq/API_Processing_Data_Formatter"
	"strings"

	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.MetaData, error) {
	var err error
	var metaData *api_processing_data_formatter.MetaData

	metaData, err = psdc.ConvertToMetaData(sdc)
	if err != nil {
		return nil, err
	}

	return metaData, nil
}

func (f *SubFunction) OrderIDByNumberSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	var args []interface{}

	billFromParty := sdc.InvoiceDocumentInputParameters.BillFromParty
	billToParty := sdc.InvoiceDocumentInputParameters.BillToParty

	if len(*billFromParty) != len(*billToParty) {
		return nil, nil
	}

	dataKey, err := psdc.ConvertToOrderIDByNumberSpecificationKey(sdc, len(*billFromParty))
	if err != nil {
		return nil, err
	}

	for i := range *billFromParty {
		dataKey.BillFromParty[i] = (*billFromParty)[i]
		dataKey.BillToParty[i] = (*billToParty)[i]
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.BillFromParty)-1) + "(?,?)"
	for i := range dataKey.BillFromParty {
		args = append(args, dataKey.BillFromParty[i], dataKey.BillToParty[i])
	}

	args = append(
		args,
		dataKey.HeaderCompleteDeliveryIsDefined,
		dataKey.HeaderDeliveryStatus,
		dataKey.HeaderBillingBlockStatus,
		dataKey.HeaderBillingStatus,
	)

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByNumberSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDByRangeSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	dataKey, err := psdc.ConvertToOrderIDByRangeSpecificationKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.BillFromPartyFrom = sdc.InvoiceDocumentInputParameters.BillFromPartyFrom
	dataKey.BillFromPartyTo = sdc.InvoiceDocumentInputParameters.BillFromPartyTo
	dataKey.BillToPartyFrom = sdc.InvoiceDocumentInputParameters.BillToPartyFrom
	dataKey.BillToPartyTo = sdc.InvoiceDocumentInputParameters.BillToPartyTo

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByRangeSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDByReferenceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	dataKey, err := psdc.ConvertToOrderIDByReferenceDocumentKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.OrderID = sdc.InvoiceDocumentInputParameters.ReferenceDocument
	dataKey.BillFromParty = append(dataKey.BillFromParty, (*sdc.InvoiceDocumentInputParameters.BillFromParty)[0])
	dataKey.BillToParty = append(dataKey.BillToParty, (*sdc.InvoiceDocumentInputParameters.BillToParty)[0])

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderBillingBlockStatus) = (?, ?, ?, ?, ?)
		AND HeaderDeliveryStatus <> ?;`, dataKey.OrderID, dataKey.BillFromParty[0], dataKey.BillToParty[0], dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderBillingBlockStatus, dataKey.HeaderDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByReferenceDocument(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderItem, error) {
	var args []interface{}

	orderID := psdc.OrderID

	dataKey, err := psdc.ConvertToOrderItemKey(sdc, len(*orderID))
	if err != nil {
		return nil, err
	}

	for i := range *orderID {
		dataKey.OrderID[i] = (*orderID)[i].OrderID
	}

	repeat := strings.Repeat("?,", len(dataKey.OrderID)-1) + "?"
	for _, tag := range dataKey.OrderID {
		args = append(args, tag)
	}

	args = append(
		args,
		dataKey.ItemCompleteDeliveryIsDefined,
		dataKey.ItemDeliveryStatus,
		dataKey.ItemBillingBlockStatus,
		dataKey.ItemBillingStatus,
	)

	rows, err := f.db.Query(
		`SELECT OrderID, ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBillingStatus, ItemBillingBlockStatus, OrderItem
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE OrderID IN ( `+repeat+` )
		AND (ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBillingBlockStatus) = (?, ?, ?)
		AND ItemBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderItem(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentByNumberSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocument, error) {
	var args []interface{}

	billFromParty := sdc.InvoiceDocumentInputParameters.BillFromParty
	billToParty := sdc.InvoiceDocumentInputParameters.BillToParty

	if len(*billFromParty) != len(*billToParty) {
		return nil, nil
	}

	dataKey, err := psdc.ConvertToDeliveryDocumentByNumberSpecificationKey(sdc, len(*billFromParty))
	if err != nil {
		return nil, err
	}

	for i := range *billFromParty {
		dataKey.BillFromParty[i] = (*billFromParty)[i]
		dataKey.BillToParty[i] = (*billToParty)[i]
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.BillFromParty)-1) + "(?,?)"
	for i := range dataKey.BillFromParty {
		args = append(args, dataKey.BillFromParty[i], dataKey.BillToParty[i])
	}

	args = append(
		args,
		dataKey.HeaderCompleteDeliveryIsDefined,
		dataKey.HeaderDeliveryStatus,
		dataKey.HeaderBillingBlockStatus,
		dataKey.HeaderBillingStatus,
	)

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentByNumberSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentByRangeSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocument, error) {
	dataKey, err := psdc.ConvertToDeliveryDocumentByRangeSpecificationKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.BillFromPartyFrom = sdc.InvoiceDocumentInputParameters.BillFromPartyFrom
	dataKey.BillFromPartyTo = sdc.InvoiceDocumentInputParameters.BillFromPartyTo
	dataKey.BillToPartyFrom = sdc.InvoiceDocumentInputParameters.BillToPartyFrom
	dataKey.BillToPartyTo = sdc.InvoiceDocumentInputParameters.BillToPartyTo

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentByRangeSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentByReferenceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocument, error) {
	dataKey, err := psdc.ConvertToDeliveryDocumentByReferenceDocumentKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.DeliveryDocument = sdc.InvoiceDocumentInputParameters.ReferenceDocument
	dataKey.BillFromParty = append(dataKey.BillFromParty, (*sdc.InvoiceDocumentInputParameters.BillFromParty)[0])
	dataKey.BillToParty = append(dataKey.BillToParty, (*sdc.InvoiceDocumentInputParameters.BillToParty)[0])

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus) = (?, ?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.DeliveryDocument, dataKey.BillFromParty[0], dataKey.BillToParty[0], dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentByReferenceDocument(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocumentItem, error) {
	var args []interface{}

	deliveryDocument := psdc.DeliveryDocument

	dataKey, err := psdc.ConvertToDeliveryDocumentItemKey(sdc, len(*deliveryDocument))
	if err != nil {
		return nil, err
	}

	for i := range *deliveryDocument {
		dataKey.DeliveryDocument[i] = (*deliveryDocument)[i].DeliveryDocument

	}

	dataKey.ConfirmedDeliveryDateFrom = sdc.InvoiceDocumentInputParameters.ConfirmedDeliveryDateFrom
	dataKey.ConfirmedDeliveryDateTo = sdc.InvoiceDocumentInputParameters.ConfirmedDeliveryDateTo
	dataKey.ActualGoodsIssueDateFrom = sdc.InvoiceDocumentInputParameters.ActualGoodsIssueDateFrom
	dataKey.ActualGoodsIssueDateTo = sdc.InvoiceDocumentInputParameters.ActualGoodsIssueDateTo

	repeat := strings.Repeat("?,", len(dataKey.DeliveryDocument)-1) + "?"
	for _, tag := range dataKey.DeliveryDocument {
		args = append(args, tag)
	}

	args = append(
		args,
		dataKey.ConfirmedDeliveryDateFrom,
		dataKey.ConfirmedDeliveryDateTo,
		dataKey.ActualGoodsIssueDateFrom,
		dataKey.ActualGoodsIssueDateTo,
		dataKey.ItemCompleteDeliveryIsDefined,
		dataKey.ItemDeliveryStatus,
		dataKey.ItemBillingBlockStatus,
		dataKey.ItemBillingStatus,
	)

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, ConfirmedDeliveryDate, ActualGoodsIssueDate, ItemCompleteDeliveryIsDefined, ItemDeliveryIncompletionStatus, ItemBillingStatus, ItemBillingBlockStatus, DeliveryDocumentItem
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		WHERE DeliveryDocument IN ( `+repeat+` )
		AND ConfirmedDeliveryDate BETWEEN ? AND ?
		AND ActualGoodsIssueDate BETWEEN ? AND ?
		AND (ItemCompleteDeliveryIsDefined, ItemDeliveryIncompletionStatus, ItemBillingBlockStatus) = (?, ?, ?)
		AND ItemBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocumentItem(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CalculateInvoiceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateInvoiceDocument, error) {
	metaData := psdc.MetaData
	dataKey, err := psdc.ConvertToCalculateInvoiceDocumentKey()
	if err != nil {
		return nil, err
	}

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateInvoiceDocumentQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	calculateInvoiceDocument := CalculateInvoiceDocument(*dataQueryGets.InvoiceDocumentLatestNumber)

	data, err := psdc.ConvertToCalculateInvoiceDocument(calculateInvoiceDocument)
	if err != nil {
		return nil, err
	}

	return data, err
}

func CalculateInvoiceDocument(latestNumber int) *int {
	res := latestNumber + 1
	return &res
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(3)

	psdc.MetaData, err = f.MetaData(sdc, psdc)
	if err != nil {
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// // I-1-1. OrderIDの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		// psdc.OrderID, e = f.OrderIDByNumberSpecification(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// I-1-1. OrderIDの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		psdc.OrderID, e = f.OrderIDByRangeSpecification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// // II-1-1. OrderIDが未請求対象であることの確認
		// psdc.OrderID, e = f.OrderIDByReferenceDocument(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// I-1-3. OrderItemの絞り込み
		psdc.OrderItem, e = f.OrderItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-I-4. オーダー参照レコード・値の取得（オーダー明細）
		psdc.ItemOrdersItem, e = f.ItemOrdersItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// // I-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		// psdc.DeliveryDocument, e = f.DeliveryDocumentByNumberSpecification(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// I-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		psdc.DeliveryDocument, e = f.DeliveryDocumentByRangeSpecification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// // II-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
		// psdc.DeliveryDocument, e = f.DeliveryDocumentByReferenceDocument(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// II-1-3. Delivery Document Itemの絞り込み
		psdc.DeliveryDocumentItem, e = f.DeliveryDocumentItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-II-4. 入出荷伝票参照レコード・値の取得（入出荷伝票明細）
		psdc.ItemDeliveryDocumentItem, e = f.ItemDeliveryDocumentItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-1. InvoiceDocument
		psdc.CalculateInvoiceDocument, e = f.CalculateInvoiceDocument(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)

	osdc, err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
