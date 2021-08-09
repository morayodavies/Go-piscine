#!/bin/bash
find . -name "*sh" | tr "/" "."| cut -d "." -f3 | sed 's/.sh//g'| sort -r