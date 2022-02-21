module ir_receiver

go 1.17

require (
	github.com/neildavis/tinygo_modules v0.0.0-00010101000000-000000000000
	tinygo.org/x/drivers v0.19.0
)

replace github.com/neildavis/tinygo_modules => ../../modules

replace tinygo.org/x/drivers v0.19.0 => github.com/neildavis/drivers v0.19.1-0.20220221142459-2f58b6ba7994
