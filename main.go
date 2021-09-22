package main

import (
	"fmt"
	"log"

	"github.com/praetorian-inc/mithril/auditors"
	"github.com/praetorian-inc/mithril/pkg/types"

	_ "github.com/praetorian-inc/mithril/auditors/auth"
)

func main() {
	auditors, err := auditors.New(types.Config{})
	if err != nil {
		log.Fatalf("failed to initialize auditors: %s", err)
	}

	var ctx types.IstioContext

	var results []types.AuditResult
	for _, auditor := range auditors {
		res, err := auditor.Audit(ctx)
		if err != nil {
			log.Printf("%s failed to run: %s", auditor.Name(), err)
		}
		results = append(results, res...)
	}

	for _, res := range results {
		fmt.Printf("%s: %s\n", res.Name, res.Description)
	}
}