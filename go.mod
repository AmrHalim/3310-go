module github.com/AmrHalim/3310-go

go 1.21.5

replace utils v0.0.0 => ./packages/utils

replace encoder v0.0.0 => ./packages/encoder

replace decoder v0.0.0 => ./packages/decoder

require (
	decoder v0.0.0
	encoder v0.0.0
)

require utils v0.0.0 // indirect
