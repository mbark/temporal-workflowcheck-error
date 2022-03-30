TOOLS_BIN = ${CURDIR}/bin

workflowcheck: $(TOOLS_BIN)/.tools_versions
	PATH="$(TOOLS_BIN):$(PATH)" workflowcheck ./...

$(TOOLS_BIN)/.tools_versions: tools.go go.mod
	mkdir -p $(dir $@)
	cat $< | grep _ | awk -F'"' '{print $$2}' > $@
	cat $@ | xargs -tI % sh -c "GOBIN=$(TOOLS_BIN) go install %"
