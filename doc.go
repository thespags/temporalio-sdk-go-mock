// Package temporalmocks exists so go.temporal.io/sdk stays a real dependency in
// go.mod (and therefore visible to Dependabot and mockery's package loader).
// The generated mocks live in ./mocks.
package temporalmocks

import (
	_ "go.temporal.io/sdk/activity"
	_ "go.temporal.io/sdk/client"
	_ "go.temporal.io/sdk/converter"
	_ "go.temporal.io/sdk/interceptor"
	_ "go.temporal.io/sdk/log"
	_ "go.temporal.io/sdk/temporal"
	_ "go.temporal.io/sdk/testsuite"
	_ "go.temporal.io/sdk/worker"
	_ "go.temporal.io/sdk/workflow"
)
