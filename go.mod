module github.com/windmilleng/tilt-extensions-ci

go 1.13

require github.com/windmilleng/tilt v0.12.2-0.20200217213441-80ebaa2c932a

// TODO(dmiller) remove this replace once https://github.com/moby/buildkit/pull/1297 is merged
replace github.com/moby/buildkit => github.com/zachbadgett/buildkit v0.6.2-0.20191220071605-814e2794095f
