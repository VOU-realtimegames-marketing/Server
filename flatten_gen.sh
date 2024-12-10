#!/bin/bash
find proto/gen -type f -name '*.go' -exec mv {} proto/gen/ \;
find proto/gen -type d -empty -delete
