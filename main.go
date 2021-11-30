package main

import (
	"context"
	"fmt"
	"os"

	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	sendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

//const baseUrl string = "https://api.binance.com"

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	pair := "BTCUSDT"
	kl := getKline(logger, pair, binance.Day, 50)

	var candlesDesc []Candle
	for i := len(kl) - 1; i > 0; i-- {
		candlesDesc = append(candlesDesc, Candle{kl[i]})
	}

	for i := 0; i < len(candlesDesc); i++ {
		p := Pattern{}
		if IsMorningStar(candlesDesc[i:]) {
			p = Pattern{
				Pair:  pair,
				Type:  "Morning Star",
				Start: candlesDesc[i-2].OpenTime,
				End:   candlesDesc[i].CloseTime,
			}
		} else if IsEveningStar(candlesDesc[i:]) {
			p = Pattern{
				Pair:  pair,
				Type:  "Evening Star",
				Start: candlesDesc[i-2].OpenTime,
				End:   candlesDesc[i].CloseTime,
			}
		} else if IsThreeWhiteSoldiers(candlesDesc[i:]) {
			p = Pattern{
				Pair:  pair,
				Type:  "Three White Soldiers",
				Start: candlesDesc[i-2].OpenTime,
				End:   candlesDesc[i].CloseTime,
			}
		} else if IsThreeBlackCrows(candlesDesc[i:]) {
			p = Pattern{
				Pair:  pair,
				Type:  "Three Black Crows",
				Start: candlesDesc[i-2].OpenTime,
				End:   candlesDesc[i].CloseTime,
			}
		} else if IsWhiteMarubozu(candlesDesc[i:]) {
			p = Pattern{
				Pair:  pair,
				Type:  "White Marubozu",
				Start: candlesDesc[i].OpenTime,
				End:   candlesDesc[i].CloseTime,
			}
		} else if IsBlackMarubozu(candlesDesc[i:]) {
			p = Pattern{
				Pair:  pair,
				Type:  "Black Marubozu",
				Start: candlesDesc[i].OpenTime,
				End:   candlesDesc[i].CloseTime,
			}
		}

		if p.Type != "" {
			trace(p)
			notifyContacts(p)
		}
	}
	fmt.Printf("end.")
}

func trace(pattern Pattern) {
	fmt.Printf("%s on %s \n", pattern.Type, pattern.Start)
}

func notifyContacts(pattern Pattern) {

	var ctx context.Context
	cfg := sendinblue.NewConfiguration()
	cfg.AddDefaultHeader("api-key", os.Getenv("SENDINBLUE_APIKEY"))

	var contactsQuery = &sendinblue.ContactsApiGetContactsOpts{}

	sib := sendinblue.NewAPIClient(cfg)
	AllContacts, _, errContacts := sib.ContactsApi.GetContacts(ctx, contactsQuery)
	if errContacts != nil {
		fmt.Println("Error when calling get_contacts: ", errContacts.Error())
		return
	}

	var templateParams interface{}
	templateParams = map[string]interface{}{
		"pair":    pattern.Pair,
		"pattern": pattern.Type,
	}

	body := sendinblue.SendSmtpEmail{
		To:         []sendinblue.SendSmtpEmailTo{},
		Headers:    nil,
		TemplateId: 5,
		Params:     &templateParams,
	}

	for i := 0; i < len(AllContacts.Contacts); i++ {
		body.To = append(body.To, sendinblue.SendSmtpEmailTo{Email: AllContacts.Contacts[i].Email})
	}

	email, _, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, body)

	if err != nil {
		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
		return
	}
	fmt.Println("send template 5:", email)
}

func notifyTwitter(pattern Pattern) {

}

func notifyFacebook(pattern Pattern) {

}

func getKline(logger log.Logger, symbol string, interval binance.Interval, limit int) []*binance.Kline {
	var ctx, _ = context.WithCancel(context.Background())

	binanceService := binance.NewAPIService(
		"https://www.binance.com",
		"", //os.Getenv("BINANCE_APIKEY"),
		nil,
		logger,
		ctx,
	)
	b := binance.NewBinance(binanceService)

	kl, err := b.Klines(binance.KlinesRequest{
		Symbol:   symbol,
		Interval: interval,
		Limit:    limit,
	})
	if err != nil {
		panic(err)
	}
	return kl
}
