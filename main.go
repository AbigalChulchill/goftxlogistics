<<<<<<< HEAD
package main

import (
	"github.com/sdcoffey/techan"
)

func main() {
	RestActions()
	WebsocketActions()
	techan.NewEMAIndicator()

}
=======
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

func main() {
	RestActions()
	WebsocketActions()
	series := techan.NewTimeSeries()

	// fetch this from your preferred exchange
	dataset := [][]string{
		// Timestamp, Open, Close, High, Low, volume
		{"1234567", "1", "2", "3", "5", "6"},
	}

	for _, datum := range dataset {
		start, _ := strconv.ParseInt(datum[0], 10, 64)
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(datum[1])
		candle.ClosePrice = big.NewFromString(datum[2])
		candle.MaxPrice = big.NewFromString(datum[3])
		candle.MinPrice = big.NewFromString(datum[4])

		series.AddCandle(candle)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewEMAIndicator(closePrices, 10) // Create an exponential moving average with a window of 10

	fmt.Println(movingAverage.Calculate(0).FormattedString(2))

}
>>>>>>> 24392a3 (commit for usage on laptop)
