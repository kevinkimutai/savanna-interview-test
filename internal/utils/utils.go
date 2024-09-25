package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
)

func ConvertStringToInt64(str string) (int64, error) {
	// Convert string to int64
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {

		return num, err
	}

	return num, nil
}

func ConvertInt64ToString(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

func ConvertNumericToFloat64(numeric pgtype.Numeric) float64 {
	fval, _ := numeric.Value()

	//Convert To Float64
	var floatVal float64
	if strVal, ok := fval.(string); ok {
		floatVal, _ = strconv.ParseFloat(strVal, 64)
	}
	return floatVal

}

var LIMIT, OFFSET int32 = 10, 0

func GetProductAPIParams(params domain.ProductParams) queries.ListProductsParams {
	var searchStr pgtype.Text
	var start_price, end_price pgtype.Numeric

	//Get Params
	if params.Search != "" {
		searchStr.Scan(params.Search)
	}

	if params.PriceStart != "" {
		//convert to float64 first
		price, _ := strconv.ParseFloat(params.PriceStart, 64)
		start_price = ConvertFloat64ToNumeric(price)
	}

	if params.PriceEnd != "" {
		//convert to float64 first
		price, _ := strconv.ParseFloat(params.PriceEnd, 64)
		end_price = ConvertFloat64ToNumeric(price)
	}

	if params.Limit != "" {
		items, _ := strconv.Atoi(params.Limit)

		LIMIT = int32(items)

	}
	if params.Page != "" {
		page, _ := strconv.Atoi(params.Page)

		if page < 1 {
			page = 1
		}

		OFFSET = (int32(page) - 1) * LIMIT

	}

	return queries.ListProductsParams{
		Column1: searchStr,
		Price:   start_price,
		Price_2: end_price,
		Limit:   LIMIT,
		Offset:  OFFSET,
	}
}
func GetOrderAPIParams(params domain.OrderParams) queries.ListOrdersParams {
	var searchStr pgtype.Text
	var release_date_start, release_date_end pgtype.Date
	var start_price, end_price pgtype.Numeric

	//Get Params
	if params.Search != "" {
		searchStr.Scan(params.Search)
	}

	if params.TotalAmountStart != "" {
		//convert to float64 first
		price, _ := strconv.ParseFloat(params.TotalAmountStart, 64)
		start_price.Scan(price)
	}

	if params.TotalAmountEnd != "" {
		//convert to float64 first
		price, _ := strconv.ParseFloat(params.TotalAmountEnd, 64)
		start_price.Scan(price)
	}

	if params.StartDate != "" {
		//convert to type time
		date, _ := time.Parse("2006-01-02", params.StartDate)
		release_date_start.Scan(date)
	}

	if params.EndDate != "" {
		//convert to type time
		date, _ := time.Parse("2006-01-02", params.EndDate)
		release_date_end.Scan(date)
	}

	if params.Limit != "" {
		items, _ := strconv.Atoi(params.Limit)

		LIMIT = int32(items)

	}
	if params.Page != "" {
		page, _ := strconv.Atoi(params.Page)

		if page < 1 {
			page = 1
		}

		OFFSET = (int32(page) - 1) * LIMIT

	}

	return queries.ListOrdersParams{
		Column1:       searchStr,
		TotalAmount:   start_price,
		TotalAmount_2: end_price,
		CreatedAt:     pgtype.Timestamptz(release_date_start),
		CreatedAt_2:   pgtype.Timestamptz(release_date_end),
		Limit:         LIMIT,
		Offset:        OFFSET,
	}
}

func GetPage(offset, limit int32) uint {
	return uint((offset / limit) + 1)
}

func ConvertFloat64ToNumeric(f float64) pgtype.Numeric {
	// Convert float64 to string with desired precision (e.g., 2 decimal places)
	strVal := strconv.FormatFloat(f, 'f', -2, 64)

	// Create a new pgtype.Numeric object
	numeric := pgtype.Numeric{}

	// Scan the string representation using pgtype.Scan
	err := numeric.Scan(strVal)
	if err != nil {
		fmt.Println(err)
	}

	return numeric
}
