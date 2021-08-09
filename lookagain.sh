#!/bin/bash
find . -name "*sh" | tr -d / | cut -d "." -f2 | sed 's/.sh//g'| sort -r