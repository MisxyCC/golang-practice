module misxy/core

go 1.23.4

replace misxy/concurrency => ../concurrency

require (
	misxy/concurrency v0.0.0-00010101000000-000000000000
	misxy/context v0.0.0-00010101000000-000000000000
)

replace misxy/context => ../context
