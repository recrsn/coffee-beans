.PHONY all: clean assemble

TEMPLATES := $(shell find ./templates -type f -name '*.html')
PUBLIC := $(shell find ./public -type f)

templates: $(TEMPLATES)
	touch templates

assets/index.css: templates
	touch assets/index.css

static/index.css: assets/index.css
	npm run build

public: $(PUBLIC)
	mkdir -p static
	cp -r public/ static/
	touch public

static: public static/index.css
	touch static

coffee-beans/bin/coffee-beans: static templates
	mkdir -p coffee-beans/bin
	go build -o ./coffee-beans/bin/coffee-beans -tags=${BUILD_MODE:-debug} -ldflags "-X main.version=${GITHUB_REF#refs/*/} -X main.commit=${GITHUB_SHA:-HEAD}"

coffee-beans/etc/coffee-beans.yaml:
	mkdir -p coffee-beans/etc
	cp coffee-beans.yaml coffee-beans/etc/coffee-beans.yaml

coffee-beans: coffee-beans/bin/coffee-beans coffee-beans/etc/coffee-beans.yaml
	touch coffee-beans

coffee-beans.tar.gz: coffee-beans
	tar -czf coffee-beans.tar.gz coffee-beans

# Aliases

.PHONY compile: coffee-beans/bin/coffee-beans
.PHONY res: coffee-beans/etc/coffee-beans.yaml
.PHONY assemble: coffee-beans

.PHONY package: coffee-beans.tar.gz

.PHONY clean:
	rm -rf coffee-beans static

