package postgresworker

import (
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/codebase/factory"
)

type (
	option struct {
		maxGoroutines int
		debugMode     bool
		locker        candiutils.Locker

		sources map[string]*PostgresSource
	}

	// OptionFunc type
	OptionFunc func(*option)
)

func getDefaultOption(service factory.ServiceFactory) option {
	opt := option{
		maxGoroutines: 1,
		debugMode:     true,
		sources:       make(map[string]*PostgresSource),
	}
	if redisPool := service.GetDependency().GetRedisPool(); redisPool != nil {
		opt.locker = candiutils.NewRedisLocker(redisPool.WritePool())
	} else {
		opt.locker = &candiutils.NoopLocker{}
	}
	return opt
}

// SetPostgresDSN option func
func SetPostgresDSN(dsn string) OptionFunc {
	return func(o *option) {
		o.sources[""] = &PostgresSource{dsn: dsn}
	}
}

// SetMaxGoroutines option func
func SetMaxGoroutines(maxGoroutines int) OptionFunc {
	return func(o *option) {
		o.maxGoroutines = maxGoroutines
	}
}

// SetDebugMode option func
func SetDebugMode(debugMode bool) OptionFunc {
	return func(o *option) {
		o.debugMode = debugMode
	}
}

// SetLocker option func
func SetLocker(locker candiutils.Locker) OptionFunc {
	return func(o *option) {
		o.locker = locker
	}
}

// AddPostgresDSN option func for add multple postgres source to be listen
func AddPostgresDSN(sourceName, dsn string) OptionFunc {
	return func(o *option) {
		o.sources[sourceName] = &PostgresSource{
			name: sourceName, dsn: dsn,
		}
	}
}
