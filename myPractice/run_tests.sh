#!/bin/bash
# ------------------------------------------------------------------
# [Ma Yuan] Run golang test
#          Options for golang test 
# ------------------------------------------------------------------

VERSION=0.1.0
SUBJECT="TEST_EXECUTOR"
RUNNING_SAMPLE=false
RUNNING_ALL=false
USAGE="Usage: run_tests -hva package name"

# --- Options processing -------------------------------------------
if [ $# == 0 ] ; then
    echo $USAGE
    exit 1;
fi

while getopts ":i:vh:as" optname
  do
    case "$optname" in
      "v")
        echo "Version $VERSION"
        exit 0;
        ;;
      "i")
        echo "-i argument: $OPTARG"
        ;;
      "h")
        echo $USAGE
        exit 0;
        ;;
      "a")
        RUNNING_ALL=true
        ;;
      "s")
        RUNNING_SAMPLE=true
        go test ./pkg/helloword
        exit 0;
        ;;
      "?")
        echo "Unknown option $OPTARG"
        exit 0;
        ;;
    esac
  done

shift $(($OPTIND - 1))

param1=$1

# --- Locks -------------------------------------------------------
LOCK_FILE=/tmp/$SUBJECT.lock
if [ -f "$LOCK_FILE" ]; then
   echo "Script is already running"
   exit
fi

trap "rm -f $LOCK_FILE" EXIT
touch $LOCK_FILE

# --- Body --------------------------------------------------------
#  SCRIPT LOGIC GOES HERE
if [ "$RUNNING_ALL" = true ] ; then
    echo 'all test will be run!'
    go test -v ./pkg/leetcode/*
  else
    echo 'running pkg '${param1}
    go test -v ./pkg/${param1}
fi
# -----------------------------------------------------------------


# go test ./pkg/*