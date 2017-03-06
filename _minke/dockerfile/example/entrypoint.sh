#!/bin/sh

consul-template -consul $CONSUL -exec-reload-signal=SIGTERM -template /example/config.ctmpl:/example/config.json -exec "/example/example /example/config.json /"
