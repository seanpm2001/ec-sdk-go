package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

func main() {

	//Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve the Managed Rule for")
	managedRuleID := flag.String("managed-rule-id", "", "Manged Rule ID you wish to retrieve for the provided account number")

	flag.Parse()

	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("Error creating WAF Service: %v\n", err)
		return
	}

	//Get Managed Rule Example
	managedRule, err := wafService.GetManagedRule(*accountNumber, *managedRuleID)

	if err != nil {
		fmt.Printf("Error retrieving managed rule: %v\n", err)
		return
	}

	fmt.Println(managedRule)
}