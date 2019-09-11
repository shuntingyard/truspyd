EXEC := truspyd.service

INSTALLROOT := /usr/lib/systemd/system

UTIL := /bin/systemctl

.PHONY: help install uninstall

help:
	@grep '^[^#[:space:]].*:' Makefile

install: $(EXEC)
	@cp $(EXEC) $(INSTALLROOT)
	@echo "$(EXEC) installed"

uninstall:
	$(UTIL) stop $(EXEC)
	$(UTIL) disable $(EXEC)
	@rm -f $(INSTALLROOT)/$(EXEC)
	@echo "$(EXEC) successfully uninstalled"
