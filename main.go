package main

//oapi-codegen --old-config-style -package=main -generate=client,types -o ./ugclient.gen.go :url:

import (
	"context"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func main() {
	keyId := "KEY_ID"
	secretKey := "SECRET_KEY"

	basicAuthProvider, basicAuthProviderErr := securityprovider.NewSecurityProviderBasicAuth(keyId, secretKey)
	if basicAuthProviderErr != nil {
		panic(basicAuthProviderErr)
	}

	c, err := NewClientWithResponses("https://staging.services.api.unity.com/ugc/", WithRequestEditorFn(basicAuthProvider.Intercept))
	if err != nil {
		panic(err)
	}

	params := S2SSearchContentParams{}

	projectId := "PROJECT_ID"
	environmentId := "ENVIRONMENT_ID"

	resp, err := c.S2SSearchContentWithResponse(context.Background(), projectId, environmentId, &params)
	if err != nil {
		panic(err)
	}

	contents := *resp.JSON200

	if len(contents.Results) == 0 {
		fmt.Println(resp.StatusCode())
		fmt.Println(resp.Body)
		os.Exit(1)
	}

	fmt.Println(*&contents.Results[0].Id)
	fmt.Println(contents.Results[0].Name)
	os.Exit(0)
}
