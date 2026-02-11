#!/bin/bash

go build -o bandb ./cmd/web/*.go && ./bandb serve