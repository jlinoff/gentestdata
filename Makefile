# Assumes that go is in your path.

# Compile.
.PHONY: compile test
compile:  bin/gentestdata

# These are not really tests because I am verifying the functionality
# manually.
define Banner
	@echo
	@echo "# ================================================================"
	@echo "# $1"
	@echo "# ================================================================"
endef

define Test
	$(call Banner,$1 $2)
	$1 $2
endef

test: bin/gentestdata
	$(call Test,$<,-h)
	$(call Test,$<,-V)
	$(call Test,$<,)
	$(call Test,$<,-w 32 -n 16)
	$(call Test,$<,-w 32 -n 16 -l)
	$(call Test,$<,-w 32 -n 16 -l -i 5)
	$(call Test,$<,-w 32 -n 16 -l -i 5 1>/dev/null)
	$(call Test,$<,-w 32 -n 16 -l -i 5 2>/dev/null)
	$(call Banner,done - passed)

clean:
	find . -type f -name '*~' -delete
	rm -f bin/gentestdata

bin/gentestdata: src/jlinoff/gentestdata/main.go src/jlinoff/gentestdata/options.go
	GOPATH=$$(pwd) go install jlinoff/gentestdata

