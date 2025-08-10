package main

import (
	sdk "github.com/webdock-io/go-sdk/sdk"
)

var New = sdk.New

type CreatePublicKeyOptions = sdk.CreatePublicKeyOptions
type CreateAccountScriptOptions = sdk.CreateAccountScriptOptions
type AccountScriptUpdateOptions = sdk.AccountScriptUpdateOptions
type CreateEventHookOptions = sdk.CreateEventHookOptions
type ListProfilesOptions = sdk.ListProfilesOptions
type GetProfilesOptions = sdk.GetProfilesOptions
type ListServersQuery = sdk.ListServersQuery
type CreateServerScriptOptions = sdk.CreateServerScriptOptions
type GetServerScriptGetByIdOption = sdk.GetServerScriptGetByIdOption
type CreateShellUserOptions = sdk.CreateShellUserOptions
type UpdateServerShellUserOptions = sdk.UpdateServerShellUserOptions
type UpdateServerOptions = sdk.UpdateServerOptions
