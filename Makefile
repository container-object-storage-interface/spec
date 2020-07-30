all: build

COSI_SPEC := spec.md
COSI_PROTO := cosi.proto

build:
	$(MAKE) -C lib/go

clean:
	$(MAKE) -C lib/go $@

clobber: clean
	$(MAKE) -C lib/go $@
	rm -f $(COSI_PROTO) $(COSI_PROTO).tmp

# check generated files for violation of standards
check: $(COSI_PROTO)
	awk '{ if (length > 122) print NR, $$0 }' $? | diff - /dev/null

.PHONY: clean clobber check
