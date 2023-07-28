/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CodeHookObservation struct {

	// The version of the request-response that you want Amazon Lex to use
	// to invoke your Lambda function. For more information, see
	// Using Lambda Functions. Must be less than or equal to 5 characters in length.
	MessageVersion *string `json:"messageVersion,omitempty" tf:"message_version,omitempty"`

	// The Amazon Resource Name (ARN) of the Lambda function.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type CodeHookParameters struct {

	// The version of the request-response that you want Amazon Lex to use
	// to invoke your Lambda function. For more information, see
	// Using Lambda Functions. Must be less than or equal to 5 characters in length.
	// +kubebuilder:validation:Required
	MessageVersion *string `json:"messageVersion" tf:"message_version,omitempty"`

	// The Amazon Resource Name (ARN) of the Lambda function.
	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type ConclusionStatementMessageObservation struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ConclusionStatementMessageParameters struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ConclusionStatementObservation struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	Message []ConclusionStatementMessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type ConclusionStatementParameters struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	// +kubebuilder:validation:Required
	Message []ConclusionStatementMessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type ConfirmationPromptMessageObservation struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ConfirmationPromptMessageParameters struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ConfirmationPromptObservation struct {

	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	Message []ConfirmationPromptMessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type ConfirmationPromptParameters struct {

	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Required
	MaxAttempts *float64 `json:"maxAttempts" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	// +kubebuilder:validation:Required
	Message []ConfirmationPromptMessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type DialogCodeHookObservation struct {

	// The version of the request-response that you want Amazon Lex to use
	// to invoke your Lambda function. For more information, see
	// Using Lambda Functions. Must be less than or equal to 5 characters in length.
	MessageVersion *string `json:"messageVersion,omitempty" tf:"message_version,omitempty"`

	// The Amazon Resource Name (ARN) of the Lambda function.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type DialogCodeHookParameters struct {

	// The version of the request-response that you want Amazon Lex to use
	// to invoke your Lambda function. For more information, see
	// Using Lambda Functions. Must be less than or equal to 5 characters in length.
	// +kubebuilder:validation:Required
	MessageVersion *string `json:"messageVersion" tf:"message_version,omitempty"`

	// The Amazon Resource Name (ARN) of the Lambda function.
	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type FollowUpPromptObservation struct {

	// Prompts for information from the user. Attributes are documented under prompt.
	Prompt []PromptObservation `json:"prompt,omitempty" tf:"prompt,omitempty"`

	// If the user answers "no" to the question defined in the prompt field,
	// Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	// documented below under statement.
	RejectionStatement []RejectionStatementObservation `json:"rejectionStatement,omitempty" tf:"rejection_statement,omitempty"`
}

type FollowUpPromptParameters struct {

	// Prompts for information from the user. Attributes are documented under prompt.
	// +kubebuilder:validation:Required
	Prompt []PromptParameters `json:"prompt" tf:"prompt,omitempty"`

	// If the user answers "no" to the question defined in the prompt field,
	// Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	// documented below under statement.
	// +kubebuilder:validation:Required
	RejectionStatement []RejectionStatementParameters `json:"rejectionStatement" tf:"rejection_statement,omitempty"`
}

type FulfillmentActivityObservation struct {

	// A description of the Lambda function that is run to fulfill the intent.
	// Required if type is CodeHook. Attributes are documented under code_hook.
	CodeHook []CodeHookObservation `json:"codeHook,omitempty" tf:"code_hook,omitempty"`

	// How the intent should be fulfilled, either by running a Lambda function or by
	// returning the slot data to the client application. Type can be either ReturnIntent or CodeHook, as documented here.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FulfillmentActivityParameters struct {

	// A description of the Lambda function that is run to fulfill the intent.
	// Required if type is CodeHook. Attributes are documented under code_hook.
	// +kubebuilder:validation:Optional
	CodeHook []CodeHookParameters `json:"codeHook,omitempty" tf:"code_hook,omitempty"`

	// How the intent should be fulfilled, either by running a Lambda function or by
	// returning the slot data to the client application. Type can be either ReturnIntent or CodeHook, as documented here.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type IntentObservation_2 struct {

	// The ARN of the Lex intent.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Checksum identifying the version of the intent that was created. The checksum is not
	// included as an argument because the resource will add it automatically when updating the intent.
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// The statement that you want Amazon Lex to convey to the user
	// after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the fulfillment_activity. If you return the intent to the client
	// application, you can't specify this element. The follow_up_prompt and conclusion_statement are
	// mutually exclusive. You can specify only one. Attributes are documented under statement.
	ConclusionStatement []ConclusionStatementObservation `json:"conclusionStatement,omitempty" tf:"conclusion_statement,omitempty"`

	// Prompts the user to confirm the intent. This question should
	// have a yes or no answer. You you must provide both the rejection_statement and confirmation_prompt,
	// or neither. Attributes are documented under prompt.
	ConfirmationPrompt []ConfirmationPromptObservation `json:"confirmationPrompt,omitempty" tf:"confirmation_prompt,omitempty"`

	// Determines if a new slot type version is created when the initial
	// resource is created and on each update. Defaults to false.
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// The date when the intent version was created.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// A description of the intent. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies a Lambda function to invoke for each user input. You can
	// invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	DialogCodeHook []DialogCodeHookObservation `json:"dialogCodeHook,omitempty" tf:"dialog_code_hook,omitempty"`

	// Amazon Lex uses this prompt to solicit additional activity after
	// fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	// user to order a drink. The follow_up_prompt field and the conclusion_statement field are mutually
	// exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	FollowUpPrompt []FollowUpPromptObservation `json:"followUpPrompt,omitempty" tf:"follow_up_prompt,omitempty"`

	// Describes how the intent is fulfilled. For example, after a
	// user provides all of the information for a pizza order, fulfillment_activity defines how the bot
	// places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	FulfillmentActivity []FulfillmentActivityObservation `json:"fulfillmentActivity,omitempty" tf:"fulfillment_activity,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date when the $LATEST version of this intent was updated.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// Standard Built-in Intents
	// in the Alexa Skills Kit.
	ParentIntentSignature *string `json:"parentIntentSignature,omitempty" tf:"parent_intent_signature,omitempty"`

	// When the user answers "no" to the question defined in
	// confirmation_prompt, Amazon Lex responds with this statement to acknowledge that the intent was
	// canceled. You must provide both the rejection_statement and the confirmation_prompt, or neither.
	// Attributes are documented under statement.
	RejectionStatement []IntentRejectionStatementObservation `json:"rejectionStatement,omitempty" tf:"rejection_statement,omitempty"`

	// An array of utterances (strings) that a user might say to signal
	// the intent. For example, "I want {PizzaSize} pizza", "Order {Quantity} {PizzaSize} pizzas".
	// In each utterance, a slot name is enclosed in curly braces. Must have between 1 and 10 items in the list, and each item must be less than or equal to 200 characters in length.
	SampleUtterances []*string `json:"sampleUtterances,omitempty" tf:"sample_utterances,omitempty"`

	// An list of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. Attributes are documented under slot.
	Slot []SlotObservation `json:"slot,omitempty" tf:"slot,omitempty"`

	// The version of the bot.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type IntentParameters_2 struct {

	// The statement that you want Amazon Lex to convey to the user
	// after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	// you provide a Lambda function in the fulfillment_activity. If you return the intent to the client
	// application, you can't specify this element. The follow_up_prompt and conclusion_statement are
	// mutually exclusive. You can specify only one. Attributes are documented under statement.
	// +kubebuilder:validation:Optional
	ConclusionStatement []ConclusionStatementParameters `json:"conclusionStatement,omitempty" tf:"conclusion_statement,omitempty"`

	// Prompts the user to confirm the intent. This question should
	// have a yes or no answer. You you must provide both the rejection_statement and confirmation_prompt,
	// or neither. Attributes are documented under prompt.
	// +kubebuilder:validation:Optional
	ConfirmationPrompt []ConfirmationPromptParameters `json:"confirmationPrompt,omitempty" tf:"confirmation_prompt,omitempty"`

	// Determines if a new slot type version is created when the initial
	// resource is created and on each update. Defaults to false.
	// +kubebuilder:validation:Optional
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// A description of the intent. Must be less than or equal to 200 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies a Lambda function to invoke for each user input. You can
	// invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	// +kubebuilder:validation:Optional
	DialogCodeHook []DialogCodeHookParameters `json:"dialogCodeHook,omitempty" tf:"dialog_code_hook,omitempty"`

	// Amazon Lex uses this prompt to solicit additional activity after
	// fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	// user to order a drink. The follow_up_prompt field and the conclusion_statement field are mutually
	// exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	// +kubebuilder:validation:Optional
	FollowUpPrompt []FollowUpPromptParameters `json:"followUpPrompt,omitempty" tf:"follow_up_prompt,omitempty"`

	// Describes how the intent is fulfilled. For example, after a
	// user provides all of the information for a pizza order, fulfillment_activity defines how the bot
	// places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	// +kubebuilder:validation:Optional
	FulfillmentActivity []FulfillmentActivityParameters `json:"fulfillmentActivity,omitempty" tf:"fulfillment_activity,omitempty"`

	// A unique identifier for the built-in intent to base this
	// intent on. To find the signature for an intent, see
	// Standard Built-in Intents
	// in the Alexa Skills Kit.
	// +kubebuilder:validation:Optional
	ParentIntentSignature *string `json:"parentIntentSignature,omitempty" tf:"parent_intent_signature,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// When the user answers "no" to the question defined in
	// confirmation_prompt, Amazon Lex responds with this statement to acknowledge that the intent was
	// canceled. You must provide both the rejection_statement and the confirmation_prompt, or neither.
	// Attributes are documented under statement.
	// +kubebuilder:validation:Optional
	RejectionStatement []IntentRejectionStatementParameters `json:"rejectionStatement,omitempty" tf:"rejection_statement,omitempty"`

	// An array of utterances (strings) that a user might say to signal
	// the intent. For example, "I want {PizzaSize} pizza", "Order {Quantity} {PizzaSize} pizzas".
	// In each utterance, a slot name is enclosed in curly braces. Must have between 1 and 10 items in the list, and each item must be less than or equal to 200 characters in length.
	// +kubebuilder:validation:Optional
	SampleUtterances []*string `json:"sampleUtterances,omitempty" tf:"sample_utterances,omitempty"`

	// An list of intent slots. At runtime, Amazon Lex elicits required slot values
	// from the user using prompts defined in the slots. Attributes are documented under slot.
	// +kubebuilder:validation:Optional
	Slot []SlotParameters `json:"slot,omitempty" tf:"slot,omitempty"`
}

type IntentRejectionStatementMessageObservation struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type IntentRejectionStatementMessageParameters struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type IntentRejectionStatementObservation struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	Message []IntentRejectionStatementMessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type IntentRejectionStatementParameters struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	// +kubebuilder:validation:Required
	Message []IntentRejectionStatementMessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type PromptMessageObservation struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type PromptMessageParameters struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type PromptObservation struct {

	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	Message []PromptMessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type PromptParameters struct {

	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Required
	MaxAttempts *float64 `json:"maxAttempts" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	// +kubebuilder:validation:Required
	Message []PromptMessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type RejectionStatementMessageObservation struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type RejectionStatementMessageParameters struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type RejectionStatementObservation struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	Message []RejectionStatementMessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type RejectionStatementParameters struct {

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	// +kubebuilder:validation:Required
	Message []RejectionStatementMessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type SlotObservation struct {

	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the intent slot that you want to create. The name is case sensitive. Must be less than or equal to 100 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Directs Lex the order in which to elicit this slot value from the user.
	// For example, if the intent has two slots with priorities 1 and 2, AWS Lex first elicits a value for
	// the slot with priority 1. If multiple slots share the same priority, the order in which Lex elicits
	// values is arbitrary. Must be between 1 and 100.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`

	// If you know a specific pattern with which users might respond to
	// an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	// is optional. In most cases, Amazon Lex is capable of understanding user utterances. Must have between 1 and 10 items in the list, and each item must be less than or equal to 200 characters in length.
	SampleUtterances []*string `json:"sampleUtterances,omitempty" tf:"sample_utterances,omitempty"`

	// Specifies whether the slot is required or optional.
	SlotConstraint *string `json:"slotConstraint,omitempty" tf:"slot_constraint,omitempty"`

	// The type of the slot, either a custom slot type that you defined or one of
	// the built-in slot types. Must be less than or equal to 100 characters in length.
	SlotType *string `json:"slotType,omitempty" tf:"slot_type,omitempty"`

	// The version of the slot type. Must be less than or equal to 64 characters in length.
	SlotTypeVersion *string `json:"slotTypeVersion,omitempty" tf:"slot_type_version,omitempty"`

	// The prompt that Amazon Lex uses to elicit the slot value
	// from the user. Attributes are documented under prompt.
	ValueElicitationPrompt []ValueElicitationPromptObservation `json:"valueElicitationPrompt,omitempty" tf:"value_elicitation_prompt,omitempty"`
}

type SlotParameters struct {

	// A description of the bot. Must be less than or equal to 200 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the intent slot that you want to create. The name is case sensitive. Must be less than or equal to 100 characters in length.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Directs Lex the order in which to elicit this slot value from the user.
	// For example, if the intent has two slots with priorities 1 and 2, AWS Lex first elicits a value for
	// the slot with priority 1. If multiple slots share the same priority, the order in which Lex elicits
	// values is arbitrary. Must be between 1 and 100.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`

	// If you know a specific pattern with which users might respond to
	// an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	// is optional. In most cases, Amazon Lex is capable of understanding user utterances. Must have between 1 and 10 items in the list, and each item must be less than or equal to 200 characters in length.
	// +kubebuilder:validation:Optional
	SampleUtterances []*string `json:"sampleUtterances,omitempty" tf:"sample_utterances,omitempty"`

	// Specifies whether the slot is required or optional.
	// +kubebuilder:validation:Required
	SlotConstraint *string `json:"slotConstraint" tf:"slot_constraint,omitempty"`

	// The type of the slot, either a custom slot type that you defined or one of
	// the built-in slot types. Must be less than or equal to 100 characters in length.
	// +kubebuilder:validation:Required
	SlotType *string `json:"slotType" tf:"slot_type,omitempty"`

	// The version of the slot type. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	SlotTypeVersion *string `json:"slotTypeVersion,omitempty" tf:"slot_type_version,omitempty"`

	// The prompt that Amazon Lex uses to elicit the slot value
	// from the user. Attributes are documented under prompt.
	// +kubebuilder:validation:Optional
	ValueElicitationPrompt []ValueElicitationPromptParameters `json:"valueElicitationPrompt,omitempty" tf:"value_elicitation_prompt,omitempty"`
}

type ValueElicitationPromptMessageObservation struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The content type of the message string.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ValueElicitationPromptMessageParameters struct {

	// The text of the message. Must be less than or equal to 1000 characters in length.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The content type of the message string.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Identifies the message group that the message belongs to. When a group
	// is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Optional
	GroupNumber *float64 `json:"groupNumber,omitempty" tf:"group_number,omitempty"`
}

type ValueElicitationPromptObservation struct {

	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	MaxAttempts *float64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	Message []ValueElicitationPromptMessageObservation `json:"message,omitempty" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

type ValueElicitationPromptParameters struct {

	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	// +kubebuilder:validation:Required
	MaxAttempts *float64 `json:"maxAttempts" tf:"max_attempts,omitempty"`

	// A set of messages, each of which provides a message string and its type.
	// You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	// Attributes are documented under message. Must contain between 1 and 15 messages.
	// +kubebuilder:validation:Required
	Message []ValueElicitationPromptMessageParameters `json:"message" tf:"message,omitempty"`

	// The response card. Amazon Lex will substitute session attributes and
	// slot values into the response card. For more information, see
	// Example: Using a Response Card. Must be less than or equal to 50000 characters in length.
	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card,omitempty"`
}

// IntentSpec defines the desired state of Intent
type IntentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntentParameters_2 `json:"forProvider"`
}

// IntentStatus defines the observed state of Intent.
type IntentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntentObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Intent is the Schema for the Intents API. Provides an Amazon Lex intent resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Intent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fulfillmentActivity)",message="fulfillmentActivity is a required parameter"
	Spec   IntentSpec   `json:"spec"`
	Status IntentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntentList contains a list of Intents
type IntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Intent `json:"items"`
}

// Repository type metadata.
var (
	Intent_Kind             = "Intent"
	Intent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Intent_Kind}.String()
	Intent_KindAPIVersion   = Intent_Kind + "." + CRDGroupVersion.String()
	Intent_GroupVersionKind = CRDGroupVersion.WithKind(Intent_Kind)
)

func init() {
	SchemeBuilder.Register(&Intent{}, &IntentList{})
}
