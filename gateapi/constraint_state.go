/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ConstraintState struct {

	ArtifactVersion string `json:"artifactVersion,omitempty"`

	Attributes *interface{} `json:"attributes,omitempty"`

	Comment string `json:"comment,omitempty"`

	CreatedAt string `json:"createdAt,omitempty"`

	DeliveryConfigName string `json:"deliveryConfigName,omitempty"`

	EnvironmentName string `json:"environmentName,omitempty"`

	JudgedAt string `json:"judgedAt,omitempty"`

	JudgedBy string `json:"judgedBy,omitempty"`

	Status string `json:"status,omitempty"`

	Type_ string `json:"type,omitempty"`
}
