/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type VirtualPMemDevice struct {

	HostPath string `json:"HostPath,omitempty"`

	ReadOnly bool `json:"ReadOnly,omitempty"`

	ImageFormat string `json:"ImageFormat,omitempty"`
}
