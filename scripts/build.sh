#!/bin/bash

make sources
WORKDIR=`pwd`
TMPDIR=`mktemp -d /tmp/rpmbuild.XXXXXXXXXX` 
mv *.tar.gz ${TMPDIR} 
cd ${TMPDIR} 
tar -xzf *.tar.gz 
find . -name '*.spec' -exec yum-builddep {} \; 
rpmbuild -ta --define='dist ."el7"' *gz
mkdir -p ${WORKDIR}/target
mv ./* ${WORKDIR}/target