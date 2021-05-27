// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rookie-ninja/rk-boot"
	rkcommon "github.com/rookie-ninja/rk-common/common"
	rkentry "github.com/rookie-ninja/rk-entry/entry"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is gin server with rk-boot.
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.basic BasicAuth
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http https
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	boot.GetGinEntry("greeter").Router.GET("/v1/hello", hello)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown signal
	boot.WaitForShutdownSig()

	// Interrupt entries
	boot.Interrupt(context.TODO())
}

// @Summary Hello
// @Description Say hello to incoming name.
// @Id v1.api.hello
// @Accept  application/json
// @Tags Hello
// @version 1.0
// @Param name query string true "Your name"
// @Produce application/json
// @Success 200 {object} helloResponse
// @Failure 400 {object} httpError
// @Router /v1/hello [get]
// @Header all {string} request-id "Request id for with uuid generator."
func hello(ctx *gin.Context) {
	ctx.Header("request-id", uuid.New().String())

	if name := ctx.Query("name"); len(name) < 1 {
		NewError(ctx, http.StatusBadRequest, errors.New("name should not be nil"))
		return
	}

	ctx.JSON(http.StatusOK, &helloResponse{
		Response: "hello " + ctx.Query("name"),
	})
}

type helloResponse struct {
	Response string `json:"response" yaml:"response" example:"hello user"`
}

func NewError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, httpError{
		Code:    status,
		Message: err.Error(),
	})
}

type httpError struct {
	Code    int    `json:"code" yaml:"code" example:"400"`
	Message string `json:"message" yaml:"message" example:"status bad request"`
}

// Register entry, must be in init() function since we need to register entry at beginning
func init() {
	rkentry.RegisterEntryRegFunc(RegisterMyEntriesFromConfig)
}

// A struct which is for unmarshalled YAML
type BootConfig struct {
	MyEntry struct {
		Enabled     bool   `yaml:"enabled" json:"enabled"`
		Name        string `yaml:"name" json:"name"`
		Description string `yaml:"description" json:"description"`
		Key         string `yaml:"key" json:"key"`
		Logger      struct {
			ZapLogger struct {
				Ref string `yaml:"ref" json:"ref"`
			} `yaml:"zapLogger" json:"zapLogger"`
			EventLogger struct {
				Ref string `yaml:"ref" json:"ref"`
			} `yaml:"eventLogger" json:"eventLogger"`
		} `yaml:"logger" json:"logger"`
	} `yaml:"myEntry" json:"myEntry"`
}

// An implementation of:
// type EntryRegFunc func(string) map[string]rkentry.Entry
func RegisterMyEntriesFromConfig(configFilePath string) map[string]rkentry.Entry {
	res := make(map[string]rkentry.Entry)

	// 1: decode config map into boot config struct
	config := &BootConfig{}
	rkcommon.UnmarshalBootConfig(configFilePath, config)

	// 3: construct entry
	if config.MyEntry.Enabled {
		zapLoggerEntry := rkentry.GlobalAppCtx.GetZapLoggerEntry(config.MyEntry.Logger.ZapLogger.Ref)
		eventLoggerEntry := rkentry.GlobalAppCtx.GetEventLoggerEntry(config.MyEntry.Logger.EventLogger.Ref)

		entry := RegisterMyEntry(
			WithName(config.MyEntry.Name),
			WithDescription(config.MyEntry.Description),
			WithKey(config.MyEntry.Key),
			WithZapLoggerEntry(zapLoggerEntry),
			WithEventLoggerEntry(eventLoggerEntry))
		res[entry.GetName()] = entry
	}

	return res
}

func RegisterMyEntry(opts ...MyEntryOption) *MyEntry {
	entry := &MyEntry{
		EntryName:        "default",
		EntryType:        "myEntry",
		EntryDescription: "Please contact maintainers to add description of this entry.",
		ZapLoggerEntry:   rkentry.GlobalAppCtx.GetZapLoggerEntryDefault(),
		EventLoggerEntry: rkentry.GlobalAppCtx.GetEventLoggerEntryDefault(),
	}

	for i := range opts {
		opts[i](entry)
	}

	if len(entry.EntryName) < 1 {
		entry.EntryName = "my-default"
	}

	if len(entry.EntryDescription) < 1 {
		entry.EntryDescription = "Please contact maintainers to add description of this entry."
	}

	rkentry.GlobalAppCtx.AddEntry(entry)

	return entry
}

type MyEntryOption func(*MyEntry)

func WithName(name string) MyEntryOption {
	return func(entry *MyEntry) {
		entry.EntryName = name
	}
}

func WithDescription(description string) MyEntryOption {
	return func(entry *MyEntry) {
		entry.EntryDescription = description
	}
}

func WithKey(key string) MyEntryOption {
	return func(entry *MyEntry) {
		entry.Key = key
	}
}

func WithZapLoggerEntry(zapLoggerEntry *rkentry.ZapLoggerEntry) MyEntryOption {
	return func(entry *MyEntry) {
		if zapLoggerEntry != nil {
			entry.ZapLoggerEntry = zapLoggerEntry
		}
	}
}

func WithEventLoggerEntry(eventLoggerEntry *rkentry.EventLoggerEntry) MyEntryOption {
	return func(entry *MyEntry) {
		if eventLoggerEntry != nil {
			entry.EventLoggerEntry = eventLoggerEntry
		}
	}
}

type MyEntry struct {
	EntryName        string                    `json:"entryName" yaml:"entryName"`
	EntryType        string                    `json:"entryType" yaml:"entryType"`
	EntryDescription string                    `json:"entryDescription" yaml:"entryDescription"`
	Key              string                    `json:"key" yaml:"key"`
	ZapLoggerEntry   *rkentry.ZapLoggerEntry   `json:"zapLoggerEntry" yaml:"zapLoggerEntry"`
	EventLoggerEntry *rkentry.EventLoggerEntry `json:"eventLoggerEntry" yaml:"eventLoggerEntry"`
}

func (entry *MyEntry) Bootstrap(context.Context) {
	event := entry.EventLoggerEntry.GetEventHelper().Start("bootstrap")
	event.AddPair("key", entry.Key)
	entry.EventLoggerEntry.GetEventHelper().Finish(event)
}

func (entry *MyEntry) Interrupt(context.Context) {}

func (entry *MyEntry) GetName() string {
	return entry.EntryName
}

func (entry *MyEntry) GetType() string {
	return entry.EntryType
}

func (entry *MyEntry) String() string {
	bytes, _ := json.Marshal(entry)

	return string(bytes)
}

// Marshal entry
func (entry *MyEntry) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"entryName":        entry.EntryName,
		"entryType":        entry.EntryType,
		"entryDescription": entry.EntryDescription,
		"eventLoggerEntry": entry.EventLoggerEntry.GetName(),
		"zapLoggerEntry":   entry.ZapLoggerEntry.GetName(),
		"key":              entry.Key,
	}

	return json.Marshal(&m)
}

// Unmarshal entry
func (entry *MyEntry) UnmarshalJSON([]byte) error {
	return nil
}

// Get description of entry
func (entry *MyEntry) GetDescription() string {
	return entry.EntryDescription
}
