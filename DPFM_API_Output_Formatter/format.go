package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToPaymentMethod(rows *sql.Rows) (*PaymentMethod, error) {
	defer rows.Close()
	paymentMethod := PaymentMethod{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&paymentMethod.PaymentMethod,
			&paymentMethod.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentMethod, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &paymentMethod, nil
	}

	return &paymentMethod, nil
}
