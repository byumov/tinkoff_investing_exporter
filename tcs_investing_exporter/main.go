package main

import (
	"context"
	"fmt"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics(token string) {

	labels := []string{"currency", "name", "ticker", "type", "account"}

	tcsItemTotalPrice := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("tcs_item"),
		Help: "Total price for portfolio item"},
		labels)

	tcsExpectedYield := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("tcs_expected_yield"),
		Help: "Total expected yield for portfolio item"},
		labels)

	tcsCurrency := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("tcs_currency"),
		Help: "Currencies in rubles"},
		labels)

	tcsItemCurrentPrice := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("tcs_current_price"),
		Help: "Current item price"},
		labels)

	go func() {

		client := sdk.NewRestClient(token)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		accounts, err := client.Accounts(ctx)
		if err != nil {
			log.Fatalf("Can't get accounts: %v", err)
		}

		for _, account := range accounts {
			portfolio, err := client.Portfolio(ctx, account.ID)
			if err != nil {
				log.Fatalf("Can't get portfolio for account %v: %v", account.ID, err)
			}

			for _, item := range portfolio.Positions {

				// get current price
				orders, err := client.Orderbook(ctx, 5, item.FIGI)
				if err != nil {
					log.Fatalf("Can't get orders for %q: %v", item.Ticker, err)
				}

				labelsValues := []string{
					string(item.AveragePositionPrice.Currency),
					item.Name,	item.Ticker,
					string(item.InstrumentType),
					string(account.Type)}

				tcsItemTotalPrice.WithLabelValues(labelsValues...).Set(item.AveragePositionPrice.Value * float64(item.Balance))
				if item.InstrumentType == sdk.InstrumentTypeCurrency {
					tcsCurrency.WithLabelValues(labelsValues...).Set(item.AveragePositionPrice.Value)
				}
				tcsExpectedYield.WithLabelValues(labelsValues...).Set(item.ExpectedYield.Value)
				tcsItemCurrentPrice.WithLabelValues(labelsValues...).Set(orders.LastPrice)
			}

		}
		// update data every 2 minute
		time.Sleep(120 * time.Second)
	}()
}

func main() {

	token, exist := os.LookupEnv("TCS_TOKEN")
	if !exist {
		log.Fatal("Env 'TCS_TOKEN' must be set, exit")
	}

	recordMetrics(token)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}
