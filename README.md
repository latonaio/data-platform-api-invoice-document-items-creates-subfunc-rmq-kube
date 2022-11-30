# data-platform-api-invoice-document-items-creates-subfunc-rmq-kube
data-platform-api-invoice-document-items-creates-subfunc-rmq-kube は、データ連携基盤において、請求APIサービスの明細登録補助機能を担うマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 対象APIサービス
data-platform-api-invoice-document-items-creates-subfunc-rmq-kube の 対象APIサービスは次の通りです。

*  APIサービス URL: https://xxx.xxx.io/api/API_INVOICE_DOCUMENT_SRV/creates/

## 本レポジトリ が 対応する データ
data-platform-api-invoice-document-items-creates-subfunc-rmq-kube が対応する データ は、次のものです。

* InvoiceDocumentItem（請求 - 明細データ）
* InvoiceDocumentItemPartner（請求 - 明細取引先データ）
* InvoiceDocumentItemPricingElement（請求 - 明細価格条件データ）

## Output
data-platform-api-delivery-document-items-creates-subfunc-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。以下の項目のうち、"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
XXXXXXXXXXX
```