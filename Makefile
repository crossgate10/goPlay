.PHONY: all

new:
	@cd pkg && mkdir exp$(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)) && \
	cd exp$(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)) && \
	go run ../new/main.go -no $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

