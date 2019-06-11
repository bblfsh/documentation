module github.com/bzz/documentation

go 1.12

require (
	github.com/bblfsh/sdk/v3 v3.1.0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/heroku/docker-registry-client v0.0.0-20181004091502-47ecf50fd8d4 // indirect
	golang.org/x/tools v0.0.0-20190610231749-f8d1dee965f7
)

// FIXME: remove when https://github.com/bblfsh/sdk/pull/408 is merged
replace github.com/bblfsh/sdk/v3 => github.com/dennwc/sdk/v3 v3.0.0-20190611200859-87bf1e433b56
