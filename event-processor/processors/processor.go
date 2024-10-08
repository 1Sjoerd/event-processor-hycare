package processors

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/lovoo/goka/storage"
)

var (
	brokers             = []string{"localhost:9092"}
	topic   goka.Stream = "iot-stream"
	group   goka.Group  = "device-info"
)

func StartEventProcessor() {
	storagePath := "./cache"
	err := os.MkdirAll(storagePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating cache directory: %v", err)
	}

	log.Printf("Starting event processor with cache directory: %s", storagePath)

	cb := func(ctx goka.Context, msg interface{}) {
		var parsedEvent map[string]interface{}
		err := json.Unmarshal([]byte(msg.(string)), &parsedEvent)
		if err != nil {
			log.Fatalf("Error parsing event: %v", err)
		}

		deviceInfo, ok := parsedEvent["device_info"].(map[string]interface{})
		if !ok {
			log.Fatalf("Error extracting device_info from event: expected map but got different type")
		}

		devEUI, ok := deviceInfo["dev_eui"].(string)
		if !ok {
			log.Fatalf("Error extracting dev_eui from event: expected string but got different type")
		}

		var currentData map[string]interface{}
		if val := ctx.Value(); val != nil {
			currentData, ok = val.(map[string]interface{})
			if !ok {
				log.Println("Continuing without current data.")
				currentData = make(map[string]interface{})
			}
		} else {
			currentData = make(map[string]interface{})
		}

		processedEvent := ProcessEvent(ctx.Key(), msg.(string), currentData)

		SendToTopic("device-info-table", processedEvent, devEUI)
	}

	g := goka.DefineGroup(group,
		goka.Input(topic, new(codec.String), cb),
		goka.Persist(new(codec.String)),
	)

	storageBuilder := storage.DefaultBuilder(storagePath)

	processor, err := goka.NewProcessor(brokers, g, goka.WithStorageBuilder(storageBuilder))
	if err != nil {
		log.Fatalf("Error creating processor: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		if err := processor.Run(ctx); err != nil {
			log.Fatalf("Error running processor: %v", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	select {
	case <-sigs:
		log.Println("Shutting down processor...")
		cancel()
	case <-done:
	}
}
