 VERSION := v0.0.1

BASEDIR := $(shell pwd)
PROJECTNAME := round_robin_with_weight

# note the diff of mac and windows
SOURCEFILES := $(shell find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./src/*" -print)

# init the dir for golang project
# build target for round robin with weight
BUILDSRCDIR := ${BASEDIR}/src
BUILDPKGDIR := ${BASEDIR}/pkg
BUILDDIR := ${BASEDIR}/bin
BUILDTARGET := ${BUILDDIR}/round_robin_with_weight
BUILDSODIR := ${BASEDIR}/so
BUILDSOTARGET := ${BUILDSODIR}/round_robin_with_weight.so

GOPATH := ${BASEDIR}
GO15VENDOREXPERIMENT := 1
GOENV := GOPATH=${GOPATH} GO15VENDOREXPERIMENT=${GO15VENDOREXPERIMENT}

# json package
# ref: https://github.com/json-iterator/go
BUILDARGS := -tags=jsoniter

LDFLAGS := -X round_robin_with_weight/config.Edition=default -X round_robin_with_weight/config.Version=${VERSION} -X round_robin_with_weight/config.BuildTime=`date +%Y-%m-%dT%T%z` -X round_robin_with_weight/config.Hash=`git rev-parse HEAD` -X round_robin_with_weight/config.Tag=`git describe --dirty --always`
DEBUGLDFLAGS := -n -X round_robin_with_weight/config.Mode=debug ${LDFLAGS}
RELEASELDFLAGS := -s -w -X round_robin_with_weight/config.Mode=release ${LDFLAGS}
GENSOFLAGS := -buildmode=c-shared

.PHONY: build
build: ${BUILDDIR} ${BUILDSRCDIR} ${SOURCEFILES}
	${GOENV} go build ${BUILDARGS} -i -ldflags "${DEBUGLDFLAGS}" -o ${BUILDTARGET} ${PROJECTNAME}

.PHONY: release
release: ${BUILDDIR} ${BUILDSRCDIR}
	${GOENV} GIN_MODE=release go build ${BUILDARGS} -v -ldflags "${RELEASELDFLAGS}" -o ${BUILDTARGET} ${PROJECTNAME}

.PHONY: genso
genso: ${BUILDDIR} ${BUILDSRCDIR} ${SOURCEFILES}
	${GOENV} go build ${GENSOFLAGS} -o ${BUILDSOTARGET} ${PROJECTNAME}

${BUILDDIR}:
	mkdir -p ${BUILDPKGDIR}
	mkdir -p ${BUILDDIR}

${BUILDSRCDIR}:
	mkdir -p ${BUILDSRCDIR}
	ln -s ${BASEDIR} ${BUILDSRCDIR}/${PROJECTNAME}


.PHONY: dev-init
dev-init: ${BUILDDIR} ${BUILDSRCDIR}


.PHONY: test
test:
	cd tests; ${GOENV} go test

.PHONY: codecheck
codecheck:
	${GOENV} find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./src/*" -exec gofmt -d {} \; -exec golint {} \;


.PHONY: clean
clean:
	rm -rf ${BUILDSRCDIR} || true
	rm -rf ${BUILDTARGET} || true
	rm -rf ${BUILDPKGDIR} || true
	rm -rf vendor || true
