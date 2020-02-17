local_resource("Create test dir", "git clone git@github.com:windmilleng/tilt-extensions.git /tmp/fake-extensions || cd /tmp/fake-extensions ; git pull", trigger_mode=TRIGGER_MODE_MANUAL)

local_resource("tilt-extensions-ci", "go run main.go /tmp/fake-extensions", deps=['go.mod', 'go.sum', 'main.go'])
