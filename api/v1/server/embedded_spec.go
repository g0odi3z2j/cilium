package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Cilium",
    "title": "Cilium API",
    "version": "v1beta"
  },
  "basePath": "/v1beta",
  "paths": {
    "/config": {
      "get": {
        "description": "Returns the configuration of the Cilium daemon.\n",
        "tags": [
          "daemon"
        ],
        "summary": "Get configuration of Cilium daemon",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/DaemonConfigurationResponse"
            }
          }
        }
      },
      "patch": {
        "description": "Updates the daemon configuration by applying the provided\nConfigurationMap and regenerates \u0026 recompiles all required datapath\ncomponents.\n",
        "tags": [
          "daemon"
        ],
        "summary": "Modify daemon configuration",
        "parameters": [
          {
            "name": "configuration",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ConfigurationMap"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Bad configuration parameters",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Recompilation failed",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/endpoint": {
      "get": {
        "description": "Returns an array of all local endpoints.\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Get list of all endpoints",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Endpoint"
              }
            }
          }
        }
      }
    },
    "/endpoint/{id}": {
      "get": {
        "description": "Returns endpoint information\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Get endpoint by endpoint ID",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Endpoint"
            }
          },
          "400": {
            "description": "Invalid endpoint ID format for specified type",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Invalid"
          },
          "404": {
            "description": "Endpoint not found"
          }
        }
      },
      "put": {
        "description": "Updates an existing endpoint\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Create endpoint",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          },
          {
            "$ref": "#/parameters/endpoint-change-request"
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Invalid endpoint in request",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Invalid"
          },
          "409": {
            "description": "Endpoint already exists",
            "x-go-name": "Exists"
          },
          "500": {
            "description": "Endpoint creation failed",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failed"
          }
        }
      },
      "delete": {
        "description": "Deletes the endpoint specified by the ID. Deletion is imminent and\natomic, if the deletion request is valid and the endpoint exists,\ndeletion will occur even if errors are encountered in the process. If\nerrors have been encountered, the code 202 will be returned, otherwise\n200 on success.\n\nAll resources associated with the endpoint will be freed and the\nworkload represented by the endpoint will be disconnected.It will no\nlonger be able to initiate or receive communications of any sort.\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Delete endpoint",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "206": {
            "description": "Deleted with a number of errors encountered",
            "schema": {
              "type": "integer"
            },
            "x-go-name": "Errors"
          },
          "400": {
            "description": "Invalid endpoint ID format for specified type. Details in error\nmessage\n",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Invalid"
          },
          "404": {
            "description": "Endpoint not found"
          }
        }
      },
      "patch": {
        "description": "Applies the endpoint change request to an existing endpoint\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Modify existing endpoint",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          },
          {
            "$ref": "#/parameters/endpoint-change-request"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid modify endpoint request",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Invalid"
          },
          "404": {
            "description": "Endpoint does not exist"
          },
          "500": {
            "description": "Endpoint update failed",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failed"
          }
        }
      }
    },
    "/endpoint/{id}/config": {
      "get": {
        "description": "Retrieves the configuration of the specified endpoint.\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Retrieve endpoint configuration",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Configuration"
            }
          },
          "404": {
            "description": "Endpoint not found"
          }
        }
      },
      "patch": {
        "description": "Update the configuration of an existing endpoint and regenerates \u0026\nrecompiles the corresponding programs automatically.\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Modify mutable endpoint configuration",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          },
          {
            "name": "configuration",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ConfigurationMap"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid configuration request",
            "x-go-name": "Invalid"
          },
          "404": {
            "description": "Endpoint not found"
          },
          "500": {
            "description": "Update failed. Details in message.",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failed"
          }
        }
      }
    },
    "/endpoint/{id}/labels": {
      "get": {
        "tags": [
          "endpoint"
        ],
        "summary": "Retrieves the list of labels associated with an endpoint.",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/LabelConfiguration"
            }
          },
          "404": {
            "description": "Endpoint not found"
          }
        }
      },
      "put": {
        "description": "Updates the list of labels associated with an endpoint by applying\na label modificator structure to the label configuration of an\nendpoint.\n\nThe label configuration mutation is only executed as a whole, i.e.\nif any of the labels to be deleted are not either on the list of\norchestration system labels, custom labels, or already disabled,\nthen the request will fail. Labels to be added which already exist\non either the orchestration list or custom list will be ignored.\n",
        "tags": [
          "endpoint"
        ],
        "summary": "Modify label configuration of endpoint",
        "parameters": [
          {
            "$ref": "#/parameters/endpoint-id"
          },
          {
            "name": "configuration",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LabelConfigurationModifier"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "404": {
            "description": "Endpoint not found"
          },
          "460": {
            "description": "Label to be deleted not found",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "LabelNotFound"
          },
          "500": {
            "description": "Error while updating labels",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "UpdateFailed"
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "description": "Returns health and status information of the Cilium daemon and related\ncomponents such as the local container runtime, connected datastore,\nKubernetes integration.\n",
        "tags": [
          "daemon"
        ],
        "summary": "Get health of Cilium daemon",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/StatusResponse"
            }
          }
        }
      }
    },
    "/identity": {
      "get": {
        "tags": [
          "policy"
        ],
        "summary": "Retrieve identity by labels",
        "parameters": [
          {
            "name": "labels",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Labels"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Identity"
            }
          },
          "404": {
            "description": "Identity not found"
          },
          "520": {
            "description": "Identity storage unreachable. Likely a network problem.",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Unreachable"
          },
          "521": {
            "description": "Invalid identity format in storage",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "InvalidStorageFormat"
          }
        }
      }
    },
    "/identity/{id}": {
      "get": {
        "tags": [
          "policy"
        ],
        "summary": "Retrieve identity",
        "parameters": [
          {
            "$ref": "#/parameters/identity-id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Identity"
            }
          },
          "400": {
            "description": "Invalid identity provided"
          },
          "404": {
            "description": "Identity not found"
          },
          "520": {
            "description": "Identity storage unreachable. Likely a network problem.",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Unreachable"
          },
          "521": {
            "description": "Invalid identity format in storage",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "InvalidStorageFormat"
          }
        }
      }
    },
    "/ipam": {
      "post": {
        "tags": [
          "ipam"
        ],
        "summary": "Allocate an IP address",
        "parameters": [
          {
            "$ref": "#/parameters/ipam-family"
          }
        ],
        "responses": {
          "201": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/IPAM"
            }
          },
          "502": {
            "description": "Allocation failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/ipam/{ip}": {
      "post": {
        "tags": [
          "ipam"
        ],
        "summary": "Allocate an IP address",
        "parameters": [
          {
            "$ref": "#/parameters/ipam-ip"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid IP address",
            "x-go-name": "Invalid"
          },
          "409": {
            "description": "IP already allocated",
            "x-go-name": "Exists"
          },
          "500": {
            "description": "IP allocation failure. Details in message.",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          },
          "501": {
            "description": "Allocation for address family disabled",
            "x-go-name": "Disabled"
          }
        }
      },
      "delete": {
        "tags": [
          "ipam"
        ],
        "summary": "Release an allocated IP address",
        "parameters": [
          {
            "$ref": "#/parameters/ipam-ip"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid IP address",
            "x-go-name": "Invalid"
          },
          "404": {
            "description": "IP address not found"
          },
          "500": {
            "description": "Address release failure",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          },
          "501": {
            "description": "Allocation for address family disabled",
            "x-go-name": "Disabled"
          }
        }
      }
    },
    "/policy": {
      "get": {
        "description": "Returns the entire policy tree with all children.\n",
        "tags": [
          "policy"
        ],
        "summary": "Retrieve entire policy tree",
        "parameters": [
          {
            "name": "labels",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Labels"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PolicyTree"
            }
          },
          "404": {
            "description": "No policy rules found"
          }
        }
      },
      "put": {
        "tags": [
          "policy"
        ],
        "summary": "Create or update a policy (sub)tree",
        "parameters": [
          {
            "$ref": "#/parameters/policy-rules"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PolicyTree"
            }
          },
          "400": {
            "description": "Invalid policy",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "InvalidPolicy"
          },
          "460": {
            "description": "Invalid path",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "InvalidPath"
          },
          "500": {
            "description": "Policy import failed",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      },
      "delete": {
        "tags": [
          "policy"
        ],
        "summary": "Delete a policy (sub)tree",
        "parameters": [
          {
            "name": "labels",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Labels"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Invalid"
          },
          "404": {
            "description": "Policy tree not found"
          },
          "500": {
            "description": "Error while deleting policy",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    },
    "/policy/resolve": {
      "get": {
        "tags": [
          "policy"
        ],
        "summary": "Resolve policy for an identity context",
        "parameters": [
          {
            "$ref": "#/parameters/identity-context"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PolicyTraceResult"
            }
          }
        }
      }
    },
    "/service": {
      "get": {
        "tags": [
          "service"
        ],
        "summary": "Retrieve list of all services",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Service"
              }
            }
          }
        }
      }
    },
    "/service/{id}": {
      "get": {
        "tags": [
          "service"
        ],
        "summary": "Retrieve configuration of a service",
        "parameters": [
          {
            "$ref": "#/parameters/service-id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Service"
            }
          },
          "404": {
            "description": "Service not found"
          }
        }
      },
      "put": {
        "tags": [
          "service"
        ],
        "summary": "Create or update service",
        "parameters": [
          {
            "$ref": "#/parameters/service-id"
          },
          {
            "$ref": "#/parameters/service-config"
          }
        ],
        "responses": {
          "200": {
            "description": "Updated"
          },
          "201": {
            "description": "Created"
          },
          "460": {
            "description": "Invalid frontend in service configuration",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "InvalidFrontend"
          },
          "461": {
            "description": "Invalid backend in service configuration",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "InvalidBackend"
          },
          "500": {
            "description": "Error while creating service",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      },
      "delete": {
        "tags": [
          "service"
        ],
        "summary": "Delete a service",
        "parameters": [
          {
            "$ref": "#/parameters/service-id"
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "404": {
            "description": "Service not found"
          },
          "500": {
            "description": "Service deletion failed",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "x-go-name": "Failure"
          }
        }
      }
    }
  },
  "definitions": {
    "Address": {
      "description": "IP address",
      "type": "string"
    },
    "BackendAddress": {
      "description": "Service backend address",
      "type": "object",
      "required": [
        "ip"
      ],
      "properties": {
        "ip": {
          "description": "Layer 3 address",
          "type": "string"
        },
        "port": {
          "description": "Layer 4 port number",
          "type": "integer",
          "format": "uint16"
        },
        "weight": {
          "description": "Weight for Round Robin",
          "type": "integer",
          "format": "uint16"
        }
      }
    },
    "Configuration": {
      "description": "General purpose structure to hold configuration of the daemon and\nendpoints. Split into a mutable and immutable section.\n",
      "type": "object",
      "properties": {
        "immutable": {
          "description": "Immutable configuration (read-only)",
          "$ref": "#/definitions/ConfigurationMap"
        },
        "mutable": {
          "description": "Changeable configuration",
          "$ref": "#/definitions/ConfigurationMap"
        }
      }
    },
    "ConfigurationMap": {
      "description": "Map of configuration key/value pairs.\n",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "DaemonConfigurationResponse": {
      "description": "Response to a daemon configuration request. Contains the addressing\ninformation and configuration settings.\n",
      "type": "object",
      "properties": {
        "addressing": {
          "$ref": "#/definitions/NodeAddressing"
        },
        "configuration": {
          "$ref": "#/definitions/Configuration"
        }
      }
    },
    "Endpoint": {
      "description": "Endpoint",
      "type": "object",
      "required": [
        "state"
      ],
      "properties": {
        "addressing": {
          "$ref": "#/definitions/EndpointAddressing"
        },
        "container-id": {
          "description": "ID assigned by container runtime",
          "type": "string"
        },
        "docker-endpoint-id": {
          "description": "Docker endpoint ID",
          "type": "string"
        },
        "docker-network-id": {
          "description": "Docker network ID",
          "type": "string"
        },
        "host-mac": {
          "description": "MAC address",
          "type": "string"
        },
        "id": {
          "description": "Local endpoint ID",
          "type": "integer"
        },
        "identity": {
          "description": "Security identity",
          "$ref": "#/definitions/Identity"
        },
        "interface-index": {
          "description": "Index of network device",
          "type": "integer"
        },
        "interface-name": {
          "description": "Name of network device",
          "type": "string"
        },
        "mac": {
          "description": "MAC address",
          "type": "string"
        },
        "policy": {
          "description": "Policy information of endpoint",
          "$ref": "#/definitions/EndpointPolicy"
        },
        "policy-enabled": {
          "description": "Whether policy enforcement is enabled or not",
          "type": "boolean"
        },
        "state": {
          "description": "Current state of endpoint",
          "$ref": "#/definitions/EndpointState"
        },
        "status": {
          "description": "Current state of endpoint",
          "type": "array",
          "items": {
            "$ref": "#/definitions/EndpointStatusChange"
          }
        }
      }
    },
    "EndpointAddressing": {
      "description": "Addressing information of an endpoint",
      "type": "object",
      "properties": {
        "ipv4": {
          "description": "IPv4 address",
          "type": "string"
        },
        "ipv6": {
          "description": "IPv6 address",
          "type": "string"
        }
      }
    },
    "EndpointChangeRequest": {
      "description": "Structure which contains the mutable elements of an Endpoint.\n",
      "type": "object",
      "required": [
        "state"
      ],
      "properties": {
        "addressing": {
          "$ref": "#/definitions/EndpointAddressing"
        },
        "container-id": {
          "description": "ID assigned by container runtime",
          "type": "string"
        },
        "docker-endpoint-id": {
          "description": "Docker endpoint ID",
          "type": "string"
        },
        "docker-network-id": {
          "description": "Docker network ID",
          "type": "string"
        },
        "host-mac": {
          "description": "MAC address",
          "type": "string"
        },
        "id": {
          "description": "Local endpoint ID",
          "type": "integer"
        },
        "interface-index": {
          "description": "Index of network device",
          "type": "integer"
        },
        "interface-name": {
          "description": "Name of network device",
          "type": "string"
        },
        "mac": {
          "description": "MAC address",
          "type": "string"
        },
        "policy-enabled": {
          "description": "Whether policy enforcement is enabled or not",
          "type": "boolean"
        },
        "state": {
          "description": "Current state of endpoint",
          "$ref": "#/definitions/EndpointState"
        }
      }
    },
    "EndpointPolicy": {
      "description": "Policy information of an endpoint",
      "type": "object",
      "properties": {
        "allowed-consumers": {
          "description": "List of identities allowed to communicate to this endpoint\n",
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "build": {
          "description": "Build number of calculated policy in use",
          "type": "integer"
        },
        "id": {
          "description": "Own identity of endpoint",
          "type": "integer"
        },
        "l4": {
          "$ref": "#/definitions/L4Policy"
        }
      }
    },
    "EndpointState": {
      "description": "State of endpoint",
      "type": "string",
      "enum": [
        "creating",
        "disconnected",
        "waiting-for-identity",
        "not-ready",
        "regenerating",
        "ready"
      ]
    },
    "EndpointStatusChange": {
      "description": "Indication of a change of status",
      "type": "object",
      "properties": {
        "code": {
          "description": "Code indicate type of status change",
          "type": "string",
          "enum": [
            "ok",
            "failed"
          ]
        },
        "message": {
          "description": "Status message",
          "type": "string"
        },
        "timestamp": {
          "description": "Timestamp when status change occurred",
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "string"
    },
    "FrontendAddress": {
      "description": "Layer 4 address",
      "type": "object",
      "properties": {
        "ip": {
          "description": "Layer 3 address",
          "type": "string"
        },
        "port": {
          "description": "Layer 4 port number",
          "type": "integer",
          "format": "uint16"
        },
        "protocol": {
          "description": "Layer 4 protocol",
          "type": "string",
          "enum": [
            "tcp",
            "udp",
            "any"
          ]
        }
      }
    },
    "IPAM": {
      "description": "IPAM configuration of an endpoint",
      "type": "object",
      "required": [
        "endpoint",
        "host-addressing"
      ],
      "properties": {
        "endpoint": {
          "$ref": "#/definitions/EndpointAddressing"
        },
        "host-addressing": {
          "$ref": "#/definitions/NodeAddressing"
        }
      }
    },
    "IPAMStatus": {
      "description": "Status of IP address management",
      "properties": {
        "ipv4": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ipv6": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Identity": {
      "description": "Security identity",
      "type": "object",
      "properties": {
        "id": {
          "description": "Unique identifier",
          "type": "integer"
        },
        "labels": {
          "description": "Labels describing the identity",
          "$ref": "#/definitions/Labels"
        }
      }
    },
    "IdentityContext": {
      "description": "Context describing a pair of source and destination identity",
      "type": "object",
      "properties": {
        "dports": {
          "description": "List of Layer 4 port and protocol pairs which will be used in communication\nfrom the source identity to the destination identity.\n",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Port"
          }
        },
        "from": {
          "$ref": "#/definitions/Labels"
        },
        "to": {
          "$ref": "#/definitions/Labels"
        },
        "verbose": {
          "description": "Enable verbose tracing.\n",
          "type": "boolean"
        }
      }
    },
    "L4Policy": {
      "description": "L4 endpoint policy",
      "type": "object",
      "properties": {
        "egress": {
          "description": "List of L4 egress rules",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ingress": {
          "description": "List of L4 ingress rules",
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "LabelConfiguration": {
      "description": "Label configuration of an endpoint",
      "type": "object",
      "properties": {
        "custom": {
          "description": "Custom labels in addition to orchestration system labels.",
          "$ref": "#/definitions/Labels"
        },
        "disabled": {
          "description": "Labels derived from orchestration system which have been disabled.",
          "$ref": "#/definitions/Labels"
        },
        "orchestration-system": {
          "description": "Labels derived from orchestration system",
          "$ref": "#/definitions/Labels"
        }
      }
    },
    "LabelConfigurationModifier": {
      "description": "Structure describing label mutations to be performed on a\nLabelConfiguration object.\n",
      "type": "object",
      "properties": {
        "add": {
          "description": "List of labels to add and enable. If the label is an orchestration\nsystem label which has been disabled before, it will be removed from\nthe disabled list and readded to the orchestration list. Otherwise\nit will be added to the custom label list.\n",
          "$ref": "#/definitions/Labels"
        },
        "delete": {
          "description": "List of labels to delete. If the label is an orchestration system\nlabel, then it will be deleted from the orchestration list and\nadded to the disabled list. Otherwise it will be removed from the\ncustom list.\n",
          "$ref": "#/definitions/Labels"
        }
      }
    },
    "Labels": {
      "description": "Set of labels",
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "NodeAddressing": {
      "description": "Addressing information of a node for all address families",
      "type": "object",
      "properties": {
        "ipv4": {
          "$ref": "#/definitions/NodeAddressingElement"
        },
        "ipv6": {
          "$ref": "#/definitions/NodeAddressingElement"
        }
      }
    },
    "NodeAddressingElement": {
      "description": "Addressing information",
      "type": "object",
      "properties": {
        "alloc-range": {
          "description": "Address pool to be used for local endpoints",
          "type": "string"
        },
        "enabled": {
          "description": "True if address family is enabled",
          "type": "boolean"
        },
        "ip": {
          "description": "IP address of node",
          "type": "string"
        }
      }
    },
    "PolicyTraceResult": {
      "description": "Response to a policy resolution process",
      "type": "object",
      "properties": {
        "log": {
          "type": "string"
        },
        "verdict": {
          "type": "string"
        }
      }
    },
    "PolicyTree": {
      "description": "Policy tree or subtree",
      "type": "string"
    },
    "Port": {
      "description": "Layer 4 port / protocol pair",
      "type": "object",
      "properties": {
        "port": {
          "description": "Layer 4 port number",
          "type": "integer",
          "format": "uint16"
        },
        "protocol": {
          "description": "Layer 4 protocol",
          "type": "string",
          "enum": [
            "tcp",
            "udp",
            "any"
          ]
        }
      }
    },
    "Service": {
      "description": "Collection of endpoints to be served",
      "type": "object",
      "required": [
        "frontend-address"
      ],
      "properties": {
        "backend-addresses": {
          "description": "List of backend addresses",
          "type": "array",
          "items": {
            "$ref": "#/definitions/BackendAddress"
          }
        },
        "flags": {
          "description": "Optional service configuration flags",
          "type": "object",
          "properties": {
            "active-frontend": {
              "description": "Frontend to backend translation activated",
              "type": "boolean"
            },
            "direct-server-return": {
              "description": "Perform direct server return",
              "type": "boolean"
            }
          }
        },
        "frontend-address": {
          "description": "Frontend address",
          "$ref": "#/definitions/FrontendAddress"
        },
        "id": {
          "description": "Unique identification",
          "type": "integer"
        }
      }
    },
    "Status": {
      "description": "Status of an individual component",
      "type": "object",
      "properties": {
        "msg": {
          "description": "Human readable status/error/warning message",
          "type": "string"
        },
        "state": {
          "description": "State the component is in",
          "type": "string",
          "enum": [
            "Ok",
            "Warning",
            "Failure",
            "Disabled"
          ]
        }
      }
    },
    "StatusResponse": {
      "description": "Health and status information of daemon",
      "type": "object",
      "properties": {
        "cilium": {
          "description": "Status of Cilium daemon",
          "$ref": "#/definitions/Status"
        },
        "container-runtime": {
          "description": "Status of local container runtime",
          "$ref": "#/definitions/Status"
        },
        "ipam": {
          "description": "Status of IP address management",
          "$ref": "#/definitions/IPAMStatus"
        },
        "kubernetes": {
          "description": "Status of Kubernetes integration",
          "$ref": "#/definitions/Status"
        },
        "kvstore": {
          "description": "Status of key/value datastore",
          "$ref": "#/definitions/Status"
        }
      }
    }
  },
  "parameters": {
    "endpoint-change-request": {
      "name": "endpoint",
      "in": "body",
      "required": true,
      "schema": {
        "$ref": "#/definitions/EndpointChangeRequest"
      }
    },
    "endpoint-id": {
      "type": "string",
      "description": "String describing an endpoint with the format ` + "`" + `[prefix:]id` + "`" + `. If no prefix\nis specified, a prefix of ` + "`" + `cilium-local:` + "`" + ` is assumed. Not all endpoints\nwill be addressable by all endpoint ID prefixes with the exception of the\nlocal Cilium UUID which is assigned to all endpoints.\n\nSupported endpoint id prefixes:\n  - cilium-local: Local Cilium endpoint UUID, e.g. cilium-local:3389595\n  - cilium-global: Global Cilium endpoint UUID, e.g. cilium-global:cluster1:nodeX:452343\n  - container-id: Container runtime ID, e.g. container-id:22222\n  - docker-net-endpoint: Docker libnetwork endpoint ID, e.g. docker-net-endpoint:4444\n",
      "name": "id",
      "in": "path",
      "required": true
    },
    "identity-context": {
      "description": "Context to provide policy evaluation on",
      "name": "identity-context",
      "in": "body",
      "schema": {
        "$ref": "#/definitions/IdentityContext"
      }
    },
    "identity-id": {
      "type": "string",
      "description": "Cluster wide unique identifier of a security identity.\n",
      "name": "id",
      "in": "path",
      "required": true
    },
    "ipam-family": {
      "enum": [
        "ipv4",
        "ipv6"
      ],
      "type": "string",
      "name": "family",
      "in": "query"
    },
    "ipam-ip": {
      "type": "string",
      "description": "IP address",
      "name": "ip",
      "in": "path",
      "required": true
    },
    "policy-rules": {
      "description": "Policy rules",
      "name": "policy",
      "in": "body",
      "required": true,
      "schema": {
        "type": "string"
      }
    },
    "service-address": {
      "description": "Service address configuration",
      "name": "address",
      "in": "body",
      "schema": {
        "$ref": "#/definitions/FrontendAddress"
      }
    },
    "service-config": {
      "description": "Service configuration",
      "name": "config",
      "in": "body",
      "required": true,
      "schema": {
        "$ref": "#/definitions/Service"
      }
    },
    "service-id": {
      "type": "integer",
      "description": "ID of service",
      "name": "id",
      "in": "path",
      "required": true
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
}
