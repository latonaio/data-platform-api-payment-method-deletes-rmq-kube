package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-payment-method-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-payment-method-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) PaymentMethodRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.PaymentMethod {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE paymentMethod.PaymentMethod = \"%s\" ", input.PaymentMethod.PaymentMethod),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	paymentMethod.PaymentMethod,
    	paymentMethod.BaseDate,
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_method_payment_method_data as paymentMethod 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentMethod(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
