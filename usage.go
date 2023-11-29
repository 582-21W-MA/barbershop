package main

const usage = `
Barbershop is a simple static site generator using Mustache templates.

Usage:
    barbershop <input_directory> [-h] [--help]
    barbershop serve <root_directory>
    barbershop watch <input_directory>

Subcommands:
    serve       Start a development web server at the given root directory.
    watch       Watch the given directory for change, run barberhop
                each time a file is added, removed or modified,
                and start a development web server.

Arguments:
    input_directory     Path to directory containing template files.
    root_directory      Path to website root.

Examples:
    barbershop src
    barbershop serve site
    barbershop watch src

Options:
    -h | --help         Print help.
`
