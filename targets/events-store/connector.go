package events_store

import "github.com/kubemq-hub/builder/common"

func Connector() *common.Connector {
	return common.NewConnector().
		SetKind("target.events-store").
		SetDescription("Kubemq Events-Store Target").
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("address").
				SetDescription("Sets Kubemq grpc endpoint address").
				SetMust(true).
				SetDefault("localhost:50000"),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("channel").
				SetDescription("Sets Events-Store channel").
				SetMust(true).
				SetDefault("events-store"),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("client_id").
				SetDescription("Sets Events-Store connection client Id").
				SetMust(false).
				SetDefault(""),
		).
		AddProperty(
			common.NewProperty().
				SetKind("string").
				SetName("auth_token").
				SetDescription("Sets Events-Store connection Authentication token").
				SetMust(false).
				SetDefault(""),
		)
}