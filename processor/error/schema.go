package error

func Schema() string {
	return errorSchema
}

var errorSchema = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "docs/spec/errors/wrapper.json",
    "title": "Errors Wrapper",
    "description": "List of errors wrapped in an object containing some other attributes normalized away form the errors themselves",
    "type": "object",
    "properties": {
        "service": {
                "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "doc/spec/service.json",
    "title": "Service",
    "type": "object",
    "properties": {
        "agent": {
            "description": "Name and version of the Elastic APM agent",
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name of the Elastic APM agent, e.g. \"Python\"",
                    "type": "string",
                    "maxLength": 1024
                },
                "version": {
                    "description": "Version of the Elastic APM agent, e.g.\"1.0.0\"",
                    "type": "string",
                    "maxLength": 1024
                }
            },
            "required": ["name", "version"]
        },
        "argv": {
            "description": "Command line arguments used to start this service",
            "type": ["array", "null"],
            "minItems": 0
        },
        "framework": {
            "description": "Name and version of the used web framework",
            "type": ["object", "null"],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 1024
                },
                "version": {
                    "type": "string",
                    "maxLength": 1024
                }
            },
            "required": ["name", "version"]
        },
        "language": {
            "description": "Name and version of the used programming language",
            "type": ["object", "null"],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 1024
                },
                "version": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                }
            },
            "required": ["name"]
        },
        "name": {
            "description": "Immutable name of the service emitting this event",
            "type": "string",
            "pattern": "^[a-zA-Z0-9 _-]+$",
            "maxLength": 1024
        },
        "pid": {
            "description": "Process ID of the service",
            "type": ["number", "null"]
        },
        "process_title": {
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "environment": {
            "description": "Environment name of the service, e.g. \"production\" or \"staging\"",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "runtime": {
            "description": "Name and version of the language runtime running this service",
            "type": ["object", "null"],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 1024
                },
                "version": {
                    "type": "string",
                    "maxLength": 1024
                }
            },
            "required": ["name", "version"]
        },
        "version": {
            "description": "Version of the service emitting this event",
            "type": ["string", "null"],
            "maxLength": 1024
        }
    },
    "required": ["agent", "name"]
        },
        "errors": {
            "type": "array",
            "items": {
                    "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "docs/spec/errors/error.json",
    "type": "object",
    "description": "Data captured by an agent representing an event occurring in a monitored service",
    "properties": {
        "context": {
                "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "doc/spec/context.json",
    "title": "Context",
    "description": "Any arbitrary contextual information regarding the event, captured by the agent, optionally provided by the user",
    "type": ["object", "null"],
    "properties": {
        "custom": {
            "description": "An arbitrary mapping of additional metadata to store with the event.",
            "type": ["object", "null"],
            "regexProperties": true,
            "patternProperties": {
                "^[^.*\"]*$": {}
            },
            "additionalProperties": false
        },
        "response": {
            "type": ["object", "null"],
            "properties": {
                "finished": {
                    "description": "A boolean indicating the the response was finished",
                    "type": ["boolean", "null"]
                },
                "headers": {
                    "description": "A mapping of HTTP headers of the response object",
                    "type": ["object", "null"],
                    "properties": {
                        "content-type": {
                            "type": ["string", "null"]
                        }
                    }
                },
                "headers_sent": {
                    "type": ["boolean", "null"]
                },
                "status_code": {
                    "description": "The HTTP status code of the response.",
                    "type": ["number", "null"]
                }
            }
        },
        "request": {
                "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "docs/spec/http.json",
    "title": "Request",
    "description": "If a log record was generated as a result of a http request, the http interface can be used to collect this information.",
    "type": ["object", "null"],
    "properties": {
        "body": {
            "description": "Data should only contain the request body (not the query string). It can either be a dictionary (for standard HTTP requests) or a raw request body.",
            "type": ["object", "string", "null"]
        },
        "env": {
            "description": "The env variable is a compounded of environment information passed from the webserver.",
            "type": ["object", "null"],
            "properties": {}
        },
        "headers": {
            "description": "Should include any headers sent by the requester. Cookies will be taken by headers if supplied.",
            "type": ["object", "null"],
            "properties": {
                "content-type": {
                    "type": ["string", "null"]
                },
                "cookie": {
                    "description": "Cookies sent with the request. It is expected to have values delimited by semicolons.",
                    "type": ["string", "null"]
                },
                "user-agent": {
                    "type": ["string", "null"]
                }
            }
        },
        "http_version": {
            "description": "HTTP version.",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "method": {
            "description": "HTTP method.",
            "type": "string",
            "maxLength": 1024
        },
        "socket": {
            "type": ["object", "null"],
            "properties": {
                "encrypted": {
                    "description": "Indicates whether request was sent as SSL/HTTPS request.",
                    "type": ["boolean", "null"]
                },
                "remote_address": {
                    "type": ["string", "null"]
                }
            }
        },
        "url": {
            "description": "A complete Url, with scheme, host and path.",
            "type": "object",
            "properties": {
                "raw": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "protocol": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "hostname": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "port": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "pathname": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "search": {
                    "description": "The search describes the query string of the request. It is expected to have values delimited by ampersands.",
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "hash": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                }
            }
        },
        "cookies": {
            "description": "A parsed key-value object of cookies",
            "type": ["object", "null"]
        }
    },
    "required": ["url", "method"]
        },
        "tags": {
            "type": ["object", "null"],
            "description": "A flat mapping of tags with values.",
            "regexProperties": true,
            "patternProperties": {
                "^[^.*\"]*$": {
                    "type": "string",
                    "maxLength": 1024
                }
            },
            "additionalProperties": false
        },
        "user": {
                "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "docs/spec/user.json",
    "title": "User",
    "description": "Describes the authenticated User for a request.",
    "type": ["object", "null"],
    "properties": {
        "id": {
            "description": "An id, identifying the logged in user, e.g. the primary key of the user",
            "type": ["string", "number", "null"],
            "maxLength": 1024
        },    
        "email": {
            "description": "The email address of the logged in user",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "username": {
            "description": "The username of the logged in user",
            "type": ["string", "null"],
            "maxLength": 1024
        }
    }
        }
    }
        },
        "culprit": {
            "description": "Function call which was the primary perpetrator of this event.",
            "type": ["string", "null"]
        },
        "exception": {
            "description": "A standard exception.",
            "type": ["object", "null"],
            "properties": {
                "code": {
                    "type": ["string", "number", "null"],
                    "maxLength": 1024
                },
                "message": {
                   "description": "The exception's error message.",
                   "type": "string"
                },
                "module": {
                    "description": "Describes the exception type's module namespace.",
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "attributes": {
                    "type": ["object", "null"]
                },
                "stacktrace": {
                    "type": ["array", "null"],
                    "items": {
                            "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "docs/spec/stacktrace_frame.json",
    "title": "Stacktrace",
    "type": "object",
    "description": "A stacktrace frame, contains various bits (most optional) describing the context of the frame",
    "properties": {
        "abs_path": {
            "description": "The absolute path of the file involved in the stack frame",
            "type": ["string", "null"]
        },
        "colno": {
            "description": "Column number",
            "type": ["number", "null"]
        },
        "context_line": {
            "description": "The line of code part of the stack frame",
            "type": ["string", "null"]
        },
        "filename": {
            "description": "The relative filename of the code involved in the stack frame, used e.g. to do error checksumming",
            "type": "string"
        },
        "function": {
            "description": "The function involved in the stack frame",
            "type": ["string", "null"]
        },
        "library_frame": {
            "description": "A boolean, indicating if this frame is from a library or user code",
            "type": ["boolean", "null"]
        },
        "lineno": {
            "description": "The line number of code part of the stack frame, used e.g. to do error checksumming",
            "type": "number"
        },
        "module": {
            "description": "The module to which frame belongs to",
            "type": ["string", "null"]
        },
        "post_context": {
            "description": "The lines of code after the stack frame",
            "type": ["array", "null"],
            "minItems": 0
        },
        "pre_context": {
            "description": "The lines of code before the stack frame",
             "type": ["array", "null"],
            "minItems": 0
        },
        "vars": {
            "description": "Local variables for this stack frame",
            "type": ["object", "null"],
            "properties": {}
        }
    },
    "required": ["filename", "lineno"]
                    },
                    "minItems": 0
                },
                "type": {
                    "type": ["string", "null"],
                    "maxLength": 1024
                },
                "uncaught": {
                    "type": ["boolean", "null"]
                }
            },
            "required": ["message"]
        },
        "id": {
            "type": ["string", "null"],
            "description": "UUID for the error",
            "pattern": "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"
        },
        "log": {
            "type": ["object", "null"],
            "properties": {
                "level": {
                    "description": "The record severity.",
                    "type": ["string", "null"],
                    "default": "error",
                    "enum": ["debug", "info", "warning", "error", "fatal", null],
                    "maxLength": 1024
                },
                "logger_name": {
                    "description": "The name of the logger which created the record.",
                    "type": ["string", "null"],
                    "default": "default",
                    "maxLength": 1024
                },
                "message": {
                    "description": "The exception's error message.",
                    "type": "string"
                },
                "param_message": {
                    "description": "A parametrized message. E.g. Could not connect to %s. The property message is still required, and should be equal to the param_message, but with placeholders replaced. In some situations the param_message is used to group errors together. The string is not interpreted, so feel free to use whichever placeholders makes sense in the client languange.",
                    "type": ["string", "null"],
                    "maxLength": 1024

                },
                "stacktrace": {
                    "type": ["array", "null"],
                    "items": {
                            "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "docs/spec/stacktrace_frame.json",
    "title": "Stacktrace",
    "type": "object",
    "description": "A stacktrace frame, contains various bits (most optional) describing the context of the frame",
    "properties": {
        "abs_path": {
            "description": "The absolute path of the file involved in the stack frame",
            "type": ["string", "null"]
        },
        "colno": {
            "description": "Column number",
            "type": ["number", "null"]
        },
        "context_line": {
            "description": "The line of code part of the stack frame",
            "type": ["string", "null"]
        },
        "filename": {
            "description": "The relative filename of the code involved in the stack frame, used e.g. to do error checksumming",
            "type": "string"
        },
        "function": {
            "description": "The function involved in the stack frame",
            "type": ["string", "null"]
        },
        "library_frame": {
            "description": "A boolean, indicating if this frame is from a library or user code",
            "type": ["boolean", "null"]
        },
        "lineno": {
            "description": "The line number of code part of the stack frame, used e.g. to do error checksumming",
            "type": "number"
        },
        "module": {
            "description": "The module to which frame belongs to",
            "type": ["string", "null"]
        },
        "post_context": {
            "description": "The lines of code after the stack frame",
            "type": ["array", "null"],
            "minItems": 0
        },
        "pre_context": {
            "description": "The lines of code before the stack frame",
             "type": ["array", "null"],
            "minItems": 0
        },
        "vars": {
            "description": "Local variables for this stack frame",
            "type": ["object", "null"],
            "properties": {}
        }
    },
    "required": ["filename", "lineno"]
                    },
                    "minItems": 0
                }
            },
            "required": ["message"]
        },
        "timestamp": {
            "type": "string",
            "format": "date-time",
            "pattern": "Z$",
            "description": "Recorded time of the error, UTC based and formatted as YYYY-MM-DDTHH:mm:ss.sssZ"
        }
    },
    "required": ["timestamp"],
    "anyOf": [
        {
            "required": ["exception"]
        },
        {
            "required": ["log"]
        }
    ]
            },
            "minItems": 1
        },
        "system": {
                "$schema": "http://json-schema.org/draft-04/schema#",
    "$id": "doc/spec/system.json",
    "title": "System",
    "type": ["object", "null"],
    "properties": {
        "architecture": {
            "description": "Architecture of the system the agent is running on.",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "hostname": {
            "description": "Hostname of the system the agent is running on.",
            "type": ["string", "null"],
            "maxLength": 1024
        },
        "platform": {
            "description": "Name of the system platform the agent is running on.",
            "type": ["string", "null"],
            "maxLength": 1024
        }
    }
        }
    },
    "required": ["service", "errors"]
}
`
