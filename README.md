# ugc-golang-client

UGC Golang client

## How to get the openapi yaml URL

1. Go on this url <https://staging.services.docs.unity.com/ugc-content-admin/v1>
2. Right click on the Download button and copy the target URL

## Generate Golang client

Based on the code generator OAPI-Codegen: <https://github.com/deepmap/oapi-codegen>

```go
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

Execute this command and replace the :url: with the base url

```go
oapi-codegen --old-config-style -package=main -generate=client,types -o ./ugclient.gen.go :url:
```

## Implement your client

Create your service account by following the steps here: <https://services.docs.unity.com/docs/admin-auth>

Once your service account is created, add the UGC organization role `User Generated Content Admin`

In your code, set your authentication provider follow

```go
keyId := "KEY_ID" // Service account key id
secretKey := "SECRET_KEY" // Service account secret key

basicAuthProvider, basicAuthProviderErr := securityprovider.NewSecurityProviderBasicAuth(keyId, secretKey)
```

Then finally initialize your client:

```go
c, err := NewClientWithResponses("https://staging.services.api.unity.com/ugc/", , WithRequestEditorFn(basicAuthProvider.Intercept))
```
