#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/whale" -importPath whale -srcPath "$SCRIPTPATH/src" -runMode prod
