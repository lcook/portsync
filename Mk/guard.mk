#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
_VARS=		DIR \
		LATEST \
		MAINTAINER \
		ORIGIN \
		TYPE \
		VERSION
.for var in ${_VARS}
.  if !defined(PACKAGE_${var})
.error PACKAGE_${var} is not defined.
.  endif
.endfor
