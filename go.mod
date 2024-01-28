module github.com/AmrHalim/3310-go

go 1.21.5

replace github.com/AmrHalim/utils v0.0.0 => ./packages/utils

replace github.com/AmrHalim/encoder v0.0.0 => ./packages/encoder

replace github.com/AmrHalim/decoder v0.0.0 => ./packages/decoder

require (
	github.com/AmrHalim/decoder v0.0.0
	github.com/AmrHalim/encoder v0.0.0
)

require github.com/AmrHalim/utils v0.0.0 // indirect
