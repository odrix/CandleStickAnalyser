package detectdaily

import (
	"context"
	"fmt"
	"os"
	"strings"

	sendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

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
		"pair":      pattern.Pair,
		"pattern":   pattern.Type,
		"trend":     strings.ToLower(pattern.TrendDirection),
		"timeframe": "daily",
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
		fmt.Println("Error when calling TransactionalEmailsApi->post-email: ", err.Error())
		return
	}
	fmt.Println("send template 5:", email)
}

func notifyOneEmail(pattern Pattern, emailNotify string) {

	var ctx context.Context
	cfg := sendinblue.NewConfiguration()
	cfg.AddDefaultHeader("api-key", os.Getenv("SENDINBLUE_APIKEY"))

	sib := sendinblue.NewAPIClient(cfg)

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

	body.To = append(body.To, sendinblue.SendSmtpEmailTo{Email: emailNotify})
	email, _, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, body)

	if err != nil {
		fmt.Println("Error when calling TransactionalEmailsApi->smtp email: ", err.Error())
		return
	}
	fmt.Println("send template 5:", email)
}

func notifyTwitter(pattern Pattern) {

}

func notifyFacebook(pattern Pattern) {

}
