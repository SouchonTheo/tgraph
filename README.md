# tgraph

## Purpose

`tgraph` is a command-line tool designed to generate dependency graphs for Scala repositories using Bloop and dot.

## Usage Examples

### Generate a Dependency Graph

Generate a dependency graph for the project "my_project" with output in PNG format.

``` bash
tgraph -n my_project -f output.png
```

### Exclude Specific Folders

Exclude "test" and "integration" folders while generating the graph for the project "my_project".

``` bash
tgraph -n my_project -i foo -i bar -f output.png
```

### Custom Output and Visualization

Generate a graph for the project "my_project" and save it as SVG format. Open the graph using a specific application.

``` bash
tgraph -n my_project -f output.svg -o my_visualizer_app
```

### General Use Case

``` bash
tgraph ([-n project]) ([-i ignore]) ([-f output | -o open])
```

For more options and details, you can use the help command:

``` bash
tgraph -h
```

## Installation

For now install and compile it on your own

Working to make it work with homebrew

## Prerequisites

- [Bloop](https://scalacenter.github.io/bloop/)
- [Dot](https://graphviz.org/doc/info/lang.html)
