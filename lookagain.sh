#!/bin/bash
find . -name "*sh" | sed 's/.sh//g' | sort -r