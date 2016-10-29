# These need to be set for your local environment.
GOROOT    ?= /opt/go/1.7.3/go
#ATOM_PATH ?= /opt/atom-1.11.2-amd64/atom
ATOM_PATH ?= /opt/atom/Atom.app/Contents/MacOS/Atom

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

# Pre-flight
preflight: golint crypto-ssh

# Run the editor.
edit: golint crypto-ssh
	GOPATH=$$(pwd) $(ATOM_PATH)

# Install the basic infrastructure for the atom editor.
.PHONY: golint crypto-ssh

golint: $(GOROOT)/bin/golint

crypto-ssh: src/golang.org/x/crypto/ssh

$(GOROOT)/bin/golint: golint/bin/golint
	sudo cp $< $@

golint/bin/golint:
	GOPATH=$$(pwd)/golint go get -u github.com/golang/lint/golint

src/golang.org/x/crypto/ssh:
	GOPATH=$$(pwd) go get golang.org/x/crypto/ssh
