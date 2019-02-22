# See LICENSE file for copyright and license details.

include config.mk

SRC = ${NAME}.go

all: options ${NAME}

options:
	@echo ${NAME} build options:
	@echo "GO       = ${GO}"

${OBJ}: config.mk

${NAME}: ${OBJ}
	@${GO}

install: all
	@echo installing executable file to ${DESTDIR}${PREFIX}/bin
	@mkdir -p ${DESTDIR}${PREFIX}/bin
	@cp -f ${NAME} ${DESTDIR}${PREFIX}/bin
	@chmod 755 ${DESTDIR}${PREFIX}/bin/${NAME}

uninstall:
	@echo removing executable file from ${DESTDIR}${PREFIX}/bin
	@rm -f ${DESTDIR}${PREFIX}/bin/${NAME}
	@rm -f ${NAME}

.PHONY: all options clean dist install uninstall
