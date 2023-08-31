#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
PROG=		portsync
VERSION=	0.2

LOCALBASE?=	/usr/local
BINDIR=		${LOCALBASE}/bin
SHAREDIR=	${LOCALBASE}/share

GO_DEFAULT?=	1.21
GO_SUFFIX=	${GO_DEFAULT:S/.//}
GO_CMD=		${BINDIR}/go${GO_SUFFIX}

GO_MODULE=	github.com/lcook/${PROG}
GO_FLAGS=	-v -ldflags "-s -w -X '${GO_MODULE}/cmd.version=${VERSION}'"
.if exists(${.CURDIR}/.git)
SHA!=		git rev-parse --short HEAD
BRANCH!=	git symbolic-ref HEAD | sed 's,refs/heads/,,'
VERSION:=	${BRANCH}/${VERSION}-${SHA}
.endif

all: build

build:
	${GO_CMD} build ${GO_FLAGS} -o ${PROG}

clean:
	${GO_CMD} clean -x

install:
	mkdir -p ${SHAREDIR}/${PROG}/Mk
	cp -vR Mk ${SHAREDIR}/${PROG}
	cp -v ${PROG} ${BINDIR}

uninstall:
	rm -rfv ${SHAREDIR}/${PROG}
	rm -rf ${BINDIR}/${PROG}

.PHONY: all build clean install
