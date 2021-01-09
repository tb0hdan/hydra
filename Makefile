VERSION = $(shell cat ./VERSION)

tag:
	@git tag -a v$(VERSION) -m v$(VERSION)
	@git push --tags
