package watcher

import "time"

type WatcherBuilder interface {
	SetDelay(delay time.Duration) WatcherBuilder
	SetDebug(enabled bool) WatcherBuilder
	SetExtFilter(filters ...string) WatcherBuilder
	Build() Watcher
}

type watcherBuilder struct {
	filters []string
	delay time.Duration
	debug bool
}

func New() WatcherBuilder {
	return &watcherBuilder{}
}

// Set extension files to filter
func (wb *watcherBuilder) SetExtFilter(filters ...string) WatcherBuilder {
	wb.filters = filters
	return wb
}

// SetDelay fo next execution
func (wb *watcherBuilder) SetDelay(delay time.Duration) WatcherBuilder {
	if delay == 0 {
		// default value
		delay = 3 * time.Second
	}
	wb.delay = delay
	return wb
}

// Enable/disable debugging output
func (wb *watcherBuilder) SetDebug(enabled bool) WatcherBuilder {
	wb.debug = enabled
	return wb
}

func (wb *watcherBuilder) Build() Watcher {
	return &watcher{
		filters: wb.filters,
		files: []string{},
		debug: wb.debug,
		delay: wb.delay,
	}
}