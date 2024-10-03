package processors

import "fmt"

func StartProcessor() error {
	// TODO: Implement starting processor logic

	fmt.Println("StartProcessor() was called")
	return nil
}

//Stops the processor
func StopProcessor() {
	// TODO: Implement stopping processor
}

//Pauses the processor
func PauseProcessor() {
	// TODO: Implement pausing processor
}

//Resumes the processor
func ResumeProcessor() {
	// TODO: Implement resuming processor
}

// ProcessEvent processes an event with the provided processor
func ProcessEvent() {
	// TODO: Implement event processing logic
}

//Send processed events to (compacted)topics
func ProduceToTopic() {
	// TODO: Implement producing to (compacted)topics
}
