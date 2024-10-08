package processors

import (
	"log"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

func SendToTopic(topic string, event string, devEUI string) {
	emitter, err := goka.NewEmitter([]string{"localhost:9092"}, goka.Stream(topic), new(codec.String))
	if err != nil {
		log.Fatalf("Error creating emitter: %v", err)
	}
	defer emitter.Finish()

	err = emitter.EmitSync(devEUI, event)
	if err != nil {
		log.Fatalf("Error emitting message: %v", err)
	}
}
