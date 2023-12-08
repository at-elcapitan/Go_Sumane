# Copyright (C) 2023 ElCapitan; pungentee
# 
# This file is part of GoSM.
# 
# GoSM is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
# 
# GoSM is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
# 
# You should have received a copy of the GNU General Public License
# along with GoSM.  If not, see <http://www.gnu.org/licenses/>.

GOCOM 				:= go
GO_FLAGS 			:= -buildmode=exe
OUTPUT_DIRECTORY 	:= out/
SOURCE_DIRECTORY 	:= src/
SOURCES     		:= $(wildcard $(SOURCE_DIRECTORY)*.go)
OUTPUT 				:= main

.PHONY : all
all : build
	@echo "Done!"

build: dirs
	@echo "Building project..."
	go build -o $(OUTPUT_DIRECTORY)$(OUTPUT) $(GO_FLAGS) $(SOURCES)

dirs:
	@mkdir -p $(OUTPUT_DIRECTORY)

.PHONY : clean
clean :
	rm -rf out/*

.PHONY : run
run : 
	@$(OUTPUT_DIRECTORY)$(OUTPUT)