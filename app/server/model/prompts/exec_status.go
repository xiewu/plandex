package prompts

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"

	shared "plandex-shared"
)

const SysExecStatusFinishedSubtaskXml = `You are tasked with evaluating a response generated by another AI (AI 1) that has been given a coding task to implement.

Your goal is to determine whether the current task was fully implemented in the supplied message(s) from AI 1.

To do this, you need to analyze the latest message from AI 1, and possibly previous messages, and then carefully decide based on the following criteria:

First, examining any previous messages along with the current message, assess whether the current task was fully implemented when these messages are taken together. A task is only considered fully implemented if all necessary code changes for that task have been completed with no remaining todo placeholders or partial implementations.

You MUST output a valid XML response that includes a <subtaskStatus> tag. The <subtaskStatus> tag must contain two child tags:
- <reasoning>: A brief explanation of whether the task was completed and why
- <subtaskFinished>: Either "true" or "false" indicating if the task is done

Do not use XML attributes - put all data as tag content.

Example response:
<subtaskStatus>
<reasoning>Task is complete - all required code changes implemented with no placeholders</reasoning>
<subtaskFinished>true</subtaskFinished>
</subtaskStatus>`

const SysExecStatusFinishedSubtask = `You are tasked with evaluating a response generated by another AI (AI 1) that has been given a coding task to implement.

Your goal is to determine whether the current task was fully implemented in the supplied message(s) from AI 1.

To do this, you need to analyze the latest message from AI 1, and possibly previous messages, and then carefully and decide based on the following criteria:

First, examining any previous messages along with the current message, assess whether the current task was fully implemented when these messages are taken together. A task is only considered fully implemented if all necessary code changes for that task have been completed with no remaining todo placeholders or partial implementations.

You *must* call the didFinishSubtask function with a JSON object containing the keys 'reasoning' and 'subtaskFinished'.

Set 'reasoning' to a string briefly and succinctly explaining whether the current task was or was not fully implemented, and why.

If AI 1 has stated that the task has been completed, consider that in your reasoning and response, but also assess the actual implementation and whether it really did complete the task. Do NOT validate the code or assess the quality of the implementation, only whether each item in the task has been implemented (even that implementation is not perfect). Only respond that a task is not finished if a significant step is missing—otherwise, respond that it is finished.

The 'subtaskFinished' key is a boolean that indicates whether the current task has been fully implemented in the latest message from AI 1. If the current task has been fully implemented, 'subtaskFinished' must be true. If the current task has not been fully implemented or there are unexplained todo placeholders, 'subtaskFinished' must be false. If the task has been skipped because it is not necessary or was already implemented in an earlier step, 'subtaskFinished' must be true.

You must always call 'didFinishSubtask'. Don't call any other function.`

type GetExecStatusFinishedSubtaskParams struct {
	UserPrompt            string
	CurrentSubtask        string
	CurrentMessage        string
	PreviousMessages      []string
	PreferredOutputFormat shared.ModelOutputFormat
}

func GetExecStatusFinishedSubtask(params GetExecStatusFinishedSubtaskParams) string {
	userPrompt := params.UserPrompt
	currentSubtask := params.CurrentSubtask
	currentMessage := params.CurrentMessage
	previousMessages := params.PreviousMessages
	preferredOutputFormat := params.PreferredOutputFormat

	var s string
	if preferredOutputFormat == shared.ModelOutputFormatXml {
		s = SysExecStatusFinishedSubtaskXml
	} else {
		s = SysExecStatusFinishedSubtask
	}

	if userPrompt != "" {
		s += "\n\n**Here is the user's prompt:**\n" + userPrompt
	}
	s += "\n\n**Here is the current task:**\n" + currentSubtask

	for _, msg := range previousMessages {
		s += "\n\n**Here is a previous message from AI 1 that was working on the same task:**\n" + msg
	}

	s += "\n\n**Here is the latest message from AI 1:**\n" + currentMessage

	return s
}

var DidFinishSubtaskFn = openai.FunctionDefinition{
	Name: "didFinishSubtask",
	Parameters: &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"reasoning": {
				Type: jsonschema.String,
			},
			"subtaskFinished": {
				Type: jsonschema.Boolean,
			},
		},
		Required: []string{"reasoning", "subtaskFinished"},
	},
}
